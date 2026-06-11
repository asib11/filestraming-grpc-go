# Filestraming using gRPC-go

A file streaming service built with **gRPC** and **Go**, demonstrating how to transfer files in chunks using gRPC's client-side streaming. This project is a practical example of using Protocol Buffers (protobuf) and the gRPC framework to build efficient, high-performance file transfer between a client and server.

---

## Features

- File transfer using **gRPC client-side streaming**
- Chunked file upload — handles files of any size
- Clean separation of `server`, `client`, and `proto` packages
- Built with **proto3** syntax
- Uses the latest `google.golang.org/grpc` and `google.golang.org/protobuf` libraries

---

## Project Structure

```
filestraming-grpc-go/
├── client/         # gRPC client — reads file and streams chunks to server
├── server/         # gRPC server — receives streamed chunks and saves file
├── proto/          # Protobuf definition and generated Go code
├── go.mod
├── go.sum
└── README.md
```

---

## Prerequisites

Make sure the following are installed on your system:

- [Go 1.21+](https://go.dev/dl/)
- [protoc](https://grpc.io/docs/protoc-installation/) — Protocol Buffer compiler (v3)
- protoc Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Update your `PATH` so `protoc` can find the plugins:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Setup & Installation

### 1. Clone the repository

```bash
git clone https://github.com/asib11/filestraming-grpc-go.git
cd filestraming-grpc-go
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. (Optional) Regenerate protobuf code

If you modify the `.proto` file, regenerate the Go code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/*.proto
```

---

## Running the Project

### Start the Server

```bash
cd server
go run main.go
```

The server will start listening on the configured port (default: `:50051`).

### Run the Client

Open a new terminal and run:

```bash
cd client
go run main.go
```

The client will read the target file, split it into chunks, and stream them to the server via gRPC.

---

## gRPC Service Definition

The service is defined in `proto/` using proto3 syntax. The core structure follows the client-streaming RPC pattern:

```protobuf
syntax = "proto3";

message FileUploadRequest {
  string file_name = 1;
  bytes  chunk     = 2;
}

message FileUploadResponse {
  string file_name = 1;
  uint32 size      = 2;
}

service FileService {
  rpc Upload(stream FileUploadRequest) returns (FileUploadResponse);
}
```

The client streams multiple `FileUploadRequest` messages (each carrying a chunk of the file), and the server responds with a single `FileUploadResponse` once the entire file is received.

---

## Dependencies

| Package | Version |
|---|---|
| `google.golang.org/grpc` | v1.81.1 |
| `google.golang.org/protobuf` | v1.36.11 |
| `golang.org/x/net` | v0.51.0 |

---

## How It Works

1. **Client** opens a local file and reads it in fixed-size chunks.
2. Each chunk is wrapped in a `FileUploadRequest` message and sent to the server over a gRPC stream.
3. **Server** listens for the incoming stream, reassembles the chunks, and writes the file to disk.
4. Once the client closes the stream, the server sends back a `FileUploadResponse` with the file name and total size.

```
Client                          Server
  |---[FileUploadRequest]-------->|
  |---[FileUploadRequest]-------->|
  |---[FileUploadRequest]-------->|
  |---[CloseStream]-------------->|
  |<--[FileUploadResponse]--------|
```

---

## Learn More

- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers (proto3) Guide](https://protobuf.dev/programming-guides/proto3/)
- [grpc-go on GitHub](https://github.com/grpc/grpc-go)

---


