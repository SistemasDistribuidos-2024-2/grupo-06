package main

import (
	"context"
	"fmt"
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

// --------------------------------------------Jayce------------------------------------------------
// **ObtenerProducto: Implementación del método gRPC para obtener un producto(Peticion de Jeyce)**
func (b *Broker) ObtenerProducto(ctx context.Context, req *pb.JayceRequest) (*pb.JayceResponse, error) {
	log.Printf("Solicitud recibida: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Selecciona un servidor de la lista de servidores
	b.serversMu.Lock()
	defer b.serversMu.Unlock()
	if len(b.servers) == 0 {
		return nil, fmt.Errorf("no hay servidores disponibles")
	}
	server := b.servers[0] // Aquí puedes implementar una lógica de balanceo de carga más avanzada

	response := &pb.JayceResponse{
		Status:  pb.ResponseStatus_OK,
		Message: server,
	}

	log.Printf("Redirigiendo al servidor: %s", server)
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
