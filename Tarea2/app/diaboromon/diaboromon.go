package main

import (
	"context"
	"log"
	"net"

	pb "diaboromon/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedDiaboromonServiceServer
}

// Atacar implementa el método RPC para atacar al Nodo Tai
func (s *server) Atacar(ctx context.Context, ataque *pb.Ataque) (*pb.Confirmacion, error) {
    log.Printf("Diaboromon ataca: Daño infligido: %d", ataque.Dano)
    return &pb.Confirmacion{Mensaje: "Ataque ejecutado correctamente"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Error al iniciar el listener: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterDiaboromonServiceServer(grpcServer, &server{})

    log.Println("Diaboromon corriendo en :50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
