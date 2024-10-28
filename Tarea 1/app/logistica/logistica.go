package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	ostronitasQueue  []*pb.PackageOrder
	prioritarioQueue []*pb.PackageOrder
	normalQueue      []*pb.PackageOrder
	// Mutex para manejar concurrencia en las colas
	mu sync.Mutex
}

type PaqueteSeguimiento struct {
	CodigoSeguimiento string
	IdPaquete         string
	Faccion           string
	TipoPaquete       string
	Estado            string
	IdCaravana        string
	Intentos          int
}

var seguimientoPaquetes = make(map[string]PaqueteSeguimiento)

// Implementación del método para recibir órdenes de facciones
func (s *logisticsServer) SendOrder(ctx context.Context, order *pb.PackageOrder) (*pb.OrderResponse, error) {
	log.Printf("Recibida orden de la facción: %s, paquete: %s\n", order.Faccion, order.IdPaquete)

	codigoSeguimiento := generateTrackingCode()

	// Almacenar en el registro de seguimiento
	seguimientoPaquetes[codigoSeguimiento] = PaqueteSeguimiento{
		CodigoSeguimiento: codigoSeguimiento,
		IdPaquete:         order.IdPaquete,
		Faccion:           order.Faccion,
		TipoPaquete:       order.TipoPaquete,
		Estado:            "En Cetus",
		IdCaravana:        "",
		Intentos:          0,
	}

	// Asignar la orden a la cola correspondiente
	s.mu.Lock()
	defer s.mu.Unlock()
	if order.Faccion == "Ostronitas" {
		s.ostronitasQueue = append(s.ostronitasQueue, order)
	} else if order.TipoPaquete == "Prioritario" {
		s.prioritarioQueue = append(s.prioritarioQueue, order)
	} else {
		s.normalQueue = append(s.normalQueue, order)
	}

	// Enviar paquete a RabbitMQ
	err := sendToRabbitMQ(order)
	if err != nil {
		log.Printf("Error al enviar a RabbitMQ: %v", err)
	}

	return &pb.OrderResponse{
		CodigoSeguimiento: codigoSeguimiento,
		Mensaje:           "Orden procesada exitosamente",
	}, nil
}

// failOnError logs the error message and exits the application if an error occurs.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Función para enviar un paquete a RabbitMQ
func sendToRabbitMQ(order *pb.PackageOrder) error {
	var conn *amqp.Connection
	var err error
	// Intentar conectarse a RabbitMQ con reintentos
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/") //El servidor RabbitMQ se encuentra en el mismo contenedor que el servidor de logística
		if err == nil {
			log.Printf("Servidor Rabbit MQ conectado exitosamente")
			break
		}
		log.Printf("Failed to connect to RabbitMQ, retrying in 5 seconds... (%d/10)", i+1)
		time.Sleep(5 * time.Second)
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"paquetes_entregados", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return fmt.Errorf("error declarando cola: %v", err)
	}

	// Crear el paquete en formato JSON para enviarlo a finanzas
	paquete := Paquete{
		ID:       order.IdPaquete,
		Valor:    float64(order.ValorSuministro),
		Intentos: 1,
		Estado:   "enviado",
		Servicio: order.Faccion,
	}

	body, err := json.Marshal(paquete)
	if err != nil {
		return fmt.Errorf("error serializando paquete: %v", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("error publicando en RabbitMQ: %v", err)
	}

	log.Printf("Paquete enviado a RabbitMQ: %s\n", body)
	return nil
}

// Función para calcular el número máximo de intentos
func calcularMaxIntentos(order *pb.PackageOrder) int {
	if order.Faccion == "Ostronitas" {
		return 3 // Ostronitas siempre tiene un máximo de 3 intentos
	} else if order.Faccion == "Grineer" {
		// Para los Grineer, calculamos el número de intentos en función del valor del suministro
		intentos := int(order.ValorSuministro) / 100
		if intentos < 1 {
			return 1 // Al menos 1 intento si el valor es menor a 100
		}
		return intentos
	}
	return 3 // Valor predeterminado por seguridad
}

// Implementación del método para consultar el estado de los paquetes
func (s *logisticsServer) CheckOrderStatus(ctx context.Context, req *pb.TrackingRequest) (*pb.TrackingResponse, error) {
	log.Printf("Consulta de estado para código de seguimiento: %s\n", req.CodigoSeguimiento)

	// Aquí va la lógica para verificar el estado del paquete en base al código de seguimiento
	//Este estado servira para colocarlo en la cola y enviarlo a finanzas
	//Aqui tengo que hacer una instancia del Paquete struct!

	paquete, existe := seguimientoPaquetes[req.CodigoSeguimiento]
	if !existe {
		return nil, fmt.Errorf("no se encontró el paquete con el código de seguimiento %s", req.CodigoSeguimiento)
	}

	return &pb.TrackingResponse{
		Estado:     paquete.Estado,
		IdCaravana: paquete.IdCaravana,
		Intentos:   int32(paquete.Intentos),
	}, nil
}

// Función para crear la conexión gRPC con las caravanas
func connectToCaravans() (pb.CaravanServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial("dist022:50052", grpc.WithInsecure()) // Puerto del servicio de caravanas
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewCaravanServiceClient(conn)
	return client, conn, nil
}

// Función para asignar paquetes a caravanas
func (s *logisticsServer) assignPackagesToCaravans() {
	caravanClient, conn, err := connectToCaravans()
	if err != nil {
		log.Fatalf("No se pudo conectar al servicio de caravanas: %v", err)
	}
	defer conn.Close()

	for {
		s.mu.Lock()
		var packageToAssign *pb.PackageOrder

		// Procesar primero paquetes prioritarios
		if len(s.prioritarioQueue) > 0 {
			packageToAssign = s.prioritarioQueue[0]
			s.prioritarioQueue = s.prioritarioQueue[1:]
		} else if len(s.ostronitasQueue) > 0 {
			packageToAssign = s.ostronitasQueue[0]
			s.ostronitasQueue = s.ostronitasQueue[1:]
		} else if len(s.normalQueue) > 0 {
			packageToAssign = s.normalQueue[0]
			s.normalQueue = s.normalQueue[1:]
		}

		if packageToAssign != nil {
			seguimiento := seguimientoPaquetes[packageToAssign.IdPaquete]
			intentos := seguimiento.Intentos

			// Calculamos el número máximo de intentos según la facción y el valor
			maxIntentos := calcularMaxIntentos(packageToAssign)

			for intentos < maxIntentos {
				instruction := &pb.DeliveryInstruction{
					IdPaquete:       packageToAssign.IdPaquete,
					TipoCaravana:    packageToAssign.TipoPaquete,
					WarframeEscolta: "Excalibur",
					Destino:         packageToAssign.Destino,
					TipoPaquete:     packageToAssign.TipoPaquete,
					Seguimiento:     seguimiento.CodigoSeguimiento,
					Valor:           packageToAssign.ValorSuministro,
					Intentos:        int32(intentos + 1),
				}

				log.Printf("Intentando entrega del paquete %s, intento %d de %d\n", packageToAssign.IdPaquete, intentos+1, maxIntentos)
				response, err := caravanClient.AssignDelivery(context.Background(), instruction)
				if err != nil {
					log.Printf("Error asignando el paquete a la caravana: %v", err)
					break
				}

				intentos++
				seguimiento.Intentos = intentos
				seguimientoPaquetes[packageToAssign.IdPaquete] = seguimiento

				// Si se entregó correctamente, actualizamos el estado
				if response.Estado == "Entregado" {
					log.Printf("Paquete %s entregado exitosamente\n", packageToAssign.IdPaquete)
					seguimiento.Estado = "Entregado"
					seguimientoPaquetes[packageToAssign.IdPaquete] = seguimiento
					break
				} else {
					log.Printf("Paquete %s no entregado, intento %d\n", packageToAssign.IdPaquete, intentos)
				}

				// Esperamos antes del siguiente intento
				time.Sleep(2 * time.Second)
			}

			if seguimiento.Intentos == maxIntentos && seguimiento.Estado != "Entregado" {
				log.Printf("Paquete %s no se pudo entregar después de %d intentos\n", packageToAssign.IdPaquete, maxIntentos)
				seguimiento.Estado = "No Entregado"
				seguimientoPaquetes[packageToAssign.IdPaquete] = seguimiento
			}
		}

		s.mu.Unlock()
		time.Sleep(2 * time.Second) // Simulación de tiempo entre asignaciones
	}
}

func generateTrackingCode() string {
	return "T" + time.Now().Format("20060102150405") + string(rand.Intn(1000))
}

/*
	 func startRabbitMQ() {
		var conn *amqp.Connection
		var err error
		// Intentar conectarse a RabbitMQ con reintentos
		for i := 0; i < 10; i++ {
			conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
			if err == nil {
				log.Printf("Servidor Rabbit MQ conectado exitosamente")
				break
			}
			log.Printf("Failed to connect to RabbitMQ, retrying in 5 seconds... (%d/10)", i+1)
			time.Sleep(5 * time.Second)
		}
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

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

		body, _ := json.Marshal(paquete)
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
*/
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
	/* 	go startRabbitMQ()

	   	log.Printf("RabbitMQ:Servidor de logística corriendo en %v", port)
	   	if err := grpcServer.Serve(lis); err != nil {
	   		log.Printf("A %v", port)
	   		log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
	   	} */
	log.Printf("Servidor de logística corriendo en %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
	}
}

// Añadir función para manejar errores en RabbitMQ
/* func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
} */

// Definir estructura para manejar el paquete en la cola de RabbitMQ
type Paquete struct {
	ID       string  `json:"id"`
	Valor    float64 `json:"valor"`
	Intentos int     `json:"intentos"`
	Estado   string  `json:"estado"`   // "entregado" o "no_entregado"
	Servicio string  `json:"servicio"` // "Ostronitas" o "Grineer Normal, Grineer Prioritario"
}
