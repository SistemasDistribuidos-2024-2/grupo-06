package main

import (
	"context"
	"log"
	pb "logistica/proto/grpc/proto"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
    port = ":50051"
)

type logisticsServer struct {
    pb.UnimplementedLogisticsServiceServer
    // Colas para los paquetes
    ostronitasQueue  []pb.PackageOrder
    prioritarioQueue []pb.PackageOrder
    normalQueue      []pb.PackageOrder
    // Mutex para manejar concurrencia en las colas
    mu sync.Mutex
}

// Implementación del método para recibir órdenes de facciones
func (s *logisticsServer) SendOrder(ctx context.Context, order *pb.PackageOrder) (*pb.OrderResponse, error) {
    log.Printf("Recibida orden de la facción: %s, paquete: %s\n", order.Faccion, order.IdPaquete)

    // Asignar la orden a la cola correspondiente
    s.mu.Lock()
    defer s.mu.Unlock()
    if order.Faccion == "Ostronitas" {
        s.ostronitasQueue = append(s.ostronitasQueue, *order)
    } else if order.TipoPaquete == "Prioritario" {
        s.prioritarioQueue = append(s.prioritarioQueue, *order)
    } else {
        s.normalQueue = append(s.normalQueue, *order)
    }

    codigoSeguimiento := generateTrackingCode()

    return &pb.OrderResponse{
        CodigoSeguimiento: codigoSeguimiento,
        Mensaje:           "Orden procesada exitosamente",
    }, nil
}

// Implementación del método para consultar el estado de los paquetes
func (s *logisticsServer) CheckOrderStatus(ctx context.Context, req *pb.TrackingRequest) (*pb.TrackingResponse, error) {
    log.Printf("Consulta de estado para código de seguimiento: %s\n", req.CodigoSeguimiento)

    // Aquí va la lógica para verificar el estado del paquete en base al código de seguimiento
    return &pb.TrackingResponse{
        Estado:      "En camino",
        IdCaravana:  "Caravana 1",
        Intentos:    1,
    }, nil
}

// Función para asignar paquetes a caravanas
func (s *logisticsServer) assignPackagesToCaravans() {
    for {
        s.mu.Lock()
        if len(s.ostronitasQueue) > 0 {
            log.Println("Asignando paquete de ostronitas a una caravana...")
            // Aquí enviaríamos el paquete a una caravana a través del gRPC de Caravanas
            // ...
            s.ostronitasQueue = s.ostronitasQueue[1:]
        } else if len(s.prioritarioQueue) > 0 {
            log.Println("Asignando paquete prioritario a una caravana...")
            // Aquí enviaríamos el paquete prioritario a una caravana
            // ...
            s.prioritarioQueue = s.prioritarioQueue[1:]
        } else if len(s.normalQueue) > 0 {
            log.Println("Asignando paquete normal a una caravana...")
            // Aquí enviaríamos el paquete normal a una caravana
            // ...
            s.normalQueue = s.normalQueue[1:]
        }
        s.mu.Unlock()
        time.Sleep(2 * time.Second) // Simulación de tiempo de espera entre asignaciones
    }
}

func generateTrackingCode() string {
    return "T" + time.Now().Format("20060102150405") + string(rand.Intn(1000))
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Fallo al escuchar en el puerto %v: %v", port, err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterLogisticsServiceServer(grpcServer, &logisticsServer{})

    // Iniciar la asignación de paquetes a caravanas
    srv := &logisticsServer{}
    go srv.assignPackagesToCaravans()

    log.Printf("Servidor de logística corriendo en %v", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
    }
}

// Espacio para la futura implementación de la comunicación con el sistema de finanzas
