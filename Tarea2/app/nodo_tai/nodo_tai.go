package main

import (
	"context"
	"fmt"
	"log"
	pb "nodo_tai/grpc/proto"

	"google.golang.org/grpc"
)

const (
    primaryNodeAddress = "localhost:50051" // Primary Node address
)

func solicitarDatos(client pb.NodoTaiServiceClient) {
    solicitud := &pb.Solicitud{Mensaje: "Solicitud de datos"}
    res, err := client.SolicitarDatos(context.Background(), solicitud)
    if err != nil {
        log.Fatalf("Error al solicitar datos: %v", err)
    }
    fmt.Printf("Cantidad de datos acumulados: %f\n", res.CantidadDatos)
}

func main() {
    conn, err := grpc.Dial(primaryNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Primary Node: %v", err)
    }
    defer conn.Close()

    client := pb.NewNodoTaiServiceClient(conn)
    solicitarDatos(client)
}
