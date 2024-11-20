package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "supervisores_hexgate/grpc/sup-broker" // Importa el paquete generado por el archivo .proto

	"google.golang.org/grpc"
)

const (
    brokerAddress = "localhost:50051" // Dirección y puerto del Broker
)

type Supervisor struct {
    client pb.BrokerServiceClient
    conn   *grpc.ClientConn // Almacenamos la conexión para cerrarla después
}


//Registros para el modelo de consistencia Read Your Writes

//Estructura que representa un registro
type Registro struct {
    Region      string           // Región donde se encuentra el registro
    Producto    string           // Nombre del producto
    Valor       int32            // Valor (cantidad) del producto
    VectorClock *pb.VectorClock // Reloj vectorial de la última versión
    Servidor    string          // Dirección del servidor Hextech que respondió
}


//Estructura que representa un supervisor
type Supervisor struct {
    client   pb.BrokerServiceClient
    conn     *grpc.ClientConn
    registros map[string]Registro // Almacena el estado local
}

// NuevoSupervisor crea un nuevo cliente para comunicarse con el Broker
func NuevoSupervisor() *Supervisor {
    conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Broker: %v", err)
    }
    log.Print("Conexión exitosa al Broker") // Log de conexión exitosa
    client := pb.NewBrokerServiceClient(conn)
    return &Supervisor{client: client, conn: conn,registros: make(map[string]Registro),} // Asignamos la conexión
}

// EnviarSolicitud envía una solicitud al Broker
func (s *Supervisor) EnviarSolicitud(req *pb.SupervisorRequest) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Llamada gRPC al Broker
    res, err := s.client.ProcessRequest(ctx, req)
    if err != nil {
        log.Fatalf("Error al enviar la solicitud: %v", err)
    }
    log.Print("Solicitud enviada correctamente") // Log de solicitud enviada correctamente

    // Manejo de la respuesta del Broker
    if res.Status == pb.ResponseStatus_OK {
        fmt.Printf("Operación realizada exitosamente en el servidor")
        fmt.Printf("Reloj vectorial: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)
        //Actualizando el estado local
        key := fmt.Sprintf("%s-%s", req.Region, req.ProductName)
        s.registros[key] = Registro{
        Region:      req.Region,
        Producto:    req.ProductName,
        Valor:       req.Value,           // Almacena el valor de la cantidad del producto
        VectorClock: res.VectorClock,     // Almacena el reloj vectorial recibido
        Servidor:    brokerAddress, 
    } else {
        fmt.Printf("Error en la operación: %s\n", res.Message)
    }
}

// AgregarProducto agrega un producto en el registro de una región
func (s *Supervisor) AgregarProducto(region, producto string, valor int32) {
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_AGREGAR,
        Value:        valor,
    }
    s.EnviarSolicitud(req)
}

// RenombrarProducto renombra un producto en el registro de una región
func (s *Supervisor) RenombrarProducto(region, producto, nuevoNombre string) {
    req := &pb.SupervisorRequest{
        Region:        region,
        ProductName:   producto,
        OperationType: pb.OperationType_RENOMBRAR,
        NewProductName: nuevoNombre,
    }
    s.EnviarSolicitud(req)
}

// ActualizarValor actualiza el valor de un producto en el registro de una región
func (s *Supervisor) ActualizarValor(region, producto string, nuevoValor int32) {
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_ACTUALIZAR,
        Value:        nuevoValor,
    }
    s.EnviarSolicitud(req)
}

// BorrarProducto elimina un producto del registro de una región
func (s *Supervisor) BorrarProducto(region, producto string) {
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_BORRAR,
    }
    s.EnviarSolicitud(req)
}


//Validación de consistencia en el modelo Read Your Writes
func (s *Supervisor) LeerConsistente(region, producto string) {
    key := fmt.Sprintf("%s-%s", region, producto)
    registro, existe := s.registros[key]
    
    if !existe {
        log.Fatalf("No existe información local para %s en la región %s", producto, region)
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    // Realiza la solicitud de lectura
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_BORRAR, // Esto es solo un ejemplo; usa el tipo correcto para leer
    }
    res, err := s.client.ProcessRequest(ctx, req)
    if err != nil {
        log.Fatalf("Error al leer el producto: %v", err)
    }

    // Validar consistencia con el reloj vectorial local y el valor
    if esConsistente(res.VectorClock, registro.VectorClock) && res.Value == registro.Valor {
        fmt.Printf("Lectura consistente: [%d, %d, %d] - Valor: %d\n", 
            res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Value)
    } else {
        log.Print("Datos no consistentes, reintentando...")
        s.LeerConsistente(region, producto) // Reintenta con el mismo servidor
    }
}

// Compara dos relojes vectoriales
func esConsistente(v1, v2 *pb.VectorClock) bool {
    return v1.Server1 >= v2.Server1 && v1.Server2 >= v2.Server2 && v1.Server3 >= v2.Server3
}


func main() {
    supervisor := NuevoSupervisor()
    if supervisor == nil {
        log.Fatal("No se pudo crear el supervisor correctamente")
    } else {
        log.Print("Éxito al crear el supervisor")
    }
    defer func() {
        if err := supervisor.conn.Close(); err != nil { // Cierra la conexión usando conn
            log.Fatalf("Error al cerrar la conexión: %v", err)
        }
    }()

    // Ejemplos de uso de las funciones del supervisor
    supervisor.AgregarProducto("Noxus", "Vino", 25)
    supervisor.RenombrarProducto("Noxus", "Vino", "Cerveza")
    supervisor.ActualizarValor("Noxus", "Cerveza", 50)
    supervisor.BorrarProducto("Noxus", "Cerveza")
}
