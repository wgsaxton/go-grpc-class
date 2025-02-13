package main

import (
	"context"
	"log"

	"github.com/wgsaxton/go-grpc-class/module2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// conn, err := grpc.Dial("localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithBlock(),
	// )

	conn, err := grpc.NewClient("localhost:50051",
	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// grpc.WithBlock(), // looks like this will be deprecated
)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Garrett S"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("response received: %s", res.GetMessage())
}
