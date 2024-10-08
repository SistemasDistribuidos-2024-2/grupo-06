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

// TransferirDatos implementa el método RPC para transferir datos a Data Nodes
func (s *server) TransferirDatos(ctx context.Context, datos *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Datos recibidos para transferir: ID: %d, Nombre: %s, Atributo: %s, Sacrificado: %v", datos.Id, datos.Nombre, datos.Atributo, datos.Sacrificado)
    return &pb.Confirmacion{Mensaje: "Datos transferidos correctamente"}, nil
}

// SolicitarDatos implementa el método RPC para solicitar cantidad de datos acumulados
func (s *server) SolicitarDatos(ctx context.Context, solicitud *pb.SolicitudTai) (*pb.RespuestaTai, error) {
    log.Printf("Solicitud de Nodo Tai recibida: %s", solicitud.Mensaje)
    return &pb.RespuestaTai{CantidadDatos: 100.0}, nil // Ejemplo de respuesta
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Error al iniciar el listener: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPrimaryNodeServiceServer(grpcServer, &server{})

    log.Println("Primary Node corriendo en :50052")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
