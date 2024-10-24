package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	pb "primary_node/grpc/proto"

	"google.golang.org/grpc"
)

const (
    dataNode1Address = "localhost:50053"
    dataNode2Address = "localhost:50054"
)

type server struct {
    pb.UnimplementedPrimaryNodeServiceServer
}

func (s *server) TransferirDatos(ctx context.Context, in *pb.DatosDigimon) (*pb.Confirmacion, error) {
    log.Printf("Recibido: ID=%d, Nombre=%s, Atributo=%s, Sacrificado=%v", in.Id, in.Nombre, in.Atributo, in.Sacrificado)

    // Determinar a qué Data Node enviar los datos
    var dataNodeAddress string
    if in.Nombre[0] >= 'A' && in.Nombre[0] <= 'M' {
        dataNodeAddress = dataNode1Address
    } else {
        dataNodeAddress = dataNode2Address
    }

    conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("No se pudo conectar al Data Node: %v", err)
        return nil, err
    }
    defer conn.Close()

    dataNodeClient := pb.NewPrimaryNodeServiceClient(conn)
    _, err = dataNodeClient.TransferirDatos(context.Background(), &pb.DatosDigimon{
        Id:        in.Id,
        Nombre:    in.Nombre,
        Atributo:  in.Atributo,
        Sacrificado: in.Sacrificado,
    })
    if err != nil {
        log.Printf("Error al transferir los datos al Data Node: %v", err)
        return nil, err
    }

    return &pb.Confirmacion{Mensaje: "Datos transferidos exitosamente"}, nil
}

func (s *server) SolicitarDatos(ctx context.Context, in *pb.SolicitudTai) (*pb.RespuestaTai, error) {
    log.Printf("Solicitud de datos recibida del Nodo Tai: %s", in.Mensaje)

    inputFile := "acumulados.txt"
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo acumulados.txt: %v", err)
        return nil, err
    }
    defer file.Close()

    var cantidadDatos float32
    _, err = fmt.Fscanf(file, "%f", &cantidadDatos)
    if err != nil {
        log.Fatalf("Error al leer la cantidad de datos acumulados: %v", err)
        return nil, err
    }

    return &pb.RespuestaTai{CantidadDatos: cantidadDatos}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Error al escuchar en el puerto 50051: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPrimaryNodeServiceServer(grpcServer, &server{})

    log.Println("Primary Node escuchando en el puerto 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
    }
}
