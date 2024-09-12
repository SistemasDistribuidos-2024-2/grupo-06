package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Definir estructura para manejar el paquete
type Paquete struct {
	ID       string  `json:"id"`
	Valor    float64 `json:"valor"`
	Intentos int     `json:"intentos"`
	Estado   string  `json:"estado"`   // "entregado" o "no_entregado"
	Servicio string  `json:"servicio"` // "Ostronitas" o "Grineer Normal, Grineer Prioritario"

}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
