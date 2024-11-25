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

	pb "supervisores_hexgate/grpc/sup-broker"  // Proto para comunicarse con el Broker
	hexpb "supervisores_hexgate/grpc/sup-serv" // Proto para comunicarse con los Servidores Hextech

	"google.golang.org/grpc"
)

const (
	brokerAddress = "container_broker:50054" // Dirección del Broker
)

// **Estructura de Registro**
type Registro struct {
	Region      string           // Región donde se encuentra el registro
	Producto    string           // Nombre del producto
	Valor       int32            // Valor (cantidad) del producto
	VectorClock *hexpb.VectorClock // Reloj vectorial de la última versión
	Servidor    string           // Dirección del servidor Hextech que respondió
}

// **Estructura del Supervisor**
type Supervisor struct {
	client    pb.BrokerServiceClient // Cliente para comunicarse con el Broker
	conn      *grpc.ClientConn       // Conexión al Broker
	registros map[string]Registro    // Almacena el estado local
	hexClients map[string]hexpb.HextechServiceClient
	connPool map[string]*grpc.ClientConn
}

// **Crear un nuevo Supervisor**
func NuevoSupervisor() *Supervisor {
	log.Print("Intentando conectar al Broker...")
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Broker: %v", err)
	}
	log.Print("Conexión exitosa al Broker")
	client := pb.NewBrokerServiceClient(conn)
	return &Supervisor{client: client, conn: conn, registros: make(map[string]Registro), hexClients: make(map[string]hexpb.HextechServiceClient), connPool: make(map[string]*grpc.ClientConn)}
}

func (s *Supervisor) getHexClient(direccion string) hexpb.HextechServiceClient {
    if client, exists := s.hexClients[direccion]; exists {
        return client // Reutiliza el cliente si ya existe
    }

    conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Servidor Hextech: %v", err)
    }
    s.connPool[direccion] = conn
    client := hexpb.NewHextechServiceClient(conn)
    s.hexClients[direccion] = client
    return client
}


// **Obtener dirección del servidor Hextech desde el Broker**
func (s *Supervisor) GetServer(region string) string {
	req := &pb.ServerRequest{
		Region: region,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := s.client.GetServer(ctx, req)
	if err != nil {
		log.Fatalf("Error al obtener servidor del Broker: %v", err)
	}
	log.Printf("Servidor asignado: %s", res.ServerAddress)
	return res.ServerAddress
}

// **Comunicación directa con el Servidor Hextech**
func (s *Supervisor) EnviarSolicitudAServidor(direccion string, req *hexpb.SupervisorRequest) bool {
	client := s.getHexClient(direccion)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Manejo del reloj vectorial local
	key := fmt.Sprintf("%s-%s", req.Region, req.ProductName)
	if registro, existe := s.registros[key]; existe {
		req.KnownVectorClock = registro.VectorClock
	} else {
		req.KnownVectorClock = &hexpb.VectorClock{Server1: 0, Server2: 0, Server3: 0}
	}

	// Procesar la solicitud
	res, err := client.HandleRequest(ctx, req)
	if err != nil {
		log.Fatalf("Error al procesar la solicitud en el Servidor Hextech: %v", err)
	}

	// Manejar la respuesta del servidor
	if res.Status == hexpb.ResponseStatus_OK {
		fmt.Printf("Operación exitosa: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)

		// Actualización de registros locales
		if req.OperationType != hexpb.OperationType_BORRAR {
			s.registros[key] = Registro{
				Region:      req.Region,
				Producto:    req.ProductName,
				VectorClock: res.VectorClock,
				Servidor:    direccion,
			}
			if res.Value != nil {
				s.registros[key] = Registro{
					Valor: *(res.Value),
				}
			}
		}
		return true
	} else if res.Status == hexpb.ResponseStatus_ERROR {
		fmt.Printf("Error: %s\n", *(res.Message))
		return false
	}
	return false
}

// **Informar inconsistencia al Broker y obtener nuevo servidor**
func (s *Supervisor) ResolverInconsistencia(region, producto string, vectorLocal *hexpb.VectorClock) string {
	currentVectorClock := &pb.VectorClock{
		Server1: vectorLocal.Server1,
		Server2: vectorLocal.Server2,
		Server3: vectorLocal.Server3,
	}
	req := &pb.InconsistencyRequest{
		Region:           region,
		ProductName:      producto,
		SupervisorClock:  currentVectorClock,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := s.client.ResolveInconsistency(ctx, req)
	if err != nil {
		log.Fatalf("Error al resolver inconsistencia con el Broker: %v", err)
	}

	log.Printf("Servidor asignado para resolver inconsistencia: %s", res.ServerAddress)
	return res.ServerAddress
}

// **Validación de consistencia (Read Your Writes)**
func (s *Supervisor) LeerConsistente(region, producto string) bool {
	key := fmt.Sprintf("%s-%s", region, producto)
	registro, existe := s.registros[key]

	if !existe {
		log.Fatalf("No existe información local para %s en la región %s", producto, region)
	}

	direccion := s.GetServer(region)
	req := &hexpb.SupervisorRequest{
		Region:           region,
		ProductName:      producto,
		OperationType:    hexpb.OperationType_ACTUALIZAR,
		KnownVectorClock: registro.VectorClock,
	}

	// Primero, obtendremos el reloj vectorial del servidor para la consistencia
	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor Hextech: %v", err)
	}
	defer conn.Close()

	client := hexpb.NewHextechServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Hacemos la solicitud al servidor Hextech
	res, err := client.HandleRequest(ctx, req)
	if err != nil {
		log.Fatalf("Error al obtener respuesta del servidor Hextech: %v", err)
	}

	// Verifica si el reloj vectorial del servidor es más reciente que el local
	if !esConsistente(res.VectorClock, registro.VectorClock) {
		// Si los relojes no son consistentes, volvemos a intentarlo
		log.Print("Datos no consistentes. Reintentando...")

		// Configurar reintento con tiempo de espera
		time.Sleep(100 * time.Millisecond) // Se podría hacer configurable

		// El método podría continuar reintentando hasta que se logre consistencia
		// Es importante evitar que el bucle quede atrapado si la consistencia no se resuelve
		return false
	}

	// Si los relojes son consistentes, continuamos
	fmt.Printf("Lectura consistente: [%d, %d, %d] - Valor: %d\n",
		res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Value)
	return true
}

// **Comparación de relojes vectoriales**
func esConsistente(v1, v2 *hexpb.VectorClock) bool {
	// Si el reloj v1 es mayor o igual que v2 en todas las posiciones, se considera consistente
	return v1.Server1 >= v2.Server1 && v1.Server2 >= v2.Server2 && v1.Server3 >= v2.Server3
}

// **Operaciones del Supervisor**
func (s *Supervisor) AgregarProducto(region, producto string, valor int32) {
	direccion := s.GetServer(region)
	req := &hexpb.SupervisorRequest{
		Region:        region,
		ProductName:   producto,
		OperationType: hexpb.OperationType_AGREGAR,
		Value:         &valor,
	}
	s.EnviarSolicitudAServidor(direccion, req)
}

// RenombrarProducto renombra un producto en el registro de una región
func (s *Supervisor) RenombrarProducto(region, producto, nuevoNombre string) {
    /* if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else { */
		nuevo := nuevoNombre
		direccion := s.GetServer(region)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
			NewProductName: &nuevo,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	//}
}

// ActualizarValor actualiza el valor de un producto en el registro de una región
func (s *Supervisor) ActualizarValor(region, producto string, nuevoValor int32) {
    /* if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else { */
		direccion := s.GetServer(region)
		req := &hexpb.SupervisorRequest{
			Region:       region,
			ProductName:  producto,
			OperationType: hexpb.OperationType_ACTUALIZAR,
			Value:        &nuevoValor,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	//}
}

// BorrarProducto elimina un producto del registro de una región
func (s *Supervisor) BorrarProducto(region, producto string) {
	/* if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else { */
		direccion := s.GetServer(region)
		req := &hexpb.SupervisorRequest{
			Region:       region,
			ProductName:  producto,
			OperationType: hexpb.OperationType_BORRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	//}
}

func menu(s *Supervisor){
	scanner := bufio.NewScanner(os.Stdin)
	for {

		log.Print("Opciones disponibles:")
		log.Print("1) Agregar Producto")
		log.Print("2) Renombrar Producto")
		log.Print("3) Actualizar Valor de Producto")
		log.Print("4) Eliminar Producto")
		log.Print("5) Salir")
		log.Print("Seleccione acción: ")
		scanner.Scan()
		entrada := strings.TrimSpace(scanner.Text())

		opcion, err := strconv.Atoi(entrada)
		if err != nil {
			log.Print("Por favor ingrese un número válido")
		}

		switch opcion{
		case 1:
			scanner1 := bufio.NewScanner(os.Stdin)
			log.Print("Opcion Agregar Producto")
			log.Print("Ingrese la región a enviar: ")
			scanner1.Scan()
			region := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el producto a enviar: ")
			scanner1.Scan()
			producto := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el valor del producto: ")
			scanner1.Scan()
			a := strings.TrimSpace(scanner1.Text())
			valor, err1 := strconv.Atoi(a)
			if err1 != nil {
				log.Print("Por favor ingrese un número válido")
			}
			
			s.AgregarProducto(region, producto, int32(valor))

		case 2:
			scanner1 := bufio.NewScanner(os.Stdin)
			log.Print("Opcion Renombrar Producto")
			log.Print("Ingrese la región del producto: ")
			scanner1.Scan()
			region := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nombre del producto a renombrar: ")
			scanner1.Scan()
			producto := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nuevo nombre del producto: ")
			scanner1.Scan()
			newName := strings.TrimSpace(scanner1.Text())

			s.RenombrarProducto(region, producto, newName)

		case 3:
			scanner1 := bufio.NewScanner(os.Stdin)
			log.Print("Opcion Actualizar Valor")
			log.Print("Ingrese la región del producto: ")
			scanner1.Scan()
			region := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nombre del producto: ")
			scanner1.Scan()
			producto := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nuevo valor del producto: ")
			scanner1.Scan()
			a := strings.TrimSpace(scanner1.Text())
			valor, err1 := strconv.Atoi(a)
			if err1 != nil {
				log.Print("Por favor ingrese un número válido")
			}

			s.ActualizarValor(region, producto, int32(valor))

		case 4:
			scanner1 := bufio.NewScanner(os.Stdin)
			log.Print("Opcion Eliminar Producto")
			log.Print("Ingrese la región del producto: ")
			scanner1.Scan()
			region := strings.TrimSpace(scanner1.Text())
			log.Print("Ingrese el nombre del producto a eliminar: ")
			scanner1.Scan()
			producto := strings.TrimSpace(scanner1.Text())
			
			s.BorrarProducto(region, producto)

		case 5: 
			log.Print("Opción Salir")
			return

		default:
			log.Print("Opción no válida, por favor escoja un número entre 1 y 5.")
		}
	}
}

func main() {
	log.Print("El supervisor está corriendo...")

	supervisor := NuevoSupervisor()
	defer func() {
		if err := supervisor.conn.Close(); err != nil {
			log.Fatalf("Error al cerrar la conexión: %v", err)
		}
		for _, conn := range supervisor.connPool {
			conn.Close()
		}
	}()
	
	menu(supervisor)

/* 	supervisor.AgregarProducto("Noxus", "Vino", 25)
	supervisor.AgregarProducto("a", "A", 1)
	supervisor.AgregarProducto("b", "B", 2)
	supervisor.AgregarProducto("c", "C", 3) */
/* 	supervisor.RenombrarProducto("Noxus", "Vino", "Cerveza")
	supervisor.ActualizarValor("Noxus", "Cerveza", 50)
	supervisor.BorrarProducto("Noxus", "Cerveza") */
}
