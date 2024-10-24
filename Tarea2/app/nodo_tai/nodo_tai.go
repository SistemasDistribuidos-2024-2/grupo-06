package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "nodo_tai/grpc/proto"
	"os"

	"google.golang.org/grpc"
)

const (
    primaryNodeAddress = "localhost:50051"
)

type server struct {
    pb.UnimplementedNodoTaiServiceServer
    vida int32
}

func (s *server) SolicitarDatos(ctx context.Context, in *pb.Solicitud) (*pb.Respuesta, error) {
    log.Printf("Solicitud de datos enviada al Primary Node: %s", in.Mensaje)

    conn, err := grpc.Dial(primaryNodeAddress, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("No se pudo conectar al Primary Node: %v", err)
        return nil, err
    }
    defer conn.Close()

    client := pb.NewNodoTaiServiceClient(conn)
    res, err := client.SolicitarDatos(context.Background(), &pb.Solicitud{Mensaje: "Solicito cantidad de datos"})
    if err != nil {
        log.Fatalf("Error al solicitar datos al Primary Node: %v", err)
        return nil, err
    }

    log.Printf("Datos recibidos: %f", res.CantidadDatos)
    return &pb.Respuesta{CantidadDatos: res.CantidadDatos}, nil
}

func (s *server) RecibirAtaque(ctx context.Context, in *pb.AtaqueDiaboromon) (*pb.Respuesta, error) {
    log.Printf("Ataque recibido de Diaboromon, daño: %d", in.Dano)
    s.vida -= in.Dano
    log.Printf("Vida restante: %d", s.vida)

    if s.vida <= 0 {
        log.Println("Greymon y Garurumon han sido derrotados...")
        os.Exit(0)
    }

    return &pb.Respuesta{CantidadDatos: 0}, nil
}

func main() {
    inputFile := "input.txt"
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo input.txt: %v", err)
    }
    defer file.Close()

    var vidaInicial int32
    _, err = fmt.Fscanf(file, "%d", &vidaInicial)
    if err != nil {
        log.Fatalf("Error al leer la vida inicial: %v", err)
    }

    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Error al escuchar en el puerto 50052: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterNodoTaiServiceServer(grpcServer, &server{vida: vidaInicial})

    log.Println("Nodo Tai escuchando en el puerto 50052...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
