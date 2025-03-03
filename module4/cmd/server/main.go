package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/wgsaxton/go-grpc-class/module4/internal/hello"
	"github.com/wgsaxton/go-grpc-class/module4/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running the application", slog.String("error", err.Error()))
		os.Exit(1)
	}

}

func run(ctx context.Context) error {
	// Load the tls certs
	tlsCredentials, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		return fmt.Errorf("failed to load tls credentials: %w", err)
	}
	
	// Create and start gRPC server
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, helloService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", addr, err)
		}

		slog.Info("Starting grpc server on address", slog.String("address", addr))

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

	// Isn't it waiting for both go routines to stop?
	// I think is the 2nd go routine stops with ctx.Done()
	// then grpcServer.GracefulStop() will stop the first go routine
	// But if grpcServer.Serve(lis) gets an error and stops blocking
	// how will the 2nd go routine stop?
	// Because of errgroup.WithContext(ctx), all other go routines
	// are stopped if 1 g.Go() routine errors out
	return g.Wait()
}
