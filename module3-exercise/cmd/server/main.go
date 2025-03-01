package main

import (
	"log"
	"net"

	"github.com/wgsaxton/go-grpc-class/module3-exercise/internal/stream"
	"github.com/wgsaxton/go-grpc-class/module3-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	streamingService := &stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer, streamingService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
