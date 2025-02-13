package main

import (
	"log"
	"net"

	"github.com/wgsaxton/go-grpc-class/module2/internal/hello"
	"github.com/wgsaxton/go-grpc-class/module2/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, helloService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting grpc server on address: %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
