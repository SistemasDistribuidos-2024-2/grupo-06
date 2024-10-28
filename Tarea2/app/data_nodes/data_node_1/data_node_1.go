package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	pb "data_node_1/grpc/primary-data" // Reemplaza "path/to" por la ruta correcta

	"google.golang.org/grpc"
)

const (
	portDataNode1 = ":50051"        // Puerto para el Data Node 1
	portDataNode2 = "dist024:50052" // Puerto para el Data Node 2
)

var (
	mutex sync.Mutex
)

// server representa el Data Node
type server struct {
	pb.UnimplementedDataNodeServiceServer
	dataFile string
}

// GuardarDatos recibe los datos del Primary Node y los almacena en el archivo correspondiente
func (s *server) GuardarDatos(ctx context.Context, in *pb.DatosParaDataNode) (*pb.Confirmacion, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Abrir o crear el archivo
	file, err := os.OpenFile(s.dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo %s: %v", s.dataFile, err)
	}
	defer file.Close()

	// Escribir la información en el archivo en el formato: ID,Atributo
	linea := fmt.Sprintf("%d,%s\n", in.Id, in.Atributo)
	if _, err := file.WriteString(linea); err != nil {
		return nil, fmt.Errorf("error al escribir en el archivo %s: %v", s.dataFile, err)
	}

	log.Printf("Datos guardados en %s: ID=%d, Atributo=%s", s.dataFile, in.Id, in.Atributo)
	return &pb.Confirmacion{Mensaje: "Datos guardados correctamente"}, nil
}

// ObtenerDatos permite responder con el atributo de un Digimon específico
func (s *server) ObtenerDatos(ctx context.Context, in *pb.SolicitudDatos) (*pb.RespuestaDatos, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Abrir el archivo de datos
	file, err := os.Open(s.dataFile)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo %s: %v", s.dataFile, err)
	}
	defer file.Close()

	// Leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		if len(partes) < 2 {
			continue
		}

		// Comparar el ID
		id, err := strconv.Atoi(partes[0])
		if err != nil {
			continue
		}
		if int32(id) == in.Id {
			// Si coincide, devolver el atributo encontrado
			return &pb.RespuestaDatos{Atributo: partes[1]}, nil
		}
	}

	// Si no se encuentra el ID, devolver un error
	return nil, fmt.Errorf("no se encontró el Digimon con ID %d en %s", in.Id, s.dataFile)
}

func main() {
	// Determina si este nodo es el Data Node 1 o Data Node 2 y configura el archivo y puerto

	port := portDataNode1
	dataFile := "INFO_1.txt"

	// Crear el servidor gRPC
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto %s: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterDataNodeServiceServer(s, &server{dataFile: dataFile})

	log.Printf("Data Node escuchando en el puerto %s y guardando en %s...", port, dataFile)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
