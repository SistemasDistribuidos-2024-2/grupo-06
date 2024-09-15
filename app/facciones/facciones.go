package main

import (
	"context"
	pb "facciones/proto/grpc/proto" // Cambiar esta ruta al path de tu archivo .proto compilado
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051" //Aqui debería ir la dirrecion del contenedor de logistica ejemplo: "distXYZ:50051"
)

func sendOrder(client pb.LogisticsServiceClient) string {
	// Ejemplo de envío de una orden por parte de una facción (Ostronitas o Grineer)
	order := &pb.PackageOrder{
		IdPaquete:        "001",
		Faccion:          "Ostronitas",
		TipoPaquete:      "Ostronitas",
		NombreSuministro: "Antitoxinas",
		ValorSuministro:  150,
		Destino:          "Puesto A",
	}

	resp, err := client.SendOrder(context.Background(), order)
	if err != nil {
		log.Fatalf("Error al enviar la orden: %v", err)
	}
	log.Printf("Orden enviada exitosamente. Código de seguimiento: %s\n", resp.CodigoSeguimiento)
	return resp.CodigoSeguimiento
}

func checkOrderStatus(client pb.LogisticsServiceClient, codigoSeguimiento string) {
	// Ejemplo de consulta de estado de un paquete
	req := &pb.TrackingRequest{
		CodigoSeguimiento: codigoSeguimiento,
	}

	resp, err := client.CheckOrderStatus(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al consultar el estado: %v", err)
	}

	log.Printf("Estado del paquete: %s, Caravana: %s, Intentos: %d\n", resp.Estado, resp.IdCaravana, resp.Intentos)
}

func main() {
	// Conectarse al servidor de logística
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor de logística: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogisticsServiceClient(conn)

	// Enviar una orden y obtener el código de seguimiento
	codigoSeguimiento := sendOrder(client)	

	// Simulación de espera antes de consultar el estado
	time.Sleep(2 * time.Second)

	// Consultar el estado de la orden
	checkOrderStatus(client, codigoSeguimiento)
}
