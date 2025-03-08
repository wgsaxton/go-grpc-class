package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wgsaxton/go-grpc-class/module6/internal/resolve"
	"github.com/wgsaxton/go-grpc-class/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func main() {
	ctx := context.Background()

	builder := resolve.Builder{}
	resolver.Register(&builder)

	const serviceConfig = `{"loadBalancingPolicy": "round_robin"}`


	// whateverscheme://
	conn, err := grpc.NewClient(fmt.Sprintf("%s://", builder.Scheme()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewConfigServiceClient(conn)

	for i := range 12 {
		log.Printf("making request: %d", i)

		res, err := client.GetServerAddress(ctx, &proto.GetServerAddressRequest{})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("response received: %s", res.GetAddress())

		time.Sleep(time.Second)
	}
}
