package main

import (
	proto "fileStreaming/proto"
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedFileServiceServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer() // engine
	proto.RegisterFileServiceServer(srv, &server{})
	reflection.Register(srv)
	fmt.Println("Server is running on port 9000")

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) UploadFile(stream proto.FileService_UploadFileServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	// var fileName string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// fileName = req.GetFilePath()
			break
		}
		chunks := req.GetChunks()
		fileBytes = append(fileBytes, chunks...)
		fileSize += int64(len(chunks))
	}

	f, err := os.Create("./abc.bin")
	if err != nil {
		return err
	}

	defer f.Close()
	_, err2 := f.Write(fileBytes)

	if err2 != nil {
		return err2
	}
	return stream.SendAndClose(&proto.UploadResponse{FileSize: fileSize, Message: "File written successfully"})
}
