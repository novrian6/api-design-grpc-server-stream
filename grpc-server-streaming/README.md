# gRPC Server Streaming - Real-time Stock Price Updates

Proyek ini mendemonstrasikan implementasi **server streaming** gRPC di Go. Server mengirimkan update harga saham secara real-time sebagai respons atas satu permintaan dari klien.

## Struktur Project

```
grpc-server-streaming/
├── proto/
│   ├── stock.proto          # Definisi protobuf
│   └── stock/               # Hasil generate protobuf
│       ├── stock.pb.go
│       └── stock_grpc.pb.go
├── server/
│   └── main.go              # Implementasi server
├── client/
│   └── main.go              # Implementasi klien
├── go.mod
├── go.sum
└── README.md
```

## Cara Kerja

1. Klien mengirim satu permintaan `StockRequest` berisi simbol saham dan durasi streaming.
2. Server mengirimkan serangkaian `StockResponse` (harga, perubahan, timestamp) setiap 1 detik selama durasi yang diminta.
3. Klien menerima dan menampilkan data streaming hingga server selesai.

## Prasyarat

- Go 1.23+
- Protocol Buffers Compiler (`protoc`)
- Protobuf Go plugin:
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## Menjalankan Aplikasi

### 1. Generate Kode Protobuf (jika mengubah proto/stock.proto)

```bash
protoc --go_out=. --go-grpc_out=. proto/stock.proto
```

### 2. Download Dependencies

```bash
go mod tidy
```

### 3. Jalankan Server

Terminal 1:
```bash
go run server/main.go
```

Server akan berjalan di port `:50052`.

Output yang diharapkan:
```
Server streaming berjalan di port :50052
```

### 4. Jalankan Klien

Terminal 2:
```bash
go run client/main.go
```

Output yang diharapkan (contoh):
```
Menerima data streaming harga saham:
Symbol: GOOGL | Harga: 120.34 | Perubahan: 0.87 | Waktu: 1718511200
Symbol: GOOGL | Harga: 119.87 | Perubahan: -0.47 | Waktu: 1718511201
Symbol: GOOGL | Harga: 121.02 | Perubahan: 1.15 | Waktu: 1718511202
Symbol: GOOGL | Harga: 120.56 | Perubahan: -0.46 | Waktu: 1718511203
Symbol: GOOGL | Harga: 121.34 | Perubahan: 0.78 | Waktu: 1718511204
Server telah selesai mengirim data
```

## Detail Implementasi

- **Port**: Server berjalan di `:50052`
- **Interval**: Server mengirim update setiap 1 detik
- **Durasi default klien**: 5 detik (dapat diubah di `client/main.go`)
- **Simbol default**: GOOGL (dapat diubah di `client/main.go`)
- **Harga awal**: Random antara 100 - 150
- **Perubahan harga**: Random antara -1.0 hingga +1.0 setiap detik