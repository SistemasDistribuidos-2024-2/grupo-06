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
    brokerAddress = "localhost:50054" // Dirección y puerto del Broker
)

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
    log.Print(req.Region)
    log.Print(req.ProductName)
    log.Print(req.OperationType)
    if req.NewProductName != nil{
        log.Print(*(req.NewProductName))
    }
    

    // Llamada gRPC al Broker
    res, err := s.client.ProcessRequest(ctx, req)
    if err != nil {
        log.Fatalf("Error al enviar la solicitud: %v", err)
    }
    log.Print("Solicitud enviada correctamente") // Log de solicitud enviada correctamente

	// Manejo de la respuesta del Broker
	if res.Status == pb.ResponseStatus_OK {
		fmt.Printf("Operación realizada exitosamente en el servidor\n")
		fmt.Printf("Reloj vectorial: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)

		// Actualizando el estado local si es una operación de escritura
		if req.OperationType != pb.OperationType_BORRAR {
			key := fmt.Sprintf("%s-%s", req.Region, req.ProductName)
			s.registros[key] = Registro{
				Region:      req.Region,
				Producto:    req.ProductName,
				Valor:       *(res.Value),       // Valor actualizado del producto
				VectorClock: res.VectorClock, // Reloj vectorial recibido
				Servidor:    brokerAddress,
			}
		}
	} else {
		fmt.Printf("Error en la operación: %s\n", *(res.Message))
	}
}

// AgregarProducto agrega un producto en el registro de una región
func (s *Supervisor) AgregarProducto(region, producto string, valor int32) {
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_AGREGAR,
        Value:        &valor,
    }
    log.Print("Iniciando solicitud 1")
    s.EnviarSolicitud(req)
}

// RenombrarProducto renombra un producto en el registro de una región
func (s *Supervisor) RenombrarProducto(region, producto, nuevoNombre string) {
    req := &pb.SupervisorRequest{
        Region:        region,
        ProductName:   producto,
        OperationType: pb.OperationType_RENOMBRAR,
        NewProductName: &nuevoNombre,
    }
    s.EnviarSolicitud(req)
}

// ActualizarValor actualiza el valor de un producto en el registro de una región
func (s *Supervisor) ActualizarValor(region, producto string, nuevoValor int32) {
    req := &pb.SupervisorRequest{
        Region:       region,
        ProductName:  producto,
        OperationType: pb.OperationType_ACTUALIZAR,
        Value:        &nuevoValor,
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


// Validación de consistencia en el modelo Read Your Writes
func (s *Supervisor) LeerConsistente(region, producto string) {
    key := fmt.Sprintf("%s-%s", region, producto)
    registro, existe := s.registros[key]

    if !existe {
        log.Fatalf("No existe información local para %s en la región %s", producto, region)
    }
    
    	// Construye una solicitud de lectura
	req := &pb.SupervisorRequest{
		Region:       region,
		ProductName:  producto,
		OperationType: pb.OperationType_ACTUALIZAR, // Tipo de operación ficticio para lectura
	}

    // Llama a EnviarSolicitud para realizar la lectura y validar
    s.EnviarSolicitud(req)

    // Validación del reloj vectorial y el valor
	res := s.registros[key]
	if esConsistente(res.VectorClock, registro.VectorClock) && res.Valor == registro.Valor {
		fmt.Printf("Lectura consistente: [%d, %d, %d] - Valor: %d\n",
			res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Valor)
	} else {
		log.Print("Datos no consistentes, reintentando...")
		time.Sleep(100 * time.Millisecond) // Retraso antes de reintentar
		s.LeerConsistente(region, producto) // Reintenta
	}
/* 
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel() */

    // Realiza la solicitud de lectura consistente
    /* req := &pb.ReadRequest{
        Region:      region,
        ProductName: producto,
    } */
    /* res, err := s.client.ConsistentRead(ctx, req)
    if err != nil {
        log.Fatalf("Error al leer el producto: %v", err)
    }

    // Validar consistencia con el reloj vectorial local y el valor
    if res.Status == pb.ResponseStatus_OK {
        if esConsistente(res.VectorClock, registro.VectorClock) && res.Value == registro.Valor {
            fmt.Printf("Lectura consistente: [%d, %d, %d] - Valor: %d\n",
                res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Value)
        } else {
            log.Print("Datos no consistentes, reintentando...")
            time.Sleep(100 * time.Millisecond) // Retraso antes de reintentar
            s.LeerConsistente(region, producto) // Reintenta con el mismo servidor
        }
    } else {
        log.Printf("Error en la lectura consistente: %s\n", res.Status.String())
    } */
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
