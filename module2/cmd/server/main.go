package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/wgsaxton/go-grpc-class/module2/internal/hello"
	"github.com/wgsaxton/go-grpc-class/module2/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("Error running application", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("Closing server gracefully")
}

func run(ctx context.Context) error {
	grpcServer := grpc.NewServer()

	helloService := hello.Service{}

	proto.RegisterHelloServiceServer(grpcServer, helloService)

	const addr = ":50051"

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			return fmt.Errorf("Failed to listen on address %q: %w", addr, err)
		}

		slog.Info("Starting grpc server on", slog.String("address", addr))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve gRPC service: %w", err)
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
