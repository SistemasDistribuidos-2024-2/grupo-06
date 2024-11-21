package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	supervisor_pb "broker/grpc/sup-broker" // Protobuf para Supervisores
	servbro_pb "broker/grpc/serv-bro"     // Protobuf para Servidores Hextech
	"google.golang.org/grpc"
)

const (
	brokerPort = ":50054" // Puerto en el que escucha el Broker
)

// **Estructura del Broker**
type Broker struct {
	supervisor_pb.UnimplementedBrokerServiceServer // Servicio para Supervisores
	servers   []string                             // Lista de direcciones de los Servidores Hextech
	serversMu sync.Mutex                           // Mutex para acceso concurrente a la lista de servidores
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

// **GetRandomServer: Seleccionar servidor aleatorio**
func (b *Broker) GetRandomServer(ctx context.Context, req *supervisor_pb.ServerRequest) (*supervisor_pb.ServerResponse, error) {
	serverAddress := b.selectServer()
	log.Printf("Servidor aleatorio asignado: %s", serverAddress)

	return &supervisor_pb.ServerResponse{
		ServerAddress: serverAddress,
	}, nil
}

// **ResolveInconsistency: Elegir servidor más actualizado basado en reloj vectorial**
func (b *Broker) ResolveInconsistency(ctx context.Context, req *supervisor_pb.InconsistencyRequest) (*supervisor_pb.ServerResponse, error) {
	b.serversMu.Lock()
	defer b.serversMu.Unlock()

	var bestServer string
	var bestClock *supervisor_pb.VectorClock

	// Iterar sobre los servidores y comparar los relojes
	for _, server := range b.servers {
		serverClock, err := b.getServerClock(server, req.SupervisorClock)
		if err != nil {
			log.Printf("Error al obtener reloj vectorial del servidor %s: %v", server, err)
			continue
		}

		// Comparar el reloj del servidor con el mejor encontrado hasta ahora
		if isNewerVectorClock(serverClock, bestClock) {
			bestServer = server
			bestClock = serverClock
		}
	}

	if bestServer == "" {
		return nil, fmt.Errorf("No se encontró un servidor más actualizado")
	}

	log.Printf("Servidor más actualizado encontrado: %s", bestServer)
	return &supervisor_pb.ServerResponse{
		ServerAddress: bestServer,
	}, nil
}

// **getServerClock: Obtener reloj vectorial de un servidor Hextech**
func (b *Broker) getServerClock(serverAddress string, clientClock *supervisor_pb.VectorClock) (*supervisor_pb.VectorClock, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("No se pudo conectar al servidor %s: %v", serverAddress, err)
	}
	defer conn.Close()

	client := servbro_pb.NewHextechServerServiceClient(conn)
	req := &servbro_pb.ServerRequest{
		KnownVectorClock: &servbro_pb.VectorClock{
			Server1: clientClock.Server1,
			Server2: clientClock.Server2,
			Server3: clientClock.Server3,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetVectorClock(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener reloj vectorial del servidor %s: %v", serverAddress, err)
	}

	return res.VectorClock, nil
}

// **isNewerVectorClock: Comparar relojes vectoriales**
func isNewerVectorClock(v1, v2 *supervisor_pb.VectorClock) bool {
	if v1 == nil {
		return false
	}
	if v2 == nil {
		return true
	}
	return v1.Server1 >= v2.Server1 && v1.Server2 >= v2.Server2 && v1.Server3 >= v2.Server3
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
	supervisor_pb.RegisterBrokerServiceServer(grpcServer, broker)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al ejecutar el Broker: %v", err)
	}
}
