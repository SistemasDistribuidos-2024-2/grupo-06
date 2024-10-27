package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pbData "primary_node/grpc/primary-data"
	pb "primary_node/grpc/primary-reg"

	"google.golang.org/grpc"
)

const (
	dataNode1Address = "localhost:50053" // Dirección del Data Node 1
	dataNode2Address = "localhost:50054" // Dirección del Data Node 2
	infoFile         = "INFO.txt"        // Archivo de INFO.txt
)

var key = []byte("llave-de-grupo-6") // Clave de 32 bytes para AES-256 (misma clave que en los servidores regionales)
var mutex sync.Mutex                  // Mutex para sincronización del archivo INFO.txt
var digimonID int32 = 0               // Contador para los ID de los Digimons
var totalDigimons int32 = 0           // Total de Digimons recibidos
var sacrificados int32 = 0            // Total de Digimons sacrificados

// Server estructura que implementa los servicios del Primary Node
type server struct {
	pb.UnimplementedPrimaryNodeServiceServer
}

// Función para desencriptar usando AES-GCM
func decryptAES(encryptedHex string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", fmt.Errorf("error decodificando hex: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("error creando el cifrador AES: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creando GCM: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("tamaño del nonce inválido")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("error al desencriptar: %v", err)
	}

	return string(plaintext), nil
}

// Función para procesar la solicitud de los servidores regionales
func (s *server) RecibirDatosRegional(ctx context.Context, in *pb.DatosCifradosDigimon) (*pb.Confirmacion, error) {
	// Desencriptar los datos recibidos
	nombre, err := decryptAES(in.NombreCifrado)
	if err != nil {
		return nil, fmt.Errorf("error al desencriptar nombre: %v", err)
	}
	atributo, err := decryptAES(in.AtributoCifrado)
	if err != nil {
		return nil, fmt.Errorf("error al desencriptar atributo: %v", err)
	}
	estado, err := decryptAES(in.EstadoCifrado)
	if err != nil {
		return nil, fmt.Errorf("error al desencriptar estado: %v", err)
	}

	// Determinar a qué Data Node se enviará
	var dataNodeAddress string
	if nombre[0] >= 'A' && nombre[0] <= 'M' {
		dataNodeAddress = dataNode1Address
	} else {
		dataNodeAddress = dataNode2Address
	}

	// Generar un ID único para el Digimon (incremental)
	id := generarID()

	// Contabilizar el Digimon
	totalDigimons++
	if estado == "Sacrificado" {
		sacrificados++
	}

	// Almacenar la información en INFO.txt
	err = almacenarEnInfoTxt(id, dataNodeAddress, nombre, estado)
	if err != nil {
		return nil, err
	}

	// Enviar la información al Data Node
	err = enviarADataNode(dataNodeAddress, id, atributo)
	if err != nil {
		return nil, err
	}

	return &pb.Confirmacion{Mensaje: "Datos recibidos y procesados exitosamente"}, nil
}

// Generar un ID único para el Digimon (incremental)
func generarID() int32 {
	mutex.Lock()
	defer mutex.Unlock()
	digimonID++
	return digimonID
}

// Función para almacenar los datos en el archivo INFO.txt
func almacenarEnInfoTxt(id int32, dataNodeAddress, nombre, estado string) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.OpenFile(infoFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("no se pudo abrir INFO.txt: %v", err)
	}
	defer file.Close()

	// Determinar el número del Data Node (1 o 2)
	numDataNode := 1
	if dataNodeAddress == dataNode2Address {
		numDataNode = 2
	}

	// Formato: ID,NumDataNode,Nombre,Estado
	linea := fmt.Sprintf("%d,%d,%s,%s\n", id, numDataNode, nombre, estado)
	_, err = file.WriteString(linea)
	if err != nil {
		return fmt.Errorf("no se pudo escribir en INFO.txt: %v", err)
	}
	return nil
}

// Función para enviar datos al Data Node
func enviarADataNode(dataNodeAddress string, id int32, atributo string) error {
	conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("no se pudo conectar al Data Node: %v", err)
	}
	defer conn.Close()

	client := pbData.NewDataNodeServiceClient(conn)
	datos := &pbData.DatosParaDataNode{
		Id:       id,
		Atributo: atributo,
	}
	_, err = client.GuardarDatos(context.Background(), datos)
	if err != nil {
		return fmt.Errorf("error al enviar datos al Data Node: %v", err)
	}
	return nil
}

// Función para mostrar el porcentaje de Digimons sacrificados al final del programa
func mostrarPorcentajeSacrificados() {
	if totalDigimons == 0 {
		log.Println("No se han procesado Digimons.")
		return
	}

	porcentajeSacrificados := (float64(sacrificados) / float64(totalDigimons)) * 100
	porcentajeNoSacrificados := 100 - porcentajeSacrificados

	log.Printf("Porcentaje de Digimons sacrificados: %.2f%%", porcentajeSacrificados)
	log.Printf("Porcentaje de Digimons no sacrificados: %.2f%%", porcentajeNoSacrificados)
}

// Iniciar el servidor Primary Node
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPrimaryNodeServiceServer(grpcServer, &server{})

	log.Println("Primary Node escuchando en el puerto 50051...")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()

	// Simulación del fin del programa (puedes modificarlo según el contexto)
	time.Sleep(1000 * time.Second) // Espera a que terminen las operaciones (simulación)
	mostrarPorcentajeSacrificados()
}
