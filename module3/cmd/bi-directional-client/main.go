package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/wgsaxton/go-grpc-class/module3/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// initialize our grpc connection
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// create a client
	client := proto.NewStreamingServiceClient(conn)

	// initialize our stream
	stream, err := client.Echo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	eg, _ := errgroup.WithContext(ctx)

	// create a separate go routine to listen to the server responses
	eg.Go(func() error {
		// loop for each message from the server
		for {
			res, err := stream.Recv()
			if err != nil {
				// break out if server sends io.EOF response
				if err == io.EOF {
					break
				}
				log.Printf("exit line 48, error: %s", err.Error())
				return err
			}

			// log the message
			log.Printf("Message received from server: %s", res.GetMessage())
		}

		return nil
	})

	// send some messages from the client
	for i := range 5 {
		req := &proto.EchoRequest{
			Message: fmt.Sprintf("Hello number %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Printf("exiting line 64, error: %s", err.Error())
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)
	}

	// close the client stream
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	// wait for the server go routine to finish
	log.Println("bi-directional stream closed")

}
