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

	jayceserver_pb "servidores_hextech/grpc/jayce-server"
	servbroker_pb "servidores_hextech/grpc/serv_broker" // Comunicación con el Broker
	supserv_pb "servidores_hextech/grpc/sup-serv"       // Comunicación con Supervisores

	"google.golang.org/grpc"
)


const DominantNodeID = 1 // Nodo dominante Definido de manera estática


// HextechServer representa un servidor Hextech
type HextechServer struct {
	supserv_pb.UnimplementedHextechServiceServer    // Servicio para Supervisores
	servbroker_pb.UnimplementedHextechServerServiceServer // Servicio para el Broker
	jayceserver_pb.UnimplementedJayceServerServiceServer
	serverID       int                              // Identificador único del servidor
    vectorClock    map[string][3]int32    // Mapeo región -> reloj vectorial
	data           map[string]map[string]int32      // Almacén de productos por región
	logMutex       sync.Mutex                       // Mutex para acceso concurrente al log
	logs           []string                         // Log para registrar operaciones
	vectorMutex    sync.Mutex                       // Mutex para acceso concurrente al reloj vectorial
}

// **NuevoServidorHextech**: Crea una instancia del servidor Hextech
func NuevoServidorHextech(id int) *HextechServer {
	return &HextechServer{
		serverID:    id,
        vectorClock: make(map[string][3]int32), // Inicializa el mapeo
		data:        make(map[string]map[string]int32),
	}
}

//Log de Registro
func (s *HextechServer) registrarLog(accion, regionAfectada, productoAfectado string) {
    // Formatear el mensaje del log
    message := fmt.Sprintf("%s %s %s", accion, regionAfectada, productoAfectado)

    // Guardar la operación en el log
    s.logMutex.Lock()
    s.logs = append(s.logs, message)
    s.logMutex.Unlock()

    // Escribir el log en un archivo
	file, err := os.OpenFile("/app/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)    
	if err != nil {
        fmt.Println("Error al abrir el archivo de logs:", err)
        return
    }
    defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		fmt.Println("Error al escribir en el archivo de logs:", err)
	} else {
		fmt.Println("Log escrito correctamente en el archivo de logs")
	}
}
func (s *HextechServer) HandleRequest(ctx context.Context, req *supserv_pb.SupervisorRequest) (*supserv_pb.ServerResponse, error) {
    log.Println("HandleRequest: Locking vectorMutex")
    s.vectorMutex.Lock()
    defer s.vectorMutex.Unlock()

    log.Println("HandleRequest: Checking context deadline")
    select {
    case <-ctx.Done():
        errorMessage := "Context deadline exceeded"
        log.Println("HandleRequest: Context deadline exceeded")
        return &supserv_pb.ServerResponse{
            Status:  supserv_pb.ResponseStatus_ERROR,
            Message: &errorMessage,
        }, ctx.Err()
    default:
        // Continue processing
    }

    region := req.Region
    product := req.ProductName
    log.Printf("HandleRequest: Processing operation %s for region %s and product %s", req.OperationType.String(), region, product)

    // Procesar la operación
    switch req.OperationType {
    case supserv_pb.OperationType_AGREGAR:
        log.Println("HandleRequest: Adding product")
        s.AgregarProducto(region, product, *req.Value)
    case supserv_pb.OperationType_RENOMBRAR:
        log.Println("HandleRequest: Renaming product")
        s.RenombrarProducto(region, product, *req.NewProductName)
    case supserv_pb.OperationType_ACTUALIZAR:
        log.Println("HandleRequest: Updating product value")
        s.ActualizarValor(region, product, *req.Value)
    case supserv_pb.OperationType_BORRAR:
        log.Println("HandleRequest: Deleting product")
        s.BorrarProducto(region, product)
    default:
        errorMessage := "Operación no reconocida"
        log.Println("HandleRequest: Unrecognized operation")
        return &supserv_pb.ServerResponse{
            Status:  supserv_pb.ResponseStatus_ERROR,
            Message: &errorMessage,
        }, nil
    }



//Incrementar Reloj------------------------------------------------------------------
	reloj := s.vectorClock[region]
    
    // Incrementa la dimensión correspondiente al servidor actual
    reloj[s.serverID-1]++
    
    // Asigna el reloj vectorial modificado de nuevo al mapa
    s.vectorClock[region] = reloj

    log.Printf("Nuevo valor de vector para la siguiente region %s: [%d, %d, %d]", region, reloj[0], reloj[1], reloj[2])

//-----------------------------------------------------------------------------------------------------


    // Llama a la función registrarLog
    log.Println("HandleRequest: Registering log")
    s.registrarLog(req.OperationType.String(), region, product)

    // Prepara el reloj vectorial para la respuesta
    vectorClock := &supserv_pb.VectorClock{
        Server1: reloj[0],
        Server2: reloj[1],
        Server3: reloj[2],
    }

    log.Println("HandleRequest: Returning response")
    return &supserv_pb.ServerResponse{
        Status:      supserv_pb.ResponseStatus_OK,
        VectorClock: vectorClock,
    }, nil
}
// **GetVectorClock**: Método para devolver el reloj vectorial al Broker
func (s *HextechServer) GetVectorClock(ctx context.Context, req *servbroker_pb.ServerRequest) (*servbroker_pb.ServerResponse, error) {
	s.vectorMutex.Lock()
	defer s.vectorMutex.Unlock()
	region := req.Region
	reloj := s.vectorClock[region]
	vectorClock := &servbroker_pb.VectorClock{
		Server1: reloj[0],
		Server2: reloj[1],
		Server3: reloj[2],
	}

	log.Printf("Reloj vectorial enviado al Broker: [%d, %d, %d]", reloj[0], reloj[1], reloj[2])

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
    fileName := fmt.Sprintf("/app/mercancias/%s.txt", region)
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
    fileName := fmt.Sprintf("/app/mercancias/%s.txt", region)
    productos := make(map[string]int32)

    file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return productos
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.Split(linea, " ")
        if len(partes) != 3 {
            continue
        }
        cantidad, err := strconv.Atoi(partes[2])
        if err != nil {
            continue
        }
        productos[partes[1]] = int32(cantidad)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error al leer el archivo:", err)
    }

    return productos
}

func (s *HextechServer) mergeLogs(logs [][]string) {
    // Implementar lógica de merge aquí
    // Actualizar data y vectorClock
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
	jayceserver_pb.RegisterJayceServerServiceServer(grpcServer, s)

	log.Printf("Servidor Hextech %d escuchando en el puerto %v\n", s.serverID, port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al ejecutar el servidor Hextech %d: %v", s.serverID, err)
	}
}

func (s *HextechServer) ObtenerProducto(ctx context.Context, req *jayceserver_pb.JayceRequest) (*jayceserver_pb.JayceResponse, error){
	productos := leerArchivo(req.Region)
	cantidad := int32(0)
	for p := range productos {
		if p == req.ProductName {
			cantidad = productos[p]
		}
	}

	s.vectorMutex.Lock()
	defer s.vectorMutex.Unlock()
	region := req.Region
	reloj := s.vectorClock[region]
	vectorClock := &jayceserver_pb.VectorClock{
		Server1: reloj[0],
		Server2: reloj[1],
		Server3: reloj[2],
	}

	log.Printf("Reloj vectorial enviado a Jayce: [%d, %d, %d]",  reloj[0], reloj[1], reloj[2])

	return &jayceserver_pb.JayceResponse{
		Cantidad: cantidad,
		VectorClock: vectorClock,
	}, nil
}

func main() {
    idServerStr := os.Getenv("SERVER_ID")
    idServer, err := strconv.Atoi(idServerStr)
    if err != nil {
        fmt.Println("SERVER_ID no válido")
        return
    }

    var port string
    switch idServer {
    case 1:
        port = ":50051"
    case 2:
        port = ":50052"
    case 3:
        port = ":50053"
    default:
        fmt.Println("SERVER_ID no válido")
        return
    }

	server := NuevoServidorHextech(idServer)
	go server.RunServer(port)

	if idServer == DominantNodeID {

        // Lógica para el nodo dominante
        // Recopilar logs de otros servidores y realizar merge
    }

	select {} // Mantener el programa activo
}
