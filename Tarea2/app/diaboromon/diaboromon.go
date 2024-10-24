package main

import (
	"context"
	pb "diaboromon/grpc/proto"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
    nodoTaiAddress = "localhost:50052"
)

func atacarPeriodicamente(client pb.DiaboromonServiceClient, td int) {
    for {
        ataque := &pb.Ataque{Dano: 10}
        res, err := client.Atacar(context.Background(), ataque)
        if err != nil {
            log.Printf("Error al atacar Nodo Tai: %v", err)
        } else {
            fmt.Printf("Confirmación de ataque: %s\n", res.Mensaje)
        }
        time.Sleep(time.Duration(td) * time.Second)
    }
}

func main() {
    inputFile := "input.txt"
    file, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo input.txt: %v", err)
    }
    defer file.Close()

    var td int
    _, err = fmt.Fscanf(file, "%d", &td)
    if err != nil {
        log.Fatalf("Error al leer el tiempo de ataque (TD): %v", err)
    }

    conn, err := grpc.Dial(nodoTaiAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Nodo Tai: %v", err)
    }
    defer conn.Close()

    client := pb.NewDiaboromonServiceClient(conn)

    atacarPeriodicamente(client, td)
}
