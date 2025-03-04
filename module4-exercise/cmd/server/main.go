package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/wgsaxton/go-grpc-class/module4-exercise/internal/stream"
	"github.com/wgsaxton/go-grpc-class/module4-exercise/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running the application", slog.String("error", err.Error()))
	}
}

func run(ctx context.Context) error {
	serverCert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		return fmt.Errorf("failed to load server tls certs: %w", err)
	}

	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		return fmt.Errorf("failed to read the CA cert: %w", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return errors.New("failed to append CA cert to cert pool")
	}

	tlsCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	streamingService := stream.Service{}

	proto.RegisterFileUploadServiceServer(grpcServer, streamingService)

	const addr = "localhost:50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("starting grpc server on address", slog.String("address", addr))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve the grpc service: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return nil
	})

	return g.Wait()
}
