package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "supervisor_1/grpc/sup-broker"  // Proto para comunicarse con el Broker
	hexpb "supervisor_1/grpc/sup-serv" // Proto para comunicarse con los Servidores Hextech

	"google.golang.org/grpc"
)

const (
	brokerAddress = "localhost:50054" // Dirección del Broker
)

// **Estructura de Registro**
type Registro struct {
	Region      string           // Región donde se encuentra el registro
	Producto    string           // Nombre del producto
	Valor       int32            // Valor (cantidad) del producto
	VectorClock *hexpb.VectorClock // Reloj vectorial de la última versión
	Servidor    string           // Dirección del servidor Hextech que respondió
}

// **Estructura del Supervisor**
type Supervisor struct {
	client    pb.BrokerServiceClient // Cliente para comunicarse con el Broker
	conn      *grpc.ClientConn       // Conexión al Broker
	registros map[string]Registro    // Almacena el estado local
}

// **Crear un nuevo Supervisor**
func NuevoSupervisor() *Supervisor {
	log.Print("Intentando conectar al Broker...")
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Broker: %v", err)
	}
	log.Print("Conexión exitosa al Broker")
	client := pb.NewBrokerServiceClient(conn)
	return &Supervisor{client: client, conn: conn, registros: make(map[string]Registro)}
}

// **Obtener dirección del servidor Hextech desde el Broker**
func (s *Supervisor) ObtenerServidor() string {
	req := &pb.ServerRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := s.client.GetServer(ctx, req)
	if err != nil {
		log.Fatalf("Error al obtener servidor del Broker: %v", err)
	}
	log.Printf("Servidor asignado: %s", res.ServerAddress)
	return res.ServerAddress
}

// **Comunicación directa con el Servidor Hextech**
func (s *Supervisor) EnviarSolicitudAServidor(direccion string, req *hexpb.SupervisorRequest) bool {
	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Servidor Hextech: %v", err)
	}
	defer conn.Close()

	client := hexpb.NewHextechServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Manejo del reloj vectorial local
	key := fmt.Sprintf("%s-%s", req.Region, req.ProductName)
	if registro, existe := s.registros[key]; existe {
		req.KnownVectorClock = registro.VectorClock
	} else {
		req.KnownVectorClock = &hexpb.VectorClock{Server1: 0, Server2: 0, Server3: 0}
	}

	// Procesar la solicitud
	res, err := client.HandleRequest(ctx, req)
	if err != nil {
		log.Fatalf("Error al procesar la solicitud en el Servidor Hextech: %v", err)
	}

	// Manejar la respuesta del servidor
	if res.Status == hexpb.ResponseStatus_OK {
		fmt.Printf("Operación exitosa: [%d, %d, %d]\n", res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3)

		// Actualización de registros locales
		if req.OperationType != hexpb.OperationType_BORRAR {
			s.registros[key] = Registro{
				Region:      req.Region,
				Producto:    req.ProductName,
				Valor:       *(res.Value),
				VectorClock: res.VectorClock,
				Servidor:    direccion,
			}
		}
		return true
	} else if res.Status == hexpb.ResponseStatus_ERROR {
		fmt.Printf("Error: %s\n", *(res.Message))
		return false
	}
	return false
}

// **Informar inconsistencia al Broker y obtener nuevo servidor**
func (s *Supervisor) ResolverInconsistencia(region, producto string, vectorLocal *hexpb.VectorClock) string {
	currentVectorClock := &pb.VectorClock{
		Server1: vectorLocal.Server1,
		Server2: vectorLocal.Server2,
		Server3: vectorLocal.Server3,
	}
	req := &pb.InconsistencyRequest{
		Region:           region,
		ProductName:      producto,
		SupervisorClock:  currentVectorClock,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := s.client.ResolveInconsistency(ctx, req)
	if err != nil {
		log.Fatalf("Error al resolver inconsistencia con el Broker: %v", err)
	}

	log.Printf("Servidor asignado para resolver inconsistencia: %s", res.ServerAddress)
	return res.ServerAddress
}

// **Validación de consistencia (Read Your Writes)**
func (s *Supervisor) LeerConsistente(region, producto string) bool {
	key := fmt.Sprintf("%s-%s", region, producto)
	registro, existe := s.registros[key]

	if !existe {
		log.Fatalf("No existe información local para %s en la región %s", producto, region)
	}

	direccion := s.ObtenerServidor()
	req := &hexpb.SupervisorRequest{
		Region:           region,
		ProductName:      producto,
		OperationType:    hexpb.OperationType_ACTUALIZAR,
		KnownVectorClock: registro.VectorClock,
	}

	// Primero, obtendremos el reloj vectorial del servidor para la consistencia
	conn, err := grpc.Dial(direccion, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error al conectar con el servidor Hextech: %v", err)
	}
	defer conn.Close()

	client := hexpb.NewHextechServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Hacemos la solicitud al servidor Hextech
	res, err := client.HandleRequest(ctx, req)
	if err != nil {
		log.Fatalf("Error al obtener respuesta del servidor Hextech: %v", err)
	}

	// Verifica si el reloj vectorial del servidor es más reciente que el local
	if !esConsistente(res.VectorClock, registro.VectorClock) {
		// Si los relojes no son consistentes, volvemos a intentarlo
		log.Print("Datos no consistentes. Reintentando...")

		// Configurar reintento con tiempo de espera
		time.Sleep(100 * time.Millisecond) // Se podría hacer configurable

		// El método podría continuar reintentando hasta que se logre consistencia
		// Es importante evitar que el bucle quede atrapado si la consistencia no se resuelve
		return false
	}

	// Si los relojes son consistentes, continuamos
	fmt.Printf("Lectura consistente: [%d, %d, %d] - Valor: %d\n",
		res.VectorClock.Server1, res.VectorClock.Server2, res.VectorClock.Server3, res.Value)
	return true
}

// **Comparación de relojes vectoriales**
func esConsistente(v1, v2 *hexpb.VectorClock) bool {
	// Si el reloj v1 es mayor o igual que v2 en todas las posiciones, se considera consistente
	return v1.Server1 >= v2.Server1 && v1.Server2 >= v2.Server2 && v1.Server3 >= v2.Server3
}

// **Operaciones del Supervisor**
func (s *Supervisor) AgregarProducto(region, producto string, valor int32) {
	if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_AGREGAR,
			Value:         &valor,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else {
		direccion := s.ObtenerServidor()
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_AGREGAR,
			Value:         &valor,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	}
}

// RenombrarProducto renombra un producto en el registro de una región
func (s *Supervisor) RenombrarProducto(region, producto, nuevoNombre string) {
    if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else {
		direccion := s.ObtenerServidor()
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
			NewProductName: &nuevoNombre,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	}
}

// ActualizarValor actualiza el valor de un producto en el registro de una región
func (s *Supervisor) ActualizarValor(region, producto string, nuevoValor int32) {
    if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else {
		direccion := s.ObtenerServidor()
		req := &hexpb.SupervisorRequest{
			Region:       region,
			ProductName:  producto,
			OperationType: hexpb.OperationType_ACTUALIZAR,
			Value:        &nuevoValor,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	}
}

// BorrarProducto elimina un producto del registro de una región
func (s *Supervisor) BorrarProducto(region, producto string) {
	if !s.LeerConsistente(region, producto) {
		direccion := s.ResolverInconsistencia(region, producto, nil)
		req := &hexpb.SupervisorRequest{
			Region:        region,
			ProductName:   producto,
			OperationType: hexpb.OperationType_RENOMBRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	} else {
		direccion := s.ObtenerServidor()
		req := &hexpb.SupervisorRequest{
			Region:       region,
			ProductName:  producto,
			OperationType: hexpb.OperationType_BORRAR,
		}
		s.EnviarSolicitudAServidor(direccion, req)
	}
}

func main() {
	log.Print("El supervisor está corriendo...")

	supervisor := NuevoSupervisor()
	defer func() {
		if err := supervisor.conn.Close(); err != nil {
			log.Fatalf("Error al cerrar la conexión: %v", err)
		}
	}()

	supervisor.AgregarProducto("Noxus", "Vino", 25)
	supervisor.RenombrarProducto("Noxus", "Vino", "Cerveza")
	supervisor.ActualizarValor("Noxus", "Cerveza", 50)
	supervisor.BorrarProducto("Noxus", "Cerveza")
}
