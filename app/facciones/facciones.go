package main

import (
	"bufio"
	"context"
	pb "facciones/proto/grpc/proto" // Cambiar esta ruta al path de tu archivo .proto compilado
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

const (
	address   = "container_logistica:50051" //Aqui debería ir la dirrecion de logistica ejemplo: "dist021:50051"
	inputFile = "input.txt"
)

// Estructura para contener las órdenes desde el archivo
type Order struct {
	IdPaquete         string
	Faccion           string
	TipoPaquete       string
	NombreSuministro  string
	ValorSuministro   int32
	Destino           string
	CodigoSeguimiento string
}

// Función para leer órdenes desde un archivo
func readOrdersFromFile(filePath string) ([]Order, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var orders []Order
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Separa la línea por comas
		fields := strings.Split(line, ",")
		if len(fields) != 7 {
			log.Printf("Formato incorrecto en la línea: %s", line)
			continue
		}
		valorSuministro := parseInt(fields[3])
		order := Order{
			IdPaquete:         fields[0],
			Faccion:           fields[1],
			TipoPaquete:       fields[2],
			NombreSuministro:  fields[3],
			ValorSuministro:   valorSuministro,
			Destino:           fields[5],
			CodigoSeguimiento: fields[6],
		}
		orders = append(orders, order)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// Función para convertir un string a int32
func parseInt(s string) int32 {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return int32(i)
}

func sendOrder(client pb.LogisticsServiceClient, order Order) string {

	// Crear un objeto de orden gRPC
	grpcOrder := &pb.PackageOrder{
		IdPaquete:        order.IdPaquete,
		Faccion:          order.Faccion,
		TipoPaquete:      order.TipoPaquete,
		NombreSuministro: order.NombreSuministro,
		ValorSuministro:  order.ValorSuministro,
		Destino:          order.Destino,
	}

	resp, err := client.SendOrder(context.Background(), grpcOrder)
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

	// Leer las órdenes desde el archivo
	orders, err := readOrdersFromFile(inputFile)
	if err != nil {
		log.Fatalf("Error al leer las órdenes del archivo: %v", err)
	}

	// Procesar y enviar cada orden
	for _, order := range orders {
		codigoSeguimiento := sendOrder(client, order)

		// Simulación de espera antes de consultar el estado
		// time.Sleep(2 * time.Second) // Descomenta si deseas hacer una pausa

		// Consultar el estado de la orden
		checkOrderStatus(client, codigoSeguimiento)
	}
}
