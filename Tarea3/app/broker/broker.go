package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	jayce_pb "broker/grpc/jayce-broker"    // Protobuf para Jayce
	server_pb "broker/grpc/serv-broker"    // Protobuf para Servidores Hextech
	supervisor_pb "broker/grpc/sup-broker" // Protobuf para Supervisores Hexgate

	"google.golang.org/grpc"
)

const (
    brokerPort = ":50054" // Puerto en el que escucha el Broker
)

type Broker struct {
    supervisor_pb.UnimplementedBrokerServiceServer  // Servicio para Supervisores Hexgate
    jayce_pb.UnimplementedJayceBrokerServiceServer  // Servicio para Jayce

    servers   map[int]server_pb.HextechServerServiceClient // Mapa de clientes de los Servidores Hextech por ID
    serversMu sync.Mutex                                  // Mutex para el acceso concurrente al mapa de servidores
}

// NewBroker crea una instancia del Broker y establece conexiones con los servidores Hextech
func NewBroker() *Broker {
    servers := make(map[int]server_pb.HextechServerServiceClient)

    // Conecta el Broker a los tres Servidores Hextech en sus puertos correspondientes
    for i := 1; i <= 3; i++ {
        address := fmt.Sprintf("localhost:%d", 50050+i)
        conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
            log.Fatalf("No se pudo conectar al Servidor Hextech %d: %v", i, err)
        }
        servers[i] = server_pb.NewHextechServerServiceClient(conn)
    }
    return &Broker{servers: servers}
}

// ProcessRequest recibe solicitudes de los Supervisores Hexgate y las redirige al Servidor Hextech correspondiente
func (b *Broker) ProcessRequest(ctx context.Context, req *supervisor_pb.SupervisorRequest) (*supervisor_pb.ServerResponse, error) {
    serverID := b.selectServerID() // ID del servidor destino
    b.serversMu.Lock()
    serverClient, exists := b.servers[serverID]
    b.serversMu.Unlock()

    if !exists {
        return &supervisor_pb.ServerResponse{
            Status:  supervisor_pb.ResponseStatus_ERROR,
            Message: "Servidor Hextech no encontrado",
        }, nil
    }

    // Convierte la solicitud del supervisor en el formato de los servidores Hextech
    serverReq := &server_pb.ServerRequest{
        Region:        req.Region,
        ProductName:   req.ProductName,
        OperationType: server_pb.OperationType(req.OperationType),
        Value:         req.Value,
        NewProductName: req.NewProductName,
    }

    // Envía la solicitud al servidor y espera la respuesta
    serverRes, err := serverClient.ProcessRequest(ctx, serverReq)
    if err != nil {
        return &supervisor_pb.ServerResponse{
            Status:  supervisor_pb.ResponseStatus_ERROR,
            Message: "Error al procesar la solicitud en el Servidor Hextech",
        }, nil
    }

    convertedVectorClock := &supervisor_pb.VectorClock{
        Server1: serverRes.VectorClock.Server1,
        Server2: serverRes.VectorClock.Server2,
        Server3: serverRes.VectorClock.Server3,
    }

    // Responde al supervisor con los datos de la operación y el reloj vectorial actualizado
    return &supervisor_pb.ServerResponse{
        Status:      supervisor_pb.ResponseStatus_OK,
        VectorClock: convertedVectorClock,
    }, nil
}

// ObtenerProducto maneja las solicitudes de Jayce para consultar la cantidad de un producto en una región
func (b *Broker) ObtenerProducto(ctx context.Context, req *jayce_pb.JayceRequest) (*jayce_pb.JayceResponse, error) {
    latestVectorClock := &jayce_pb.VectorClock{Server1: 0, Server2: 0, Server3: 0}

    // Consulta cada servidor para obtener la cantidad y el reloj vectorial más actualizado
    for _, client := range b.servers {
        serverReq := &server_pb.ServerRequest{
            Region:      req.Region,
            ProductName: req.ProductName,
        }

        // Enviar la solicitud al servidor y obtener la respuesta
        serverRes, err := client.ProcessRequest(ctx, serverReq)
        if err != nil || serverRes.Status == server_pb.ResponseStatus_ERROR {
            continue // Ignora errores al consultar servidores
        }

        currentVectorClock := &jayce_pb.VectorClock{
            Server1: serverRes.VectorClock.Server1,
            Server2: serverRes.VectorClock.Server2,
            Server3: serverRes.VectorClock.Server3,
        }

        // Compara el reloj vectorial para encontrar la versión más reciente
        if isNewerVectorClock(currentVectorClock, latestVectorClock) {
            latestVectorClock = currentVectorClock
        }
    }

    // Si no se encontró el producto, devuelve un error
    if latestVectorClock == nil {
        return &jayce_pb.JayceResponse{
            Status:  jayce_pb.ResponseStatus_ERROR,
            Message: "Producto o región no encontrados",
        }, nil
    }

    // Responde a Jayce con la cantidad y el reloj vectorial más actualizado
    return &jayce_pb.JayceResponse{
        Status:      jayce_pb.ResponseStatus_OK,
        VectorClock: latestVectorClock,
    }, nil
}

// isNewerVectorClock compara dos relojes vectoriales para determinar si uno es más nuevo
func isNewerVectorClock(newer, current *jayce_pb.VectorClock) bool {
    return newer.Server1 >= current.Server1 && newer.Server2 >= current.Server2 && newer.Server3 >= current.Server3
}

func (b *Broker) selectServerID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(3) + 1 // Genera un número entre 1 y 3
}

func main() {
    // Inicializa el Broker y establece los servicios
    broker := NewBroker()

    // Configura el servidor gRPC y registra los servicios
    lis, err := net.Listen("tcp", brokerPort)
    if err != nil {
        log.Fatalf("Error al iniciar el Broker: %v", err)
    }

    grpcServer := grpc.NewServer()
    supervisor_pb.RegisterBrokerServiceServer(grpcServer, broker)
    jayce_pb.RegisterJayceBrokerServiceServer(grpcServer, broker)

    log.Printf("Broker escuchando en %v\n", brokerPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al ejecutar el Broker: %v", err)
    }
}
