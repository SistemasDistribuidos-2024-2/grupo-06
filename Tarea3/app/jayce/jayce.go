package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "jayce/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto

	"google.golang.org/grpc"
)

const (
	brokerAddress = "container_broker:50054" // Dirección y puerto del Broker
)

type Jayce struct {
	client          pb.JayceBrokerServiceClient
	lastVectorClock *pb.VectorClock // Último reloj vectorial recibido para Monotonic Reads
	consultas       []Consulta      // Almacena las consultas realizadas
}

type Consulta struct {
	Solicitud *pb.JayceRequest
	Respuesta *pb.JayceResponse
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

// ObtenerProducto envía una solicitud de consulta al Broker para obtener la cantidad de un producto en una región específica
func (j *Jayce) ObtenerProducto(region, product string) (*pb.JayceRequest, *pb.JayceResponse, error) {
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
	res, err := j.client.ObtenerProducto(ctx, req)
	if err != nil {
		log.Printf("Error al obtener el producto: %v", err)
		return req, nil, err
	}
	log.Printf("Solicitud enviada correctamente al Broker: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Verifica el estado de la respuesta
	if res.Status == pb.ResponseStatus_ERROR {
		log.Printf("Error en la consulta: %s", res.Message)
		return req, res, fmt.Errorf("error en la consulta: %s", res.Message)
	}

	// Log de la respuesta del Broker
	log.Printf("Respuesta del Broker: Status: %v, VectorClock: [%d, %d, %d], Message: %s",
		res.Status, res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Message)

	// Verifica la consistencia Monotonic Reads con el reloj vectorial
	if j.lastVectorClock != nil && !isVectorClockConsistent(j.lastVectorClock, res.VectorClock) {
		log.Println("Inconsistencia detectada: la lectura retrocede en el tiempo.")
		return req, res, fmt.Errorf("inconsistencia detectada: la lectura retrocede en el tiempo")
	}

	// Imprime los datos de la respuesta
	fmt.Printf("Producto consultado: %s en %s\n", product, region)
	fmt.Printf("Reloj vectorial: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)

	// Actualiza el último reloj vectorial para Monotonic Reads
	j.lastVectorClock = res.VectorClock

	return req, res, nil
}

// AlmacenarConsulta guarda la solicitud y la respuesta en la memoria si no hay error
func (j *Jayce) AlmacenarConsulta(req *pb.JayceRequest, res *pb.JayceResponse, err error) {
	if err == nil {
		j.consultas = append(j.consultas, Consulta{Solicitud: req, Respuesta: res, Error: err})
		log.Printf("Consulta agregada correctamente: Región: %s, Producto: %s", req.Region, req.ProductName)
	} else {
		log.Printf("No se pudo agregar la consulta, HUBO UN ERROR AL OBTENER EL PRODUCTO(FUNCION Obtener Producto)!: Región: %s, Producto: %s, Error: %v", req.Region, req.ProductName, err)
	}
}

// isVectorClockConsistent verifica si el nuevo reloj vectorial es consistente con el último (Monotonic Reads)
func isVectorClockConsistent(last, current *pb.VectorClock) bool {
	return current.Server1 >= last.Server1 && current.Server2 >= last.Server2 && current.Server3 >= last.Server3
}

func main() {
	// Inicializa a Jayce y realiza consultas
	log.Println("Iniciando servidor Jayce...")
	jayce := NewJayce()
	log.Println("Jayce ha sido inicializado")

	// Ejemplo de consultas realizadas por Jayce
	req1, res1, err1 := jayce.ObtenerProducto("Noxus", "Vino")
	jayce.AlmacenarConsulta(req1, res1, err1)

	req2, res2, err2 := jayce.ObtenerProducto("Demacia", "Espadas")
	jayce.AlmacenarConsulta(req2, res2, err2)

	req3, res3, err3 := jayce.ObtenerProducto("Piltover", "Cristales Hextech")
	jayce.AlmacenarConsulta(req3, res3, err3)

	// Imprime las consultas almacenadas
	for _, consulta := range jayce.consultas {
		log.Printf("Consulta: Región: %s, Producto: %s, Respuesta: %v, Error: %v",
			consulta.Solicitud.Region, consulta.Solicitud.ProductName, consulta.Respuesta, consulta.Error)
	}
}
