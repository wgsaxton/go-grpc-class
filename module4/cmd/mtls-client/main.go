package main

import (
	"context"
	"crypto/tls"
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

	certPool := x509.NewCertPool()
	cert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("failed to append CA cert")
	}

	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs: certPool,
	}

	tlsCredentials := credentials.NewTLS(tlsConfig)

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
