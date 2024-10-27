package main

import (
	"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pbData "primary_node/grpc/primary-data"
	pbReg "primary_node/grpc/primary-reg"
	pbTai "primary_node/grpc/tai-primary" // Importa el proto para Tai-Primary

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
	pbReg.UnimplementedPrimaryNodeServiceServer
	pbTai.UnimplementedTaiNodeServiceServer
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
func (s *server) RecibirDatosRegional(ctx context.Context, in *pbReg.DatosCifradosDigimon) (*pbReg.Confirmacion, error) {
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

	return &pbReg.Confirmacion{Mensaje: "Datos recibidos y procesados exitosamente"}, nil
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

// Función para solicitar atributos de Digimons sacrificados a los Data Nodes y calcular datos acumulados
func (s *server) SolicitarCantidadDatos(ctx context.Context, in *pbTai.SolicitudTai) (*pbTai.RespuestaTai, error) {
	idsSacrificados, err := obtenerIDsSacrificados()
	if err != nil {
		return nil, err
	}

	// Dividir los IDs de los Digimons sacrificados en base a su asignación de Data Node
	dataNode1IDs, dataNode2IDs := dividirIDsPorDataNode(idsSacrificados)

	// Calcular la cantidad de datos acumulados desde ambos Data Nodes
	cantidadDatos1, err := calcularDatosDesdeDataNode(dataNode1Address, dataNode1IDs)
	if err != nil {
		return nil, err
	}
	cantidadDatos2, err := calcularDatosDesdeDataNode(dataNode2Address, dataNode2IDs)
	if err != nil {
		return nil, err
	}

	cantidadDatos := cantidadDatos1 + cantidadDatos2

	return &pbTai.RespuestaTai{CantidadDatos: cantidadDatos}, nil
}

// Dividir los IDs de los Digimons sacrificados según el Data Node al que pertenecen
func dividirIDsPorDataNode(ids []int32) (dataNode1IDs, dataNode2IDs []int32) {

	idsMap := make(map[int32]bool)
	for _, id := range ids {
		idsMap[id] = true
	}

	file, err := os.Open(infoFile)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo %s: %v", infoFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		id, _ := strconv.Atoi(partes[0])
		dataNode := partes[1]

		if idsMap[int32(id)] {
			if dataNode == "1" {
				dataNode1IDs = append(dataNode1IDs, int32(id))
			} else if dataNode == "2" {
				dataNode2IDs = append(dataNode2IDs, int32(id))
			}
		}
	}
	return
}

// Solicita a un Data Node los atributos de los IDs especificados y calcula la cantidad de datos
func calcularDatosDesdeDataNode(dataNodeAddress string, ids []int32) (float32, error) {
	var cantidadDatos float32

	conn, err := grpc.Dial(dataNodeAddress, grpc.WithInsecure())
	if err != nil {
		return 0, fmt.Errorf("no se pudo conectar al Data Node: %v", err)
	}
	defer conn.Close()

	client := pbData.NewDataNodeServiceClient(conn)

	for _, id := range ids {
		req := &pbData.SolicitudDatos{Id: id}
		res, err := client.ObtenerDatos(context.Background(), req)
		if err != nil {
			log.Printf("Error al obtener atributo del Data Node: %v", err)
			continue
		}

		switch res.Atributo {
		case "Vaccine":
			cantidadDatos += 3
		case "Data":
			cantidadDatos += 1.5
		case "Virus":
			cantidadDatos += 0.8
		}
	}

	return cantidadDatos, nil
}

// Obtener los IDs de los Digimons sacrificados desde INFO.txt
func obtenerIDsSacrificados() ([]int32, error) {
	file, err := os.Open(infoFile)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir INFO.txt: %v", err)
	}
	defer file.Close()

	var ids []int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		if len(partes) != 4 {
			continue
		}

		id := partes[0]
		estado := partes[3]

		// Agregar solo los IDs de los Digimons sacrificados
		if estado == "Sacrificado" {
			idInt, _ := strconv.Atoi(id)
			ids = append(ids, int32(idInt))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer INFO.txt: %v", err)
	}
	return ids, nil
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
	pbReg.RegisterPrimaryNodeServiceServer(grpcServer, &server{})
	pbTai.RegisterTaiNodeServiceServer(grpcServer, &server{})

	log.Println("Primary Node escuchando en el puerto 50051...")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()

	// Simulación del fin del programa (puedes modificarlo según el contexto)
	time.Sleep(200 * time.Second) // Espera a que terminen las operaciones (simulación)
	mostrarPorcentajeSacrificados()
}
