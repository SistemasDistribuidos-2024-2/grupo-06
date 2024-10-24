package main

import (
	"bufio"
	"context"
	pb "continente_server/grpc/proto"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const (
    primaryNodeAddress = "localhost:50051" // Dirección del Primary Node
    rutaINPUT = "../../INPUT.txt"
    rutaDIGIMONS = "DIGIMONS.txt"
)

var PS float32 // Probabilidad de sacrificio
var TE int     // Tiempo de espera para enviar información
var TD int     // Tiempo de ataque de Diaboromon
var CD int     // Cantidad de datos necesarios para evolucionar en Omegamon
var VI int     // Vida inicial para Greymon y Garurumon

// Función para leer los valores de INPUT.txt
func leerInput() {
    file, err := os.Open(rutaINPUT)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo INPUT.txt: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.Split(linea, ",")
        if len(partes) != 5 {
            log.Fatalf("Formato incorrecto en INPUT.txt: %s", linea)
        }

        // Parsear cada valor
        psFloat, err := strconv.ParseFloat(partes[0], 32)
        if err != nil {
            log.Fatalf("Error al parsear PS: %v", err)
        }
        PS = float32(psFloat)

        TE, err = strconv.Atoi(partes[1])
        if err != nil {
            log.Fatalf("Error al parsear TE: %v", err)
        }

        TD, err = strconv.Atoi(partes[2])
        if err != nil {
            log.Fatalf("Error al parsear TD: %v", err)
        }

        CD, err = strconv.Atoi(partes[3])
        if err != nil {
            log.Fatalf("Error al parsear CD: %v", err)
        }

        VI, err = strconv.Atoi(partes[4])
        if err != nil {
            log.Fatalf("Error al parsear VI: %v", err)
        }

        log.Printf("Valores leídos de INPUT.txt -> PS: %f, TE: %d, TD: %d, CD: %d, VI: %d", PS, TE, TD, CD, VI)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error al leer INPUT.txt: %v", err)
    }
}

func enviarDatos(client pb.PrimaryNodeServiceClient, digimons []string) {
    for _, digimon := range digimons {
        partes := strings.Split(digimon, ",")
        nombre := partes[0]
        atributo := partes[1]
        sacrificado := partes[2] == "Sacrificado"

        datos := &pb.DatosDigimon{
            Nombre:      nombre,
            Atributo:    atributo,
            Sacrificado: sacrificado,
        }

        res, err := client.TransferirDatos(context.Background(), datos)
        if err != nil {
            log.Printf("Error al enviar datos al Primary Node: %v", err)
            continue
        }
        log.Printf("Confirmación del Primary Node: %s", res.Mensaje)
    }
}

func seleccionarSacrificio(digimon string) string {
    sacrificado := "No-Sacrificado"
    if rand.Float32() < PS {
        sacrificado = "Sacrificado"
    }
    return sacrificado
}

func main() {
    rand.Seed(time.Now().UnixNano())

    // Leer los valores de PS, TE, TD, CD y VI desde INPUT.txt
    leerInput()

    // Leer los datos de digimons desde el archivo DIGIMONS.txt
    file, err := os.Open(rutaDIGIMONS)
    if err != nil {
        log.Fatalf("No se pudo abrir el archivo DIGIMONS.txt: %v", err)
    }
    defer file.Close()

    conn, err := grpc.Dial(primaryNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Primary Node: %v", err)
    }
    defer conn.Close()

    client := pb.NewPrimaryNodeServiceClient(conn)

    scanner := bufio.NewScanner(file)
    var digimons []string
    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.Split(linea, ",")
        if len(partes) != 2 {
            log.Printf("Formato incorrecto: %s", linea)
            continue
        }

        nombre := partes[0]
        atributo := partes[1]
        estado := seleccionarSacrificio(linea) // Determinar si es sacrificado o no

        digimon := fmt.Sprintf("%s,%s,%s", nombre, atributo, estado)
        digimons = append(digimons, digimon)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error al leer el archivo: %v", err)
    }

    // Enviar los datos de los digimons seleccionados
    enviarDatos(client, digimons)
}
