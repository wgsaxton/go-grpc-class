package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/wgsaxton/go-grpc-class/module6/internal/config"
	"github.com/wgsaxton/go-grpc-class/module6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	cfg := config.Config{
		MethodConfig: []config.MethodConfig{{
			Name: []config.NameConfig{{
				Service: "config.ConfigService",
			}},
			RetryPolicy: config.RetryPolicy{
				MaxAttempts:          4,
				InitialBackoff:       "0.1s",
				MaxBackoff:           "1s",
				BackoffMultiplier:    2,
				RetryableStatusCodes: []string{"INTERNAL", "UNAVAILABLE"},
			},
		}},
	}

	serviceConfig, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(string(serviceConfig)),
	)
	if err != nil {
		log.Fatalf("failed to create the client: %v", err)
	}

	client := proto.NewConfigServiceClient(conn)

	_, err = client.Flaky(ctx, &proto.FlakyRequest{})
	if err != nil {
		log.Fatal(err)
	}


}
