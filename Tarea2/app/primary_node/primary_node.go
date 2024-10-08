package main

import (
	"context"
	"log"
	"net"
	pb "primary_node/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedPrimaryNodeServiceServer
}

func (s *server) TransferirDatos(ctx context.Context, datos *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Recibido datos de Digimon: %v", datos.Nombre)
    return &pb.Confirmacion{Mensaje: "Datos recibidos y transferidos"}, nil
}

func (s *server) SolicitarDatos(ctx context.Context, solicitud *pb.SolicitudTai) (*pb.RespuestaTai, error) {
    return &pb.RespuestaTai{CantidadDatos: 100.0}, nil // Datos de prueba
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Fallo al escuchar: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterPrimaryNodeServiceServer(s, &server{})
    log.Printf("Primary Node escuchando en :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Fallo al servir: %v", err)
    }
}
