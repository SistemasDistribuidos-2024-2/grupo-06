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
	"time"

	pb "servidores_hextech/grpc/serv-broker" // Importa el paquete generado por el archivo .proto

	"google.golang.org/grpc"
)

const (
    port = ":50052" // Puerto donde escuchará el servidor
)

type HextechServer struct {
    pb.UnimplementedHextechServerServiceServer
    serverID       int                  // Identificador único del servidor
    vectorClock    [3]int32             // Reloj vectorial para el servidor
    data           map[string]map[string]int32 // Almacén de productos por región
    logMutex       sync.Mutex           // Mutex para el acceso concurrente al log
    logs           []string             // Log para registrar operaciones
    vectorMutex    sync.Mutex           // Mutex para el acceso concurrente al reloj vectorial
}

// NuevoServidor crea una instancia del servidor Hextech
func NuevoServidorHextech(id int) *HextechServer {
    return &HextechServer{
        serverID:    id,
        vectorClock: [3]int32{0, 0, 0},
        data:        make(map[string]map[string]int32),
    }
}

// Maneja la escritura en un archivo de región
func escribirArchivo(region string, productos map[string]int32) {
	fileName := fmt.Sprintf("%s.txt", region)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error al crear archivo %s: %v", fileName, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for producto, cantidad := range productos {
		linea := fmt.Sprintf("%s %s %d\n", region, producto, cantidad)
		writer.WriteString(linea)
	}
	writer.Flush()
}

func leerArchivo(region string) map[string]int32 {
	fileName := fmt.Sprintf("%s.txt", region)
	productos := make(map[string]int32)

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return productos // Archivo no existe aún, devolvemos un mapa vacío
		}
		log.Fatalf("Error al leer archivo %s: %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, " ")
		if len(partes) != 3 {
			continue
		}
		producto := partes[1]
		cantidad, err := strconv.Atoi(partes[2])
		if err != nil {
			log.Printf("Error al parsear cantidad en archivo %s: %v", fileName, err)
			continue
		}
		productos[producto] = int32(cantidad)
	}
	return productos
}

// ProcessRequest maneja las solicitudes del Broker
func (s *HextechServer) ProcessRequest(ctx context.Context, req *pb.ServerRequest) (*pb.ServerResponse, error) {
    s.vectorMutex.Lock()
    defer s.vectorMutex.Unlock()
    log.Print(req)
    region := req.Region
    product := req.ProductName
    var message string

    switch req.OperationType {
    case pb.OperationType_AGREGAR:
        log.Print("AGREGANDO PRODUCTO")
        s.AgregarProducto(region, product, req.Value)
        message = fmt.Sprintf("Producto agregado: %s en %s con cantidad %d", product, region, req.Value)
    case pb.OperationType_RENOMBRAR:
        s.RenombrarProducto(region, product, req.NewProductName)
        message = fmt.Sprintf("Producto renombrado: %s en %s a %s", product, region, req.NewProductName)
    case pb.OperationType_ACTUALIZAR:
        s.ActualizarValor(region, product, req.Value)
        message = fmt.Sprintf("Producto actualizado: %s en %s con cantidad %d", product, region, req.Value)
    case pb.OperationType_BORRAR:
        s.BorrarProducto(region, product)
        message = fmt.Sprintf("Producto borrado: %s en %s", product, region)
    default:
        return &pb.ServerResponse{
            Status:  pb.ResponseStatus_ERROR,
            Message: "Operación no reconocida",
        }, nil
    }

    // Incrementa el reloj vectorial del servidor actual
    s.vectorClock[s.serverID-1]++

    // Guarda la operación en el log
    s.logMutex.Lock()
    s.logs = append(s.logs, message)
    s.logMutex.Unlock()

    // Prepara el reloj vectorial para la respuesta
    vectorClock := &pb.VectorClock{
        Server1: s.vectorClock[0],
        Server2: s.vectorClock[1],
        Server3: s.vectorClock[2],
    }

    return &pb.ServerResponse{
        Status:      pb.ResponseStatus_OK,
        VectorClock: vectorClock,
    }, nil
}

// AgregarProducto agrega un nuevo producto o actualiza su cantidad si ya existe
func (s *HextechServer) AgregarProducto(region, product string, value int32) {
    s.data[region] = leerArchivo(region)
    if s.data[region] == nil {
        s.data[region] = make(map[string]int32)
    }
    s.data[region][product] = value
    escribirArchivo(region, s.data[region])
}

// RenombrarProducto cambia el nombre de un producto
func (s *HextechServer) RenombrarProducto(region, product, newProductName string) {
    s.data[region] = leerArchivo(region)
    if _, exists := s.data[region][product]; exists {
        delete(s.data[region], product)
        s.data[region][newProductName] = s.data[region][product]
        escribirArchivo(region, s.data[region])
    }
}

// ActualizarValor modifica la cantidad de un producto en la región
func (s *HextechServer) ActualizarValor(region, product string, newValue int32) {
    s.data[region] = leerArchivo(region)
    if _, exists := s.data[region][product]; exists {
        s.data[region][product] = newValue
        escribirArchivo(region, s.data[region])
    }
}

// BorrarProducto elimina un producto del registro de una región
func (s *HextechServer) BorrarProducto(region, product string) {
    s.data[region] = leerArchivo(region)
    if _, exists := s.data[region][product]; exists {
        delete(s.data[region], product)
        escribirArchivo(region, s.data[region])
    }
}

// ConsistenciaEventual implementa la propagación de cambios entre servidores cada 30 segundos
func (s *HextechServer) ConsistenciaEventual() {
    for {
        time.Sleep(30 * time.Second)
        s.vectorMutex.Lock()
        s.logMutex.Lock()

        // Aquí se implementaría el envío de los logs a otros servidores y el proceso de merge.
        // Esto implica la elección de un servidor dominante y la resolución de conflictos.

        fmt.Println("Propagando cambios a otros servidores...")
        fmt.Printf("Estado actual del reloj vectorial: [%d, %d, %d]\n", s.vectorClock[0], s.vectorClock[1], s.vectorClock[2])

        // Limpia el log después de la propagación
        s.logs = []string{}
        s.logMutex.Unlock()
        s.vectorMutex.Unlock()
    }
}

func (s *HextechServer) RunServer(port string) {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Error al iniciar el servidor Hextech %d: %v", s.serverID, err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterHextechServerServiceServer(grpcServer, s)

    log.Printf("Servidor Hextech %d escuchando en el puerto %v\n", s.serverID, port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Error al ejecutar el servidor Hextech %d: %v", s.serverID, err)
    }
}

func main() {
    for i := 1; i <= 3; i++ {
        server := NuevoServidorHextech(i)
        port := fmt.Sprintf(":%d", 50050+i)
        go server.RunServer(port)
        go server.ConsistenciaEventual()
    }

    // Mantiene el main activo para evitar que el programa termine
    select {}
}
