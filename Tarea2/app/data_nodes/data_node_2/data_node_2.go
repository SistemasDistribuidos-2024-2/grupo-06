package main

import (
	"context"
	pb "data_node_2/grpc/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedDataNodeServiceServer
}

func (s *server) GuardarDatos(ctx context.Context, in *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Datos recibidos: ID=%d, Nombre=%s, Atributo=%s, Sacrificado=%v", in.Id, in.Nombre, in.Atributo, in.Sacrificado)

    outputFile := fmt.Sprintf("INFO%d.txt", in.Id%2+1)
    file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo %s: %v", outputFile, err)
        return nil, err
    }
    defer file.Close()

    _, err = fmt.Fprintf(file, "ID=%d, Nombre=%s, Atributo=%s, Sacrificado=%v\n", in.Id, in.Nombre, in.Atributo, in.Sacrificado)
    if err != nil {
        log.Fatalf("Error al escribir los datos: %v", err)
        return nil, err
    }

    log.Printf("Datos guardados en %s", outputFile)
    return &pb.Confirmacion{Mensaje: "Datos guardados exitosamente"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50054")
    if err != nil {
        log.Fatalf("Error al escuchar en el puerto: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterDataNodeServiceServer(grpcServer, &server{})

    log.Println("Data Node escuchando en el puerto 50054...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
