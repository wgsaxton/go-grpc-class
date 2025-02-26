package streaming

import (
	"time"

	"github.com/wgsaxton/go-grpc-class/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

func (s Service) StreamServerTime(request *proto.StreamServerTimeRequest, stream grpc.ServerStreamingServer[proto.StreamServerTimeResponse]) error {
	// initialize a ticker for our interval
	if request.GetIntervalSeconds() == 0 {
		return status.Error(codes.InvalidArgument, "interval must be set")
	}

	interval := time.Duration(request.GetIntervalSeconds()) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// loop through and listen on the ticker
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			currentTime := time.Now()

			resp := &proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}

			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}

	// get the time
	// build our response
	// return that to the client

	// make sure the context is not cancelled
}
