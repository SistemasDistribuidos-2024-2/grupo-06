package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "broker/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto

	"google.golang.org/grpc"
)

const (
	brokerPort = ":50054" // Puerto en el que escucha el Broker
)

// **Estructura del Broker**
type Broker struct {
	pb.UnimplementedJayceBrokerServiceServer            // Servicio para Jayce
	servers                                  []string   // Lista de direcciones de los Servidores Hextech
	serversMu                                sync.Mutex // Mutex para acceso concurrente a la lista de servidores
}

// **Crear un nuevo Broker**
func NewBroker() *Broker {
	log.Print("Inicializando el Broker...")
	servers := []string{
		"localhost:50051", // Servidor Hextech 1
		"localhost:50052", // Servidor Hextech 2
		"localhost:50053", // Servidor Hextech 3
	}
	log.Printf("Servidores Hextech registrados: %v", servers)
	return &Broker{servers: servers}
}

// **ObtenerProducto: Implementación del método gRPC para obtener un producto**
func (b *Broker) ObtenerProducto(ctx context.Context, req *pb.JayceRequest) (*pb.JayceResponse, error) {
	log.Printf("Solicitud recibida: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Aquí puedes agregar la lógica para consultar los servidores Hextech y obtener el producto
	// Por ahora, vamos a simular una respuesta exitosa con un reloj vectorial ficticio

	vectorClock := &pb.VectorClock{
		Server1: 1,
		Server2: 2,
		Server3: 3,
	}

	response := &pb.JayceResponse{
		Status:      pb.ResponseStatus_OK,
		VectorClock: vectorClock,
		Message:     "Producto encontrado exitosamente",
	}

	log.Printf("Respuesta enviada: %v", response)
	return response, nil
}

func main() {
	log.Print("El Servidor Broker está iniciando...")

	// Inicializa el Broker
	broker := NewBroker()
	log.Print("Broker inicializado correctamente")

	// Configura el servidor gRPC y registra los servicios
	lis, err := net.Listen("tcp", brokerPort)
	if err != nil {
		log.Fatalf("Error al iniciar el Broker: %v", err)
	}
	log.Printf("Broker escuchando en %v\n", brokerPort)

	grpcServer := grpc.NewServer()
	pb.RegisterJayceBrokerServiceServer(grpcServer, broker)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al ejecutar el Broker: %v", err)
	}
}
