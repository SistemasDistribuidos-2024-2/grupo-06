package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pbDiaboromon "nodo_tai/grpc/tai-diaboromon" // Para la comunicación con Diaboromon
	pbPrimary "nodo_tai/grpc/tai-primary"       // Para la comunicación con el Primary Node

	"google.golang.org/grpc"
)

const (
	primaryNodeAddress   = "localhost:50051" // Dirección del Primary Node
	nodoTaiPort          = ":50055"          // Puerto del Nodo Tai para actuar como servidor
	diaboromonAddress    = "localhost:50052" // Dirección de Diaboromon
	inputFile            = "../INPUT.txt"
)

var (
	PS             float64 // Probabilidad de sacrificio
	TE             int     // Tiempo de espera para enviar información
	TD             int     // Tiempo de ataque de Diaboromon
	CD             int     // Cantidad de datos necesarios para evolucionar
	VI             int     // Vida inicial para Greymon y Garurumon
	vida           int     // Vida restante
	datosAcumulados float32 // Cantidad de datos acumulados
)

// Leer los valores de INPUT.txt
func leerInput() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo INPUT.txt: %v", err)
	}
	defer file.Close()

	var line string
	fmt.Fscanln(file, &line)
	valores := strings.Split(line, ",")

	if len(valores) != 5 {
		log.Fatalf("Formato incorrecto en INPUT.txt")
	}

	// Parsear cada valor
	PS, _ = strconv.ParseFloat(valores[0], 64)
	TE, _ = strconv.Atoi(valores[1])
	TD, _ = strconv.Atoi(valores[2])
	CD, _ = strconv.Atoi(valores[3])
	VI, _ = strconv.Atoi(valores[4])

	vida = VI // Inicializar la vida con el valor de VI
}

// Solicitar datos acumulados al Primary Node
func SolicitarDatos(client pbPrimary.TaiNodeServiceClient) {
	solicitud := &pbPrimary.SolicitudTai{Mensaje: "Solicito cantidad de datos acumulados"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	respuesta, err := client.SolicitarCantidadDatos(ctx, solicitud)
	if err != nil {
		log.Fatalf("Error al solicitar datos al Primary Node: %v", err)
	}

	datosAcumulados = respuesta.CantidadDatos
	log.Printf("Cantidad de datos acumulados: %.2f", datosAcumulados)
}

// Implementación del servidor para recibir ataques de Diaboromon
type server struct {
	pbDiaboromon.UnimplementedTaiDiaboromonServiceServer
	diaboromonClient pbDiaboromon.TaiDiaboromonServiceClient
}

// Procesar el ataque de Diaboromon y responder con el estado de vida
func (s *server) AtaqueDiaboromon(ctx context.Context, req *pbDiaboromon.SolicitudAtaque) (*pbDiaboromon.ConfirmacionAtaque, error) {
	vida -= 10
	log.Printf("Diaboromon ataca: Vida restante de Greymon/Garurumon: %d", vida)

	if vida <= 0 {
		log.Println("Diaboromon ha vencido. Fin de la ejecución.")
		os.Exit(0)
	} else {
		// Enviar estado de vida de vuelta a Diaboromon
		s.enviarEstadoVida()
	}

	return &pbDiaboromon.ConfirmacionAtaque{Mensaje: "Ataque recibido por Nodo Tai"}, nil
}

// Enviar estado de vida actual a Diaboromon
func (s *server) enviarEstadoVida() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	estado := &pbDiaboromon.SolicitudEstado{Mensaje: "Solicitar estado de vida"}
	respuesta, err := s.diaboromonClient.EstadoVida(ctx, estado)
	if err != nil {
		log.Printf("Error al enviar estado de vida a Diaboromon: %v", err)
		return
	}
	log.Printf("Estado de vida enviado a Diaboromon: Vida restante %d", respuesta.VidaRestante)
}

// Atacar a Diaboromon
func atacarDiaboromon() {
	if datosAcumulados >= float32(CD) {
		log.Println("Greymon/Garurumon evolucionan a Omegamon y derrotan a Diaboromon. Fin de la ejecución.")
		os.Exit(0)
	} else {
		log.Printf("Diaboromon repele el ataque, vida restante de Greymon/Garurumon: %d", vida)
		if vida <= 0 {
			log.Println("Diaboromon ha vencido. Fin de la ejecución.")
			os.Exit(0)
		}
	}
}

// Ejecutar el ciclo principal de ataques y defensa
func cicloPrincipal(primaryClient pbPrimary.TaiNodeServiceClient) {
	// Solicitar datos al Primary Node inicialmente
	SolicitarDatos(primaryClient)

	// Iniciar el ciclo de verificación de datos
	ticker := time.NewTicker(time.Duration(TE) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Verificar si Greymon/Garurumon tienen datos suficientes para atacar
		SolicitarDatos(primaryClient)
		if datosAcumulados >= float32(CD) {
			atacarDiaboromon()
		}
	}
}

func main() {
	// Leer INPUT.txt
	leerInput()

	// Conectar al Primary Node como cliente
	primaryConn, err := grpc.Dial(primaryNodeAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Primary Node: %v", err)
	}
	defer primaryConn.Close()
	primaryClient := pbPrimary.NewTaiNodeServiceClient(primaryConn)

	// Conectar a Diaboromon como cliente
	diaboromonConn, err := grpc.Dial(diaboromonAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar a Diaboromon: %v", err)
	}
	defer diaboromonConn.Close()
	diaboromonClient := pbDiaboromon.NewTaiDiaboromonServiceClient(diaboromonConn)

	// Iniciar el servidor gRPC para recibir ataques de Diaboromon
	go func() {
		lis, err := net.Listen("tcp", nodoTaiPort)
		if err != nil {
			log.Fatalf("Error al escuchar en el puerto %s: %v", nodoTaiPort, err)
		}

		grpcServer := grpc.NewServer()
		pbDiaboromon.RegisterTaiDiaboromonServiceServer(grpcServer, &server{diaboromonClient: diaboromonClient})

		log.Printf("Nodo Tai escuchando en el puerto %s para recibir ataques de Diaboromon", nodoTaiPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()

	// Iniciar el ciclo principal de combate
	cicloPrincipal(primaryClient)
}
