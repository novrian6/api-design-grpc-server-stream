# gRPC Server Streaming

Demo gRPC server-streaming untuk streaming data harga saham secara real-time.

## Struktur Project

```
grpc-server-stream/
├── README.md                         (file ini)
├── grpc-server-streaming/            # Project utama (server + client bawaan)
│   ├── go.mod / go.sum
│   ├── server/
│   │   └── main.go                   # gRPC server (port :50052)
│   ├── client/
│   │   └── main.go                   # Client bawaan
│   └── proto/
│       ├── stock.proto               # Definisi protobuf
│       └── stock/
│           ├── stock.pb.go           # Generated message types
│           └── stock_grpc.pb.go      # Generated gRPC service
└── client-call/                      # Client project terpisah (independen)
    ├── go.mod / go.sum
    ├── main.go                       # Client dengan tampilan tabel
    └── proto/stock/
        ├── stock.pb.go               # Copy dari generated proto
        └── stock_grpc.pb.go          # Copy dari generated proto
```

## Cara Menjalankan

### 1. Jalankan Server (wajib)

```bash
# Dari root grpc-server-stream
cd grpc-server-streaming
go run server/main.go
```

Server akan berjalan di **port 50052** dan siap menerima koneksi client.

### 2. Jalankan Client

#### Opsi A — Client bawaan (`grpc-server-streaming/client`)

```bash
# Terminal terpisah, dari root grpc-server-stream
cd grpc-server-streaming
go run client/main.go
```

#### Opsi B — Client independen (`client-call`)

```bash
# Terminal terpisah, dari root grpc-server-stream
cd client-call
go run main.go
```

## Cara Kerja

1. **Server** menerima request `StockRequest` (symbol, duration_seconds) dan melakukan streaming data harga saham dengan perubahan acak setiap detik
2. **Client** menerima stream menggunakan `Recv()` dalam loop hingga server mengirim sinyal EOF

**Contoh request:**
```protobuf
StockRequest {
    symbol: "GOOGL"
    duration_seconds: 5
}
```

**Server akan mengirim 5 data** (1 data per detik) dengan format:
```protobuf
StockResponse {
    symbol: "GOOGL"
    price: 113.45
    change: 0.60
    timestamp: 1781580783
}
```

## Port

| Service  | Port  |
|----------|-------|
| gRPC     | 50052 |# api-design-grpc-server-stream
