package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "jayce/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto
	pbserver "jayce/grpc/jayce-server"

	"google.golang.org/grpc"
)

const (
	brokerAddress = "container_broker:50054" // Dirección y puerto del Broker
)

type Jayce struct {
	client    pb.JayceBrokerServiceClient
	clientServer pbserver.JayceServerServiceClient
	consultas []Consulta // Almacena las consultas realizadas
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
	return &Jayce{client: client}
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
		log.Printf("Error al obtener el producto: %v", err)
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

func (j *Jayce) ObtenerProducto(region, product string) (error) {
	req := &pbserver.JayceRequest{
		Region: region,
		ProductName: product,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Enviando solicitud al Broker: Región: %s, Producto: %s", req.Region, req.ProductName)
	res, err := j.clientServer.ObtenerProducto(ctx, req)
	if err != nil {
		log.Printf("Error al obtener el producto: %v", err)
		return err
	}
	log.Printf("Solicitud enviada correctamente al Broker: Región: %s, Producto: %s", req.Region, req.ProductName)

	cantidad := res.Cantidad

	vectorClock := res.VectorClock

	log.Printf("Respuesta del Servidor: Cantidad: %v, Reloj de Vectores: %s", cantidad, vectorClock)

	j.AlmacenarConsulta(req, res, err)

	return err
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

func main() {
	// Inicializa a Jayce y realiza consultas
	log.Println("Iniciando servidor Jayce...")
	jayce := NewJayce()
	log.Println("Jayce ha sido inicializado")

	// Ejemplo de consultas realizadas por Jayce
	req1, puerto1, err1 := jayce.ObtenerServidor("Noxus", "Vino")
	if err1 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err1)
	} else {
		log.Printf("Consulta: Región: %s, Producto: %s, Puerto: %s", req1.Region, req1.ProductName, puerto1)
	}

	req2, puerto2, err2 := jayce.ObtenerServidor("Demacia", "Espadas")
	if err2 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err2)
	} else {
		log.Printf("Consulta: Región: %s, Producto: %s, Puerto: %s", req2.Region, req2.ProductName, puerto2)
	}

	req3, puerto3, err3 := jayce.ObtenerServidor("Piltover", "Cristales Hextech")
	if err3 != nil {
		log.Printf("Error en la peticion de servidor(FUNCION OBTENER PRODUCTOS): %v", err3)
	} else {
		log.Printf("Consulta: Región: %s, Producto: %s, Puerto: %s", req3.Region, req3.ProductName, puerto3)
	}
	// Imprime las consultas almacenadas
	for _, consulta := range jayce.consultas {
		log.Printf("Consulta: Región: %s, Producto: %s, Respuesta: %v, Error: %v",
			consulta.Solicitud.Region, consulta.Solicitud.ProductName, consulta.Respuesta, consulta.Error)
	}
}
