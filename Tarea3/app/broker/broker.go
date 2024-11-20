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
    log.Print("Inicializando el Broker...")
    servers := make(map[int]server_pb.HextechServerServiceClient)

    // Conecta el Broker a los tres Servidores Hextech en sus puertos correspondientes
    for i := 1; i <= 3; i++ {
        address := fmt.Sprintf("localhost:%d", 50050+i)
        log.Printf("Intentando conectar al Servidor Hextech %d en %s...", i, address)
        conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
            log.Fatalf("No se pudo conectar al Servidor Hextech %d: %v", i, err)
        }
        servers[i] = server_pb.NewHextechServerServiceClient(conn)
        log.Printf("Conexión exitosa al Servidor Hextech %d", i)
    }
    log.Print("Broker inicializado correctamente")
    return &Broker{servers: servers}
}


// ProcessRequest recibe solicitudes de los Supervisores Hexgate y las redirige al Servidor Hextech correspondiente
func (b *Broker) ProcessRequest(ctx context.Context, req *supervisor_pb.SupervisorRequest) (*supervisor_pb.ServerResponse, error) {
    serverID := b.selectServerID() // ID del servidor destino
    b.serversMu.Lock()
    serverClient, exists := b.servers[serverID]
    b.serversMu.Unlock()

    if !exists {
        errorMessage := "Servidor Hextech no encontrado"
        return &supervisor_pb.ServerResponse{
            Status: supervisor_pb.ResponseStatus_ERROR,
            Message: &errorMessage,
        }, nil
    }

    // Convierte la solicitud del supervisor en el formato de los servidores Hextech
    serverReq := &server_pb.ServerRequest{
        Region:        req.Region,
        ProductName:   req.ProductName,
        OperationType: server_pb.OperationType(req.OperationType),
    }
    if req.Value != nil {
        serverReq.Value = *(req.Value)
    }
    if req.NewProductName != nil {
        serverReq.NewProductName = *(req.NewProductName)
    }
    log.Print(serverReq)
    // Envía la solicitud al servidor y espera la respuesta
    serverRes, err := serverClient.ProcessRequest(ctx, serverReq)
    if err != nil {
        errorMessage := "Error al procesar la solicitud en el Servidor Hextech"
        return &supervisor_pb.ServerResponse{
            Status:  supervisor_pb.ResponseStatus_ERROR,
            Message: &errorMessage,
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
        Value: &serverReq.Value,
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
        if isNewerVectorClockJayce(currentVectorClock, latestVectorClock) {
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

// isNewerVectorClockJayce compara dos relojes vectoriales para determinar si uno es más nuevo
func isNewerVectorClockJayce(newer, current *jayce_pb.VectorClock) bool {
    return newer.Server1 >= current.Server1 && newer.Server2 >= current.Server2 && newer.Server3 >= current.Server3
}

func isNewerVectorClockSupervisor(newer, current *supervisor_pb.VectorClock) bool {
    return newer.Server1 >= current.Server1 && newer.Server2 >= current.Server2 && newer.Server3 >= current.Server3
}

func (b *Broker) selectServerID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(3) + 1 // Genera un número entre 1 y 3
}

func main() {
    log.Print("El Servidor Broker está iniciando...")

    // Inicializa el Broker y establece los servicios
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
    jayce_pb.RegisterJayceBrokerServiceServer(grpcServer, broker)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al ejecutar el Broker: %v", err)
    }
}