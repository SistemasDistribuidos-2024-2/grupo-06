package main

import (
	"bufio"
	"context"
	pb "continente_server/grpc/proto"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

const (
    primaryNodeAddress = "localhost:50051" // Dirección del Primary Node
)

func enviarDatos(client pb.PrimaryNodeServiceClient, id int32, nombre, atributo string, sacrificado bool) {
    datos := &pb.DatosDigimon{
        Id:         id,
        Nombre:     nombre,
        Atributo:   atributo,
        Sacrificado: sacrificado,
    }
    
    res, err := client.TransferirDatos(context.Background(), datos)
    if err != nil {
        log.Printf("Error al enviar datos al Primary Node: %v", err)
        return
    }
    log.Printf("Confirmación del Primary Node: %s", res.Mensaje)
}

func main() {
    // Leer los datos de digimons.txt
    file, err := os.Open("../DIGIMONS.txt")
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
    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.Split(linea, ",")
        if len(partes) != 4 {
            log.Printf("Formato de línea incorrecto: %s", linea)
            continue
        }

        // Parsear los datos de la línea
        id := int32(0)
        fmt.Sscanf(partes[0], "%d", &id)
        nombre := partes[1]
        atributo := partes[2]
        sacrificado := partes[3] == "true"

        // Enviar los datos al Primary Node
        enviarDatos(client, id, nombre, atributo, sacrificado)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error al leer el archivo: %v", err)
    }
}
