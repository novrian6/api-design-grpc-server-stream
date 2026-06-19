package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "client-call/proto/stock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gagal terhubung ke server: %v", err)
	}
	defer conn.Close()

	client := pb.NewStockServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Request streaming data for GOOGL stock for 5 seconds
	req := &pb.StockRequest{
		Symbol:          "GOOGL",
		DurationSeconds: 5,
	}

	stream, err := client.GetStockUpdates(ctx, req)
	if err != nil {
		log.Fatalf("gagal memulai stream: %v", err)
	}

	fmt.Println("=== Menerima data streaming harga saham (client-call) ===")
	fmt.Printf("%-10s %-10s %-12s %-12s\n", "Symbol", "Harga", "Perubahan", "Waktu")
	fmt.Println("---------------------------------------------------")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\n✓ Server telah selesai mengirim data")
			break
		}
		if err != nil {
			log.Fatalf("kesalahan saat membaca stream: %v", err)
		}

		t := time.Unix(resp.Timestamp, 0).Format("15:04:05")
		changeSymbol := "+"
		if resp.Change < 0 {
			changeSymbol = ""
		}
		fmt.Printf("%-10s %-10.2f %s%-11.2f %-12s\n", resp.Symbol, resp.Price, changeSymbol, resp.Change, t)
	}

	fmt.Println("=== Selesai ===")
}
