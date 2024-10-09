package main

import (
	"context"
	pb "data_node_1/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedDataNodeServiceServer
}

func (s *server) AlmacenarDatos(ctx context.Context, datos *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Datos almacenados: ID %d, Atributo %s", datos.Id, datos.Atributo)
    return &pb.Confirmacion{Mensaje: "Datos almacenados correctamente"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Fallo al escuchar: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterDataNodeServiceServer(s, &server{})
    log.Printf("Data Node escuchando en :50053")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Fallo al servir: %v", err)
    }
}
