package main

import (
	"context"
	"encoding/json"
	"log"
	pb "logistica/proto/grpc/proto"
	"math/rand"
	"net"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type logisticsServer struct {
	pb.UnimplementedLogisticsServiceServer
	// Colas para los paquetes
	ostronitasQueue  []pb.PackageOrder
	prioritarioQueue []pb.PackageOrder
	normalQueue      []pb.PackageOrder
	// Mutex para manejar concurrencia en las colas
	mu sync.Mutex
}

// Implementación del método para recibir órdenes de facciones
func (s *logisticsServer) SendOrder(ctx context.Context, order *pb.PackageOrder) (*pb.OrderResponse, error) {
	log.Printf("Recibida orden de la facción: %s, paquete: %s\n", order.Faccion, order.IdPaquete)

	// Asignar la orden a la cola correspondiente
	s.mu.Lock()
	defer s.mu.Unlock()
	if order.Faccion == "Ostronitas" {
		s.ostronitasQueue = append(s.ostronitasQueue, *order)
	} else if order.TipoPaquete == "Prioritario" {
		s.prioritarioQueue = append(s.prioritarioQueue, *order)
	} else {
		s.normalQueue = append(s.normalQueue, *order)
	}

	codigoSeguimiento := generateTrackingCode()

	return &pb.OrderResponse{
		CodigoSeguimiento: codigoSeguimiento,
		Mensaje:           "Orden procesada exitosamente",
	}, nil
}

// Implementación del método para consultar el estado de los paquetes
func (s *logisticsServer) CheckOrderStatus(ctx context.Context, req *pb.TrackingRequest) (*pb.TrackingResponse, error) {
	log.Printf("Consulta de estado para código de seguimiento: %s\n", req.CodigoSeguimiento)

	// Aquí va la lógica para verificar el estado del paquete en base al código de seguimiento
	//Este estado servira para colocarlo en la cola y enviarlo a finanzas
	//Aqui tengo que hacer una instancia del Paquete struct!

	return &pb.TrackingResponse{
		Estado:     "En camino",
		IdCaravana: "Caravana 1",
		Intentos:   1,
	}, nil
}

// Función para asignar paquetes a caravanas
func (s *logisticsServer) assignPackagesToCaravans() {
	for {
		s.mu.Lock()
		if len(s.ostronitasQueue) > 0 {
			log.Println("Asignando paquete de ostronitas a una caravana...")
			// Aquí enviaríamos el paquete a una caravana a través del gRPC de Caravanas
			// ...
			s.ostronitasQueue = s.ostronitasQueue[1:]
		} else if len(s.prioritarioQueue) > 0 {
			log.Println("Asignando paquete prioritario a una caravana...")
			// Aquí enviaríamos el paquete prioritario a una caravana
			// ...
			s.prioritarioQueue = s.prioritarioQueue[1:]
		} else if len(s.normalQueue) > 0 {
			log.Println("Asignando paquete normal a una caravana...")
			// Aquí enviaríamos el paquete normal a una caravana
			// ...
			s.normalQueue = s.normalQueue[1:]
		}
		s.mu.Unlock()
		time.Sleep(2 * time.Second) // Simulación de tiempo de espera entre asignaciones
	}
}

func generateTrackingCode() string {
	return "T" + time.Now().Format("20060102150405") + string(rand.Intn(1000))
}

func startRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	log.Printf("A %v", port)

	q, err := ch.QueueDeclare(
		"paquetes_entregados", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paquete := Paquete{
		ID:       "12345",
		Valor:    100.50,
		Intentos: 1,
		Estado:   "no_entregado",
		Servicio: "Ostronitas",
	}

	body, err := json.Marshal(paquete)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func main() {
	//----------------------------------------------------Servidor gRPC----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al escuchar en el puerto %v: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLogisticsServiceServer(grpcServer, &logisticsServer{})

	// Iniciar la asignación de paquetes a caravanas
	srv := &logisticsServer{}
	go srv.assignPackagesToCaravans()

	// Iniciar RabbitMQ en una goroutine
	go startRabbitMQ()

	log.Printf("Servidor de logística corriendo en %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("A %v", port)
		log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
	}
}

// Añadir función para manejar errores en RabbitMQ
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Definir estructura para manejar el paquete en la cola de RabbitMQ
type Paquete struct {
	ID       string  `json:"id"`
	Valor    float64 `json:"valor"`
	Intentos int     `json:"intentos"`
	Estado   string  `json:"estado"`   // "entregado" o "no_entregado"
	Servicio string  `json:"servicio"` // "Ostronitas" o "Grineer Normal, Grineer Prioritario"
}
