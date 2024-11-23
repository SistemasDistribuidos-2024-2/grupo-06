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

	servbroker_pb "servidor_3/grpc/serv_broker" // Comunicación con el Broker
	supserv_pb "servidor_3/grpc/sup-serv"       // Comunicación con Supervisores

	"google.golang.org/grpc"
)

const (
	idServer = 3
	port = ":50053" // Puerto donde escuchará el servidor
)

// HextechServer representa un servidor Hextech
type HextechServer struct {
	supserv_pb.UnimplementedHextechServiceServer    // Servicio para Supervisores
	servbroker_pb.UnimplementedHextechServerServiceServer // Servicio para el Broker
	serverID       int                              // Identificador único del servidor
	vectorClock    [3]int32                         // Reloj vectorial del servidor
	data           map[string]map[string]int32      // Almacén de productos por región
	logMutex       sync.Mutex                       // Mutex para acceso concurrente al log
	logs           []string                         // Log para registrar operaciones
	vectorMutex    sync.Mutex                       // Mutex para acceso concurrente al reloj vectorial
}

// **NuevoServidorHextech**: Crea una instancia del servidor Hextech
func NuevoServidorHextech(id int) *HextechServer {
	return &HextechServer{
		serverID:    id,
		vectorClock: [3]int32{0, 0, 0},
		data:        make(map[string]map[string]int32),
	}
}

// **HandleRequest**: Manejo de solicitudes desde Supervisores
func (s *HextechServer) HandleRequest(ctx context.Context, req *supserv_pb.SupervisorRequest) (*supserv_pb.ServerResponse, error) {
	s.vectorMutex.Lock()
	defer s.vectorMutex.Unlock()

	region := req.Region
	product := req.ProductName
	var message string

	// Procesar la operación
	switch req.OperationType {
	case supserv_pb.OperationType_AGREGAR:
		s.AgregarProducto(region, product, *req.Value)
		message = fmt.Sprintf("Producto agregado: %s en %s con cantidad %d", product, region, req.Value)
	case supserv_pb.OperationType_RENOMBRAR:
		s.RenombrarProducto(region, product, *req.NewProductName)
		message = fmt.Sprintf("Producto renombrado: %s en %s a %s", product, region, *(req.NewProductName))
	case supserv_pb.OperationType_ACTUALIZAR:
		s.ActualizarValor(region, product, *req.Value)
		message = fmt.Sprintf("Producto actualizado: %s en %s con cantidad %d", product, region, req.Value)
	case supserv_pb.OperationType_BORRAR:
		s.BorrarProducto(region, product)
		message = fmt.Sprintf("Producto borrado: %s en %s", product, region)
	default:
		errorMessage := "Operación no reconocida"
		return &supserv_pb.ServerResponse{
			Status:  supserv_pb.ResponseStatus_ERROR,
			Message: &errorMessage,
		}, nil
	}

	// Incrementa el reloj vectorial del servidor actual
	s.vectorClock[s.serverID-1]++

	// Guarda la operación en el log
	s.logMutex.Lock()
	s.logs = append(s.logs, message)
	s.logMutex.Unlock()

	// Prepara el reloj vectorial para la respuesta
	vectorClock := &supserv_pb.VectorClock{
		Server1: s.vectorClock[0],
		Server2: s.vectorClock[1],
		Server3: s.vectorClock[2],
	}

	return &supserv_pb.ServerResponse{
		Status:      supserv_pb.ResponseStatus_OK,
		VectorClock: vectorClock,
	}, nil
}

// **GetVectorClock**: Método para devolver el reloj vectorial al Broker
func (s *HextechServer) GetVectorClock(ctx context.Context, req *servbroker_pb.ServerRequest) (*servbroker_pb.ServerResponse, error) {
	s.vectorMutex.Lock()
	defer s.vectorMutex.Unlock()

	vectorClock := &servbroker_pb.VectorClock{
		Server1: s.vectorClock[0],
		Server2: s.vectorClock[1],
		Server3: s.vectorClock[2],
	}

	log.Printf("Reloj vectorial enviado al Broker: [%d, %d, %d]", s.vectorClock[0], s.vectorClock[1], s.vectorClock[2])

	return &servbroker_pb.ServerResponse{
		ServerClock: vectorClock,
	}, nil
}

// **Funciones Auxiliares para manejar productos**
func (s *HextechServer) AgregarProducto(region, product string, value int32) {
	s.data[region] = leerArchivo(region)
	if s.data[region] == nil {
		s.data[region] = make(map[string]int32)
	}
	s.data[region][product] = value
	escribirArchivo(region, s.data[region])
}

func (s *HextechServer) RenombrarProducto(region, product, newProductName string) {
	s.data[region] = leerArchivo(region)
	if _, exists := s.data[region][product]; exists {
		s.data[region][newProductName] = s.data[region][product]
		delete(s.data[region], product)
		escribirArchivo(region, s.data[region])
	}
}

func (s *HextechServer) ActualizarValor(region, product string, newValue int32) {
	s.data[region] = leerArchivo(region)
	if _, exists := s.data[region][product]; exists {
		s.data[region][product] = newValue
		escribirArchivo(region, s.data[region])
	}
}

func (s *HextechServer) BorrarProducto(region, product string) {
	s.data[region] = leerArchivo(region)
	if _, exists := s.data[region][product]; exists {
		delete(s.data[region], product)
		escribirArchivo(region, s.data[region])
	}
}

// **Funciones para archivos**
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
			return productos // Archivo no existe aún
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

// **Ejecución del Servidor**
func (s *HextechServer) RunServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor Hextech %d: %v", s.serverID, err)
	}

	grpcServer := grpc.NewServer()
	supserv_pb.RegisterHextechServiceServer(grpcServer, s)
	servbroker_pb.RegisterHextechServerServiceServer(grpcServer, s)

	log.Printf("Servidor Hextech %d escuchando en el puerto %v\n", s.serverID, port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al ejecutar el servidor Hextech %d: %v", s.serverID, err)
	}
}

func main() {
	server := NuevoServidorHextech(idServer)
	go server.RunServer(port)

	select {} // Mantener el programa activo
}
