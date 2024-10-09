package main

import (
	"context"
	pb "diaboromon/grpc/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

const (
    nodoTaiAddress = "localhost:50052" // Nodo Tai address
)

func atacar(client pb.DiaboromonServiceClient) {
    ataque := &pb.Ataque{Dano: 10}
    res, err := client.Atacar(context.Background(), ataque)
    if err != nil {
        log.Fatalf("Error al atacar Nodo Tai: %v", err)
    }
    fmt.Printf("Confirmación de ataque: %s\n", res.Mensaje)
}

func main() {
    conn, err := grpc.Dial(nodoTaiAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Nodo Tai: %v", err)
    }
    defer conn.Close()

    client := pb.NewDiaboromonServiceClient(conn)
    atacar(client)
}
