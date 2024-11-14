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

// NuevoSupervisor crea un nuevo cliente para comunicarse con el Broker
func NuevoSupervisor() *Supervisor {
    conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("No se pudo conectar al Broker: %v", err)
    }

    client := pb.NewBrokerServiceClient(conn)
    return &Supervisor{client: client, conn: conn} // Asignamos la conexión
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

    // Manejo de la respuesta del Broker
    if res.Status == pb.ResponseStatus_OK {
        fmt.Printf("Operación realizada exitosamente en el servidor")
        fmt.Printf("Reloj vectorial: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)
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

func main() {
    supervisor := NuevoSupervisor()
    log.Print("Exito al crear el supervisor")
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
