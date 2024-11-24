package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "jayce/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto
	pbserver "jayce/grpc/jayce-server"

	"google.golang.org/grpc"
)

const (
	brokerAddress = "container_broker:50054" // Dirección y puerto del Broker
)

type Registro struct {
    UltimaRespuesta *pbserver.JayceResponse // Última respuesta recibida
    RelojVectorial  *pbserver.VectorClock              // Último reloj vectorial asociado
	UltimoServidor string
}

type Jayce struct {
	client    pb.JayceBrokerServiceClient
	clientServer pbserver.JayceServerServiceClient
	consultas []Consulta // Almacena las consultas realizadas
	connPool map[string]*grpc.ClientConn
	serverClients map[string]pbserver.JayceServerServiceClient
	registros map[string]Registro
	inconsistencias []Consulta
}

type Consulta struct {
	Solicitud *pbserver.JayceRequest
	Respuesta *pbserver.JayceResponse
	Error     error
}

// NewJayce crea una nueva instancia de Jayce y conecta con el Broker
func NewJayce() *Jayce {
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Broker: %v", err)
	}

	client := pb.NewJayceBrokerServiceClient(conn)
	log.Println("Conectado exitosamente al Broker(instancia Jayce creada)")
	return &Jayce{client: client, connPool: make(map[string]*grpc.ClientConn), serverClients: make(map[string]pbserver.JayceServerServiceClient), registros: make(map[string]Registro), inconsistencias: []Consulta{}}
}

func (j *Jayce) VerificarMonotonicidad(region, producto string, nuevoReloj *pbserver.VectorClock) bool {
    key := region + "-" + producto
    registro, existe := j.registros[key]
    if !existe {
        return true // Si no hay registro previo, la lectura es válida
    }

    // Compara los valores de los relojes vectoriales

    if nuevoReloj.Server1 < registro.RelojVectorial.Server1 || nuevoReloj.Server2 < registro.RelojVectorial.Server2 || nuevoReloj.Server3 < registro.RelojVectorial.Server3{
        return false // Retrocede en el tiempo
    }
    

    return true // La lectura es consistente
}


func (j *Jayce) ActualizarRegistro(region, producto string, res *pbserver.JayceResponse, servidor string) {
    key := region + "-" + producto
	log.Print("ACTUALIZAR REGISTRO")
	log.Print(key)
    j.registros[key] = Registro{
        UltimaRespuesta: res,
        RelojVectorial:  res.VectorClock,
		UltimoServidor: servidor,
    }
    log.Printf("Registro actualizado: Región: %s, Producto: %s, Reloj: %v", region, producto, res.VectorClock)
}



// ObtenerProducto envía una solicitud de consulta al Broker para obtener el puerto del servidor asignado
func (j *Jayce) ObtenerServidor(region, product string) (*pb.JayceRequest, string, error) {
	// Crea la solicitud
	req := &pb.JayceRequest{
		Region:      region,
		ProductName: product,
	}

	// Define un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Envía la solicitud al Broker
	log.Printf("Enviando solicitud al Broker: Región: %s, Producto: %s", req.Region, req.ProductName)
	res, err := j.client.ObtenerServidor(ctx, req)
	if err != nil {
		log.Printf("Error al obtener el servidor: %v", err)
		return req, "", err
	}
	log.Printf("Solicitud enviada correctamente al Broker: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Verifica el estado de la respuesta
	if res.Status == pb.ResponseStatus_ERROR {
		log.Print("Error en la consulta")
		return req, "", fmt.Errorf("error en la consulta")
	}
	// Extrae el puerto del mensaje de respuesta
	puerto := ""
	if res.Message != nil {
		puerto = *(res.Message)
	}

	// Log de la respuesta del Broker
	log.Printf("Respuesta del Broker: Status: %v, Message: %s", res.Status, puerto)

	// Imprime los datos de la respuesta
	fmt.Printf("Producto consultado: %s en %s\n", product, region)
	fmt.Printf("Mensaje del Broker: %s\n", puerto)


	return req, puerto, nil
}

func (j *Jayce) getServerClient(direccion string) pbserver.JayceServerServiceClient {
    if client, exists := j.serverClients[direccion]; exists {
        return client // Reutiliza el cliente si ya existe
    }

    conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al servidor: %v", err)
    }

    // Almacena la conexión y el cliente en el pool
    j.connPool[direccion] = conn
    client := pbserver.NewJayceServerServiceClient(conn)
    j.serverClients[direccion] = client

    return client
}

func (j *Jayce) CloseConnections() {
    for _, conn := range j.connPool {
        if err := conn.Close(); err != nil {
            log.Printf("Error al cerrar la conexión: %v", err)
        }
    }
}


func (j *Jayce) ObtenerProducto(region, product string) (error) {
	req := &pbserver.JayceRequest{
		Region: region,
		ProductName: product,
	}

	_, direccion, err := j.ObtenerServidor(region, product)
	if err != nil {
		return err
	}

	client := j.getServerClient(direccion)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	log.Printf("Enviando solicitud al Servidor: Región: %s, Producto: %s", req.Region, req.ProductName)
	res, err := client.ObtenerProducto(ctx, req)
	if err != nil {
		log.Printf("Error al obtener el producto: %v", err)
		return err
	}
	log.Printf("Solicitud enviada correctamente al Servidor: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Verificar consistencia Monotonic Reads
	if !j.VerificarMonotonicidad(region, product, res.VectorClock) {
		log.Printf("Inconsistencia detectada: Región: %s, Producto: %s, Reloj recibido: %v, Último reloj: %v",
			region, product, res.VectorClock, j.registros[region+"-"+product].RelojVectorial)
	
		// Registrar inconsistencia para análisis posterior
		j.inconsistencias = append(j.inconsistencias, Consulta{Solicitud: req, Respuesta: res, Error: fmt.Errorf("lectura no consistente")})
	
		return fmt.Errorf("lectura no consistente: retrocede en el tiempo para Región: %s, Producto: %s", region, product)
	}

	j.ActualizarRegistro(region, product, res, direccion)

	cantidad := res.Cantidad

	vectorClock := res.VectorClock

	log.Printf("Respuesta del Servidor: Cantidad: %v, Reloj de Vectores: %s", cantidad, vectorClock)

	j.AlmacenarConsulta(req, res, err)

	return err
}
func (j *Jayce) MostrarInconsistencias() {
    if len(j.inconsistencias) == 0 {
        log.Println("No se detectaron inconsistencias.")
        return
    }

    log.Println("Inconsistencias detectadas:")
    for i, inconsistencia := range j.inconsistencias {
        log.Printf("[%d] Región: %s, Producto: %s, Error: %v",
            i+1, inconsistencia.Solicitud.Region, inconsistencia.Solicitud.ProductName, inconsistencia.Error)
    }
}


// AlmacenarConsulta guarda la solicitud y la respuesta en la memoria si no hay error
func (j *Jayce) AlmacenarConsulta(req *pbserver.JayceRequest, res *pbserver.JayceResponse, err error) {
	if err == nil {
		j.consultas = append(j.consultas, Consulta{Solicitud: req, Respuesta: res, Error: err})
		log.Printf("Consulta agregada correctamente: Región: %s, Producto: %s", req.Region, req.ProductName)
	} else {
		log.Printf("No se pudo agregar la consulta, HUBO UN ERROR AL OBTENER EL PRODUCTO(FUNCION Obtener Producto)!: Región: %s, Producto: %s, Error: %v", req.Region, req.ProductName, err)
	}
}

func menu(j *Jayce){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		log.Print("Quiere obtener un producto?")
		log.Print("1) Obtener Producto")
		log.Print("2) Salir")
		scanner.Scan()
		entrada := strings.TrimSpace(scanner.Text())

		opcion, err := strconv.Atoi(entrada)
		if err != nil {
			log.Print("Por favor ingrese un número válido")
		}
		switch opcion {
		case 1:
			scanner1 := bufio.NewScanner(os.Stdin)
			log.Print("Obtener Producto")
			log.Print("Ingrese la región del producto: ")
			scanner1.Scan()
			region := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nombre del producto a obtener: ")
			scanner1.Scan()
			producto := strings.TrimSpace(scanner1.Text())

			err1:= j.ObtenerProducto(region, producto)
			if err1 != nil {
				log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err1)
			} else {
				log.Printf("Consulta exitosa")
			}
		case 2:
			log.Print("Opción Salir")
			return
		default:
			log.Print("Opción no válida, por favor escoja un número entre 1 y 2.")
		}
	}
}

func main() {
	// Inicializa a Jayce y realiza consultas
	log.Println("Iniciando servidor Jayce...")
	jayce := NewJayce()
	defer jayce.CloseConnections()
	log.Println("Jayce ha sido inicializado")

	menu(jayce)

	// Ejemplo de consultas realizadas por Jayce
/* 	err1 := jayce.ObtenerProducto("Noxus", "Vino")
	if err1 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err1)
	} else {
		log.Printf("Consulta exitosa")
	}
	err2 := jayce.ObtenerProducto("a", "A")
	//err2 := jayce.ObtenerProducto("Demacia", "Espadas")
	if err2 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err2)
	} else {
		log.Printf("Consulta exitosa")
	}
	err3 := jayce.ObtenerProducto("b", "B")
	//err3 := jayce.ObtenerProducto("Piltover", "Cristales Hextech")
	if err3 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err3)
	} else {
		log.Printf("Consulta exitosa")
	} */
	// Imprime las consultas almacenadas
	for _, consulta := range jayce.consultas {
		log.Printf("Consulta: Región: %s, Producto: %s, Respuesta: %v, Error: %v",
			consulta.Solicitud.Region, consulta.Solicitud.ProductName, consulta.Respuesta, consulta.Error)
	}
	jayce.MostrarInconsistencias()
}
