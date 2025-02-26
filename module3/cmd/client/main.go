package main

import (
	"context"
	"io"
	"log"

	"github.com/wgsaxton/go-grpc-class/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// first initialize grpc connection
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// create the client
	client := proto.NewStreamingServiceClient(conn)

	// initialize the stream
	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	// loop through all the responses we get back from the server
	// log it
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		log.Printf("Recieved time from server: %s", res.GetCurrentTime().AsTime())
	}

	// once the server closes the stream, exit gracefully
	log.Println("Streaming server closed")

}
