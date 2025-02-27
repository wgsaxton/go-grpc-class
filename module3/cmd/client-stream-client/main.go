package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wgsaxton/go-grpc-class/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// initialize grpc connection
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	//create our client
	client := proto.NewStreamingServiceClient(conn)

	// initialize the client stream
	stream, err := client.LogStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// send some log messages
	for i := 0; i < 5; i++ {
		req := &proto.LogStreamRequest{
			Timestamp: timestamppb.New(time.Now()),
			Level:     proto.LogLevel_LOG_LEVEL_INFO,
			Message:   fmt.Sprintf("Hello log: %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}

	// close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	// log the response from server
	log.Printf("Server says number of messages received: %d", res.GetEntriesLogged())
}
