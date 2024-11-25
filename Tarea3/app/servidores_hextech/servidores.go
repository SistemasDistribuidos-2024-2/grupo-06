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
	jayceserver_pb "servidores_hextech/grpc/jayce-server"
	servbroker_pb "servidores_hextech/grpc/serv_broker" // Comunicación con el Broker
	supserv_pb "servidores_hextech/grpc/sup-serv"       // Comunicación con Supervisores
    hextech_pb "servidores_hextech/grpc/hextech"

	"google.golang.org/grpc"
)


const DominantNodeID = 1 // Nodo dominante Definido de manera estática


// HextechServer representa un servidor Hextech
type HextechServer struct {
	supserv_pb.UnimplementedHextechServiceServer    // Servicio para Supervisores
	servbroker_pb.UnimplementedHextechServerServiceServer // Servicio para el Broker
	jayceserver_pb.UnimplementedJayceServerServiceServer // Servicio para Jayce
    hextech_pb.UnimplementedConsistenciaServiceServer // Servicio para otros servidores Hextech
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





//--------------------------------------------------------Consistencias Eventual------------------------------------------------------------------------------

// Función para devolver los logs cuando se soliciten
func (s *HextechServer) GetLogs(ctx context.Context, req *hextech_pb.GetLogsRequest) (*hextech_pb.GetLogsResponse, error) {
    s.logMutex.Lock()
    defer s.logMutex.Unlock()

    var allRegionLogs []*hextech_pb.RegionLogs

    for region, vector := range s.vectorClock {
        regionLogs := make(map[string]string)
        for i, logEntry := range s.logs {
            if strings.Contains(logEntry, region) {
                regionLogs[fmt.Sprintf("log%d", i+1)] = logEntry
            }
        }

        allRegionLogs = append(allRegionLogs, &hextech_pb.RegionLogs{
            Region: region,
            Logs:   regionLogs,
            Vector: &hextech_pb.VectorClock{
                Server1: vector[0],
                Server2: vector[1],
                Server3: vector[2],
            },
        })
    }

    return &hextech_pb.GetLogsResponse{
        ServerID:   int32(s.serverID),
        RegionLogs: allRegionLogs,
    }, nil
}

// Función para recopilar logs de otros servidores
func (s *HextechServer) recopilarLogs() []*hextech_pb.GetLogsResponse {
    var logsRecopilados []*hextech_pb.GetLogsResponse
    for i := 1; i <= 3; i++ {
        if i == s.serverID {
            continue
        }
        conn, err := grpc.Dial(fmt.Sprintf("dist02%d:5005%d", i, i), grpc.WithInsecure())
        if err != nil {
            log.Printf("Error al conectar con el servidor %d: %v", i, err)
            continue
        }
        defer conn.Close()

        client := hextech_pb.NewConsistenciaServiceClient(conn)
        req := &hextech_pb.GetLogsRequest{}
        resp, err := client.GetLogs(context.Background(), req)
        if err != nil {
            log.Printf("Esperando disponibilidad del servidor %d: ", i)
            continue
        }

        logsRecopilados = append(logsRecopilados, resp)
    }
    return logsRecopilados
}

func (s *HextechServer) ejecutarLogicaDominante() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        logs := s.recopilarLogs()
        s.mergeLogs(logs)
        log.Println("Propagando cambios a otros servidores de parte de Nodo dominante")
    }
}

// Función para realizar el merge de los logs
func (s *HextechServer) mergeLogs(logs []*hextech_pb.GetLogsResponse) {
    // Implementar lógica de merge aquí
    for _, logResponse := range logs {
        log.Printf("-------------------------------------------------------Aplicando cambios en el Servidor: %d", logResponse.ServerID)
        for _, regionLog := range logResponse.RegionLogs {
            // Procesar cada regionLog
            log.Printf("Región: %s, Vector: {Server1: %d, Server2: %d, Server3: %d}", 
                regionLog.Region, regionLog.Vector.Server1, regionLog.Vector.Server2, regionLog.Vector.Server3)
            
            for logKey, logValue := range regionLog.Logs {
                log.Printf("%s: %s", logKey, logValue)
                // Aplicar cambios en el archivo del nodo dominante
                s.applyChange(logKey, logValue)
            }
        }
    }

    // Actualizar los relojes de vectores
    s.updateVectorClocks(logs)
}

// Función para aplicar cambios en el archivo del nodo dominante
func (s *HextechServer) applyChange(logKey, logValue string) {
    // Implementar lógica para aplicar cambios en el archivo
    log.Printf("Aplicando cambio: %s -> %s", logKey, logValue)
}

// Función para actualizar los relojes de vectores
func (s *HextechServer) updateVectorClocks(logs []*hextech_pb.GetLogsResponse) {
    for _, logResponse := range logs {
        for _, regionLog := range logResponse.RegionLogs {
            // Obtener el vector actual
            currentVector := s.vectorClock[regionLog.Region]

            // Actualizar el vector del nodo dominante
            updatedVector := [3]int32{
                int32(max(int(currentVector[0]), int(regionLog.Vector.Server1))),
                int32(max(int(currentVector[1]), int(regionLog.Vector.Server2))),
                int32(max(int(currentVector[2]), int(regionLog.Vector.Server3))),
            }

            // Asignar el vector actualizado
            s.vectorClock[regionLog.Region] = updatedVector

            // Mostrar el nuevo vector actualizado
            log.Printf("Región: %s, Nuevo Vector: {Server1: %d, Server2: %d, Server3: %d}", 
                regionLog.Region, updatedVector[0], updatedVector[1], updatedVector[2])
        }
    }

    // Propagar los cambios a otros servidores
    s.propagateChanges()
}

    // Propagar los cambios a otros servidores
    
// Función para propagar los cambios a otros servidores
func (s *HextechServer) propagateChanges() {
    for i := 1; i <= 3; i++ {
        if i == s.serverID {
            continue
        }
        conn, err := grpc.Dial(fmt.Sprintf("dist02%d:5005%d", i, i), grpc.WithInsecure())
        if err != nil {
            log.Printf("Error al conectar con el servidor %d: %v", i, err)
            continue
        }
        defer conn.Close()

        client := hextech_pb.NewConsistenciaServiceClient(conn)

        // Crear un mapa de vectores por región
        vectorClocks := make(map[string]*hextech_pb.VectorClock)
        for region, vector := range s.vectorClock {
            vectorClocks[region] = &hextech_pb.VectorClock{
                Server1: vector[0],
                Server2: vector[1],
                Server3: vector[2],
            }
            log.Printf("Propagando cambios al servidor %d para la región %s: Nuevo Vector: {Server1: %d, Server2: %d, Server3: %d}", 
                i, region, vector[0], vector[1], vector[2])
        }

        req := &hextech_pb.UpdateRequest{
            ServerID:     int32(s.serverID),
            VectorClocks: vectorClocks,
        }
        _, err = client.Update(context.Background(), req)
        if err != nil {
            log.Printf("Error al propagar cambios al servidor %d: %v", i, err)
        }
    }
}

// Función para actualizar los vectores en las réplicas
func (s *HextechServer) Update(ctx context.Context, req *hextech_pb.UpdateRequest) (*hextech_pb.UpdateResponse, error) {
    s.vectorMutex.Lock()
    defer s.vectorMutex.Unlock()

    // Actualizar el vector del servidor para todas las regiones
    for region, vector := range req.VectorClocks {
        s.vectorClock[region] = [3]int32{
            vector.Server1,
            vector.Server2,
            vector.Server3,
        }

        log.Printf("Vector actualizado para la región %s en esta replica ,Servidor ----> %d: Nuevo Vector: {Server1: %d, Server2: %d, Server3: %d}", 
            region, s.serverID, vector.Server1, vector.Server2, vector.Server3)
    }

    return &hextech_pb.UpdateResponse{}, nil
}


// Función auxiliar para obtener el máximo de dos enteros
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
//---------------------------------------------------------Fin consistencias Eventual------------------------------------------------------------------------------




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
    s.vectorMutex.Lock()
    defer s.vectorMutex.Unlock()

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



//------------------------------------------------------Incrementar Reloj------------------------------------------------------------------
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
    hextech_pb.RegisterConsistenciaServiceServer(grpcServer, s)

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
    
    // Lógica para el nodo dominante
	if idServer == DominantNodeID {
        go server.ejecutarLogicaDominante()

    }

	select {} // Mantener el programa activo
}
