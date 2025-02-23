package main

import (
	"log"
	"net"

	"github.com/wgsaxton/go-grpc-class/module2-exercise/internal/todo"
	"github.com/wgsaxton/go-grpc-class/module2-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	todoService := todo.NewService()

	proto.RegisterTodoServiceServer(grpcServer, todoService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server running at address %s", ":50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
