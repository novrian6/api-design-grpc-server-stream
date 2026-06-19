package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "grpc-server-streaming/proto/stock"

	"google.golang.org/grpc"
)

type stockServer struct {
	pb.UnimplementedStockServiceServer
}

func (s *stockServer) GetStockUpdates(req *pb.StockRequest, stream pb.StockService_GetStockUpdatesServer) error {
	symbol := req.Symbol
	duration := req.DurationSeconds
	endTime := time.Now().Add(time.Duration(duration) * time.Second)
	currentPrice := 100.0 + rand.Float64()*50

	for time.Now().Before(endTime) {
		change := (rand.Float64() - 0.5) * 2.0
		currentPrice += change

		resp := &pb.StockResponse{
			Symbol:    symbol,
			Price:     currentPrice,
			Change:    change,
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(resp); err != nil {
			return fmt.Errorf("error sending stream: %w", err)
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("Selesai mengalirkan data untuk %s selama %d detik", symbol, duration)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("gagal mendengarkan: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &stockServer{})

	log.Println("Server streaming berjalan di port :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gagal menjalankan server: %v", err)
	}
}
