package config

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/wgsaxton/go-grpc-class/module6/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	name string
	proto.UnimplementedConfigServiceServer
}

func NewService(name string) (*service, error) {
	if name == "" {
		return nil, errors.New("service name cannot be empty")
	}

	return &service{
		name: name,
	}, nil
}

func (s service) LongRunning(ctx context.Context, req *proto.LongRunningRequest) (*proto.LongRunningResponse, error) {
	select {
	case <-time.Tick(time.Second * 5):
		log.Println("finish request")
	case <-ctx.Done():
		log.Println("context done")
	}

	return &proto.LongRunningResponse{}, nil
}

func (s service) Flaky(ctx context.Context, req *proto.FlakyRequest) (*proto.FlakyResponse, error) {
	if rand.Intn(3) != 0 {
		log.Println("error respons returned")
		return nil, status.Error(codes.Internal, "flaky error occurred")
	}

	log.Println("successful response occurred")

	return &proto.FlakyResponse{}, nil
}

func (s service) GetServerAddress(context.Context, *proto.GetServerAddressRequest) (*proto.GetServerAddressResponse, error) {
	log.Printf("address received: %s", s.name)

	return &proto.GetServerAddressResponse{Address: s.name}, nil
}