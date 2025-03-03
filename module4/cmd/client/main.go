package main

import (
	"context"
	"crypto/x509"
	"log"
	"os"

	"github.com/wgsaxton/go-grpc-class/module4/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	// Public CA example
	// tlsCredentials := credentials.NewTLS(&tls.Config{})

	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("failed to append CA cert")
	}

	tlsCredentials := credentials.NewClientTLSFromCert(certPool, "")

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Garrett123"})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code: %s, error %s", status.Code().String(), status.Message())
		}
		log.Fatal(err)
	}

	log.Printf("Response received: %s", res.GetMessage())

}
