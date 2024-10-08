package main

import (
	"context"
	"log"
	"time"

	pb "nodo_tai/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
    // Conexión con el Primary Node
    conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Primary Node: %v", err)
    }
    defer conn.Close()

    c := pb.NewNodoTaiServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Solicita la cantidad de datos acumulados al Primary Node
    r, err := c.SolicitarDatos(ctx, &pb.Solicitud{Mensaje: "Solicitud de datos acumulados"})
    if err != nil {
        log.Fatalf("Error al solicitar datos: %v", err)
    }
    log.Printf("Respuesta recibida: Cantidad de datos: %f", r.CantidadDatos)
}
