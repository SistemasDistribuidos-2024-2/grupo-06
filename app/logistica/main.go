package main

import (
	"context"
	"log"
	"net"

	pb "logistica/proto/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedLogisticsServiceServer
}

func (s *server) ProcessOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    // Implementación de la lógica para procesar una orden
    log.Printf("Processing order: %v", req.OrderId)
    response := &pb.OrderResponse{
        OrderId:    req.OrderId,
        Status:     "Processed",
        TrackingId: "tracking_id_example",
        FinalValue: req.Value,
    }
    return response, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterLogisticsServiceServer(s, &server{})

    log.Printf("Server is running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
