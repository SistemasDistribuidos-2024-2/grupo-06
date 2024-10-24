package main

import (
	pb "caravanas/proto/grpc/proto"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
    port = ":50052"
)

type caravanServer struct {
    pb.UnimplementedCaravanServiceServer
    mu sync.Mutex
    currentDeliveries map[string]*pb.DeliveryInstruction
}

// Implementación del método para asignar entregas a las caravanas
func (s *caravanServer) AssignDelivery(ctx context.Context, instruction *pb.DeliveryInstruction) (*pb.DeliveryStatus, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    log.Printf("Asignada entrega a la caravana: %s con escolta: %s\n", instruction.TipoCaravana, instruction.WarframeEscolta)

    // Guardamos la entrega actual
    s.currentDeliveries[instruction.IdPaquete] = instruction

    // Simulación de una entrega exitosa o fallida (85% de éxito)
    success := rand.Float32() < 0.85
    var estado string
    if success {
        estado = "Entregado"
    } else {
        estado = "No Entregado"
    }

    logDeliveryStatus(instruction, estado)

    // Espera para un segundo paquete
    time.Sleep(5 * time.Second)

    return &pb.DeliveryStatus{
        IdPaquete: instruction.IdPaquete,
        Estado:    estado,
        Intentos:  instruction.Intentos, // Simulamos que se intentó una vez
    }, nil
}

// Reporte de estado de entrega (usado por caravanas para informar al sistema logístico)
func (s *caravanServer) ReportDeliveryStatus(ctx context.Context, status *pb.DeliveryStatus) (*emptypb.Empty, error) {
    log.Printf("Reporte de entrega recibido para paquete: %s, estado: %s\n", status.IdPaquete, status.Estado)

    // Aquí puedes añadir la lógica para registrar el reporte de la caravana

    return &emptypb.Empty{}, nil
}

func logDeliveryStatus(instruction *pb.DeliveryInstruction, estado string) {
    file, err := os.OpenFile("registro_entregas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Error abriendo el archivo de registro: %v", err)
    }
    defer file.Close()

    logEntry := fmt.Sprintf("Paquete: %s, Caravana: %s, Estado: %s\n", instruction.IdPaquete, instruction.TipoCaravana, estado)
    if _, err := file.WriteString(logEntry); err != nil {
        log.Fatalf("Error escribiendo en el archivo de registro: %v", err)
    }
}


func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Fallo al escuchar en el puerto %v: %v", port, err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterCaravanServiceServer(grpcServer, &caravanServer{
        currentDeliveries: make(map[string]*pb.DeliveryInstruction),
    })

    log.Printf("Servidor de caravanas corriendo en %v", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
    }
}
