package main

import (
	"context"
	"log"
	"time"

	pb "nodo_tai/grpc/tai-primary" // Asegúrate de que el paquete esté bien referenciado después de compilar el .proto

	"google.golang.org/grpc"
)

const (
	primaryNodeAddress = "localhost:50051" // Dirección del Primary Node
)

// SolicitarDatos solicita la cantidad de datos acumulados al Primary Node
func SolicitarDatos(client pb.TaiNodeServiceClient) {
	// Crear el mensaje de solicitud
	solicitud := &pb.SolicitudTai{Mensaje: "Solicito cantidad de datos acumulados"}

	// Hacer la solicitud al Primary Node
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	respuesta, err := client.SolicitarCantidadDatos(ctx, solicitud)
	if err != nil {
		log.Fatalf("Error al solicitar datos al Primary Node: %v", err)
	}

	log.Printf("Cantidad de datos acumulados de los Digimons sacrificados: %.2f", respuesta.CantidadDatos)
}

func main() {
	// Conectar al Primary Node
	conn, err := grpc.Dial(primaryNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Primary Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaiNodeServiceClient(conn)

	// Solicitar datos al Primary Node
	SolicitarDatos(client)
}
