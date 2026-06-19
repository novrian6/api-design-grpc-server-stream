package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-server-streaming/proto/stock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gagal terhubung: %v", err)
	}
	defer conn.Close()

	client := pb.NewStockServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.StockRequest{
		Symbol:          "GOOGL",
		DurationSeconds: 5,
	}

	stream, err := client.GetStockUpdates(ctx, req)
	if err != nil {
		log.Fatalf("gagal memulai stream: %v", err)
	}

	log.Println("Menerima data streaming harga saham:")
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server telah selesai mengirim data")
			break
		}
		if err != nil {
			log.Fatalf("kesalahan saat membaca stream: %v", err)
		}

		log.Printf("Symbol: %s | Harga: %.2f | Perubahan: %.2f | Waktu: %d",
			resp.Symbol, resp.Price, resp.Change, resp.Timestamp)
	}
}
