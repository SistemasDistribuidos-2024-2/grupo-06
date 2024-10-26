package main

import (
	"bufio"
	"context"
	pb "continente_folder/grpc/proto" // Usamos el archivo proto que definimos antes
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const (
    primaryNodeAddress = "localhost:50051" // Dirección del Primary Node (cambiar si está en otra máquina)
    rutaINPUT         = "../../INPUT.txt"
    rutaDIGIMONS      = "DIGIMONS.txt"
)

// Struct para almacenar la información de cada Digimon
type Digimon struct {
    Nombre     string
    Atributo   string
    Sacrificado string
}

var PS float32 // Probabilidad de sacrificio
var TE int     // Tiempo de espera para enviar información
var TD int     // Tiempo de ataque de Diaboromon
var CD int     // Cantidad de datos necesarios para evolucionar en Omegamon
var VI int     // Vida inicial para Greymon y Garurumon

var key = []byte("llave-de-grupo-6") // Clave de 32 bytes para AES-256

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

// Función para cifrar los datos usando AES
func encryptAES(plaintext string) string {
    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatalf("Error creando el cifrador AES: %v", err)
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        log.Fatalf("Error creando GCM: %v", err)
    }

    nonce := make([]byte, gcm.NonceSize())
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

    return hex.EncodeToString(ciphertext)
}

// Función para enviar los datos de los Digimons al Primary Node
func enviarDatos(client pb.RegionalServerServiceClient, digimons []Digimon) {
    for _, digimon := range digimons {
        // Cifrar la información antes de enviarla
        encryptedNombre := encryptAES(digimon.Nombre)
        encryptedAtributo := encryptAES(digimon.Atributo)
        encryptedEstado := encryptAES(digimon.Sacrificado)

        datos := &pb.DatosDigimon{
            Nombre:    encryptedNombre,
            Atributo:  encryptedAtributo,
            Estado:    encryptedEstado,
        }

        res, err := client.EnviarEstadoDigimon(context.Background(), datos)
        if err != nil {
            log.Printf("Error al enviar datos al Primary Node: %v", err)
            continue
        }
        log.Printf("Confirmación del Primary Node: %s", res.Mensaje)
    }
}

// Función que selecciona si un Digimon es sacrificado o no, según la probabilidad PS
func seleccionarSacrificio() string {
    if rand.Float32() < PS {
        return "Sacrificado"
    }
    return "No-Sacrificado"
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

    client := pb.NewRegionalServerServiceClient(conn)

    scanner := bufio.NewScanner(file)
    var digimons []Digimon

    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.Split(linea, ",")
        if len(partes) != 2 {
            log.Printf("Formato incorrecto: %s", linea)
            continue
        }

        // Crear una estructura Digimon y agregarla a la lista
        digimons = append(digimons, Digimon{
            Nombre:     partes[0],
            Atributo:   partes[1],
            Sacrificado: seleccionarSacrificio(),
        })
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error al leer el archivo: %v", err)
    }

    // Enviar los datos de los digimons seleccionados en lotes de 6 cada TE segundos
    ticker := time.NewTicker(time.Duration(TE) * time.Second)
    defer ticker.Stop()

    batchSize := 6
    for range ticker.C {
        if len(digimons) == 0 {
            log.Println("No quedan más Digimons para procesar.")
            break
        }

        // Seleccionar un lote de hasta 6 Digimons
        lote := digimons
        if len(digimons) > batchSize {
            lote = digimons[:batchSize]
            digimons = digimons[batchSize:]
        } else {
            digimons = []Digimon{}
        }

        // Enviar el lote al Primary Node
        enviarDatos(client, lote)
    }
}
