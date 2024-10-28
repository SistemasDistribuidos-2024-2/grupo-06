// diaboromon.go
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

	pb "diaboromon/grpc/tai-diaboromon"

	"google.golang.org/grpc"
)

const (
	taiAddress = ":50051" // Dirección del Nodo Tai
	port       = ":50052" // Puerto en el que escucha Diaboromon
	inputFile  = "/app/input.txt"
)

var (
	TD int // Tiempo de ataque de Diaboromon
	VI int // Vida inicial de Greymon/Garurumon
)

// Cargar el valor de TD desde INPUT.txt
func cargarParametros() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error al abrir INPUT.txt: %v", err)
	}
	defer file.Close()
	var line string
	fmt.Fscanln(file, &line)
	valores := strings.Split(line, ",")

	if len(valores) != 5 {
		log.Fatalf("Formato incorrecto en INPUT.txt")
	}

	// Ignorar PS y TE ya que Diaboromon solo necesita TD y VI
	TD, _ = strconv.Atoi(valores[2])
	VI, _ = strconv.Atoi(valores[4])
}

// Servidor Diaboromon que implementa TaiDiaboromonServiceServer
type server struct {
	pb.UnimplementedTaiDiaboromonServiceServer
}

// Ataque de Diaboromon a Nodo Tai
func (s *server) AtaqueDiaboromon(ctx context.Context, req *pb.SolicitudAtaque) (*pb.ConfirmacionAtaque, error) {
	log.Println("Diaboromon: Ataque a Greymon/Garurumon enviado al Nodo Tai.")
	return &pb.ConfirmacionAtaque{Mensaje: "Ataque recibido por Nodo Tai"}, nil
}

func main() {
	// Cargar el parámetro TD
	cargarParametros()

	// Configurar el servidor gRPC de Diaboromon
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTaiDiaboromonServiceServer(grpcServer, &server{})

	log.Printf("Diaboromon escuchando en el puerto %v...", port)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()

	// Conectar a Nodo Tai como cliente
	conn, err := grpc.Dial(taiAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Nodo Tai: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaiDiaboromonServiceClient(conn)
	// Ataques periódicos cada TD segundos
	ticker := time.NewTicker(time.Duration(TD) * time.Second)
	defer ticker.Stop()

	time.Sleep(10 * time.Second)

	for range ticker.C {
		// Enviar ataque a Nodo Tai
		req := &pb.SolicitudAtaque{Mensaje: "Diaboromon ataca"}
		_, err := client.AtaqueDiaboromon(context.Background(), req)
		if err != nil {
			log.Printf("Error al enviar ataque a Nodo Tai: %v", err)
			continue
		}

		// Solicitar estado de vida restante a Nodo Tai
		estadoReq := &pb.SolicitudEstado{Mensaje: "Solicitar estado de vida"}
		res, err := client.EstadoVida(context.Background(), estadoReq)
		if err != nil {
			log.Printf("Error al obtener estado de vida de Nodo Tai: %v", err)
			continue
		}

		log.Printf("Estado de vida recibido del Nodo Tai: %d puntos restantes", res.VidaRestante)

		if res.VidaRestante <= 0 {
			log.Println("Diaboromon ha vencido. Fin de la ejecución.")
			os.Exit(0)
		}
	}
}
