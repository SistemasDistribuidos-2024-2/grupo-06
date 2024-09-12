package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Definir la estructura Paquete que se va a recibir
type Paquete struct {
	ID       string  `json:"id"`
	Valor    float64 `json:"valor"`
	Intentos int     `json:"intentos"`
	Estado   string  `json:"estado"`   // "entregado" o "no_entregado"
	Servicio string  `json:"servicio"` // "Ostronitas", "Grineer Normal", "Grineer Prioritario"
}

// Registro individual de cada paquete procesado
type RegistroPaquete struct {
	ID       string  // Identificador del paquete
	Intentos int     // Cantidad de intentos de entrega
	Estado   string  // "entregado" o "no_entregado"
	Ganancia float64 // Ganancia o pérdida en créditos
}

// Estructura para mantener los registros en memoria
type RegistroFinanzas struct {
	Completados        []RegistroPaquete  // Paquetes entregados
	NoEntregados       []RegistroPaquete  // Paquetes no entregados
	ResumenGanancias   map[string]float64 // ID del paquete y su ganancia o pérdida
	IntentosPorPaquete map[string]int     // ID del paquete y la cantidad de intentos
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Inicializar la estructura de registros
var finanzas = RegistroFinanzas{
	Completados:        []RegistroPaquete{},
	NoEntregados:       []RegistroPaquete{},
	ResumenGanancias:   make(map[string]float64),
	IntentosPorPaquete: make(map[string]int),
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	//Declaramos la misma cola que en el productor, se declara ya que se pod´ria ejecutar antes el consumidor que el productor
	q, err := ch.QueueDeclare(
		"paquetes_entregados", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// Deserializar el mensaje JSON a la estructura Paquete
			var paquete Paquete
			err := json.Unmarshal(d.Body, &paquete)
			if err != nil {
				log.Printf("Error al deserializar el mensaje: %s", err)
				continue
			}

			// Procesar el paquete recibido según la lógica del laboratorio
			procesarPaquete(paquete)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//Cuando termine su ejecución, imprimir el balance final y los intentos por paquete
	defer imprimirBalanceFinal()
	defer imprimirIntentos()
	<-forever
}

// Actualizar los registros en memoria con cada paquete procesado
func procesarPaquete(paquete Paquete) {
	var ganancia, perdida float64

	// Actualizamos la cantidad de intentos por paquete
	finanzas.IntentosPorPaquete[paquete.ID] = paquete.Intentos

	if paquete.Estado == "entregado" {
		// Si la entrega fue exitosa
		ganancia = paquete.Valor
		if paquete.Servicio == "Grineer Prioritario" {
			ganancia *= 0.30 // El 30% del valor si es Grineer Prioritario
		}

		// Guardar la ganancia en el resumen de ganancias
		finanzas.ResumenGanancias[paquete.ID] = ganancia

		// Añadir el paquete a la lista de completados
		finanzas.Completados = append(finanzas.Completados, RegistroPaquete{
			ID:       paquete.ID,
			Intentos: paquete.Intentos,
			Estado:   "entregado",
			Ganancia: ganancia,
		})

		log.Printf("Entrega exitosa! Paquete ID: %s, Ganancia: %.2f\n", paquete.ID, ganancia)
	} else {
		// Si la entrega falló
		perdida = float64(paquete.Intentos) * 100 // Costo de 100 créditos por intento fallido

		// Guardar la pérdida en el resumen de ganancias
		finanzas.ResumenGanancias[paquete.ID] = -perdida

		// Añadir el paquete a la lista de no entregados
		finanzas.NoEntregados = append(finanzas.NoEntregados, RegistroPaquete{
			ID:       paquete.ID,
			Intentos: paquete.Intentos,
			Estado:   "no_entregado",
			Ganancia: -perdida,
		})

		log.Printf("Entrega fallida! Paquete ID: %s, Pérdida: %.2f\n", paquete.ID, perdida)
	}
}

// Imprimir el balance final en créditos
func imprimirBalanceFinal() {
	var totalGanancias float64
	for id, ganancia := range finanzas.ResumenGanancias {
		log.Printf("Paquete ID: %s, Ganancia/Pérdida: %.2f\n", id, ganancia)
		totalGanancias += ganancia
	}
	log.Printf("Balance final: %.2f créditos\n", totalGanancias)
}

// Imprimir resumen de intentos por paquete
func imprimirIntentos() {
	for id, intentos := range finanzas.IntentosPorPaquete {
		log.Printf("Paquete ID: %s, Intentos: %d\n", id, intentos)
	}
}