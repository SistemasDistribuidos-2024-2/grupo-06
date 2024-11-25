package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "broker/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto
	pbsuper "broker/grpc/sup-broker"

	"google.golang.org/grpc"
)

const (
	brokerPort = ":50054" // Puerto en el que escucha el Broker
)

// **Estructura del Broker**
type Broker struct {
	pb.UnimplementedJayceBrokerServiceServer            // Servicio para Jayce
	pbsuper.UnimplementedBrokerServiceServer
	servers                                  []string   // Lista de direcciones de los Servidores Hextech
	serversMu                                sync.Mutex // Mutex para acceso concurrente a la lista de servidores
}

// **Crear un nuevo Broker**
func NewBroker() *Broker {
	log.Print("Inicializando el Broker...")
	servers := []string{
		"dist021:50051", // Servidor Hextech 1
		"dist022:50052", // Servidor Hextech 2
		"dist023:50053", // Servidor Hextech 3
	}
	log.Printf("Servidores Hextech registrados: %v", servers)
	return &Broker{servers: servers}
}

// --------------------------------------------Jayce------------------------------------------------
// **ObtenerProducto: Implementación del método gRPC para obtener un producto(Peticion de Jeyce)**
func (b *Broker) ObtenerServidor(ctx context.Context, req *pb.JayceRequest) (*pb.JayceResponse, error) {
	log.Printf("Solicitud recibida: Región: %s, Producto: %s", req.Region, req.ProductName)

	// Selecciona un servidor de la lista de servidores
	b.serversMu.Lock()
	defer b.serversMu.Unlock()
	if len(b.servers) == 0 {
		return nil, fmt.Errorf("no hay servidores disponibles")
	}
	server := b.servers[rand.Intn(len(b.servers))] // Selecciona un servidor aleatorio

	response := &pb.JayceResponse{
		Status:  pb.ResponseStatus_OK,
		Message: &server,
	}

	log.Printf("Redirigiendo al servidor: %s", server)
	return response, nil
}

func (b *Broker) GetServer(ctx context.Context, req *pbsuper.ServerRequest) (*pbsuper.ServerResponse, error){
	log.Print("Asignando servidor...")
	direccion := b.selectServerAddress()
	return &pbsuper.ServerResponse{
		ServerAddress: direccion,
	}, nil
	
}

func (b *Broker) selectServerAddress() string {
    rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3) + 1
	return b.servers[num-1]
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
	pbsuper.RegisterBrokerServiceServer(grpcServer, broker)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al ejecutar el Broker: %v", err)
	}
}
