package main

import (
	"context"
	"log"
	"net"

	pb "data_node_1/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedDataNodeServiceServer
}

// AlmacenarDatos implementa el método RPC para almacenar datos de un Digimon
func (s *server) AlmacenarDatos(ctx context.Context, datos *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Datos recibidos para almacenar: ID: %d, Atributo: %s", datos.Id, datos.Atributo)
    return &pb.Confirmacion{Mensaje: "Datos almacenados correctamente"}, nil
}

// SolicitarAtributo implementa el método RPC para solicitar el atributo de un Digimon
func (s *server) SolicitarAtributo(ctx context.Context, solicitud *pb.SolicitudAtributo) (*pb.AtributoDigimon, error) {
    log.Printf("Solicitud de atributo para ID: %d", solicitud.Id)
    return &pb.AtributoDigimon{Atributo: "Vaccine"}, nil // Ejemplo de respuesta
}

func main() {
    lis, err := net.Listen("tcp", ":50054")
    if err != nil {
        log.Fatalf("Error al iniciar el listener: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterDataNodeServiceServer(grpcServer, &server{})

    log.Println("Data Node 1 corriendo en :50054")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
