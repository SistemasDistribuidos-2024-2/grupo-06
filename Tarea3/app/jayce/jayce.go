package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "jayce/grpc/jayce-broker" // Importa el paquete generado por el archivo .proto

	"google.golang.org/grpc"
)

const (
    brokerAddress = "localhost:50054" // Dirección y puerto del Broker
)

type Jayce struct {
    client       pb.JayceBrokerServiceClient
    lastVectorClock *pb.VectorClock // Último reloj vectorial recibido para Monotonic Reads
}

// NewJayce crea una nueva instancia de Jayce y conecta con el Broker
func NewJayce() *Jayce {
    conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Broker: %v", err)
    }

    client := pb.NewJayceBrokerServiceClient(conn)
    return &Jayce{client: client}
}

// ObtenerProducto envía una solicitud de consulta al Broker para obtener la cantidad de un producto en una región específica
func (j *Jayce) ObtenerProducto(region, product string) {
    // Crea la solicitud
    req := &pb.JayceRequest{
        Region:      region,
        ProductName: product,
    }

    // Define un contexto con timeout
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Envía la solicitud al Broker
    res, err := j.client.ObtenerProducto(ctx, req)
    if err != nil {
        log.Printf("Error al obtener el producto: %v", err)
        return
    }

    // Verifica el estado de la respuesta
    if res.Status == pb.ResponseStatus_ERROR {
        log.Printf("Error en la consulta: %s", res.Message)
        return
    }

    // Verifica la consistencia Monotonic Reads con el reloj vectorial
    if j.lastVectorClock != nil && !isVectorClockConsistent(j.lastVectorClock, res.VectorClock) {
        log.Println("Inconsistencia detectada: la lectura retrocede en el tiempo.")
        return
    }

    // Imprime los datos de la respuesta
    fmt.Printf("Producto consultado: %s en %s\n", product, region)
    fmt.Printf("Reloj vectorial: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)

    // Actualiza el último reloj vectorial para Monotonic Reads
    j.lastVectorClock = res.VectorClock
}

// isVectorClockConsistent verifica si el nuevo reloj vectorial es consistente con el último (Monotonic Reads)
func isVectorClockConsistent(last, current *pb.VectorClock) bool {
    return current.Server1 >= last.Server1 && current.Server2 >= last.Server2 && current.Server3 >= last.Server3
}

func main() {
    // Inicializa a Jayce y realiza consultas
    jayce := NewJayce()

    // Ejemplo de consultas realizadas por Jayce
    jayce.ObtenerProducto("Noxus", "Vino")
    jayce.ObtenerProducto("Demacia", "Espadas")
    jayce.ObtenerProducto("Piltover", "Cristales Hextech")
}
