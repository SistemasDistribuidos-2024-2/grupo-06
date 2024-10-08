package main

import (
	"context"
	"log"
	"net"

	pb "continente_server/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedServidorRegionalServiceServer
}

// EnviarEstado implementa el método RPC para enviar el estado de un Digimon
func (s *server) EnviarEstado(ctx context.Context, digimon *pb.EstadoDigimon) (*pb.Confirmacion, error) {
    log.Printf("Recibido Digimon: Nombre: %s, Atributo: %s, Sacrificado: %v", digimon.Nombre, digimon.Atributo, digimon.Sacrificado)
    return &pb.Confirmacion{Mensaje: "Estado recibido correctamente"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Error al iniciar el listener: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterServidorRegionalServiceServer(grpcServer, &server{})

    log.Println("Continente Server corriendo en :50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}