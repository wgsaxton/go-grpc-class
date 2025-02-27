package streaming

import (
	"io"
	"log"
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
			// get the time
			currentTime := time.Now()

			// build our response
			resp := &proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}

			// return that to the client
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}

}

func (s Service) LogStream(stream grpc.ClientStreamingServer[proto.LogStreamRequest, proto.LogStreamResponse]) error {
	// initialize a count
	count := 0
	// loop through all the received messages
	for {
		// receive our message
		logEntry, err := stream.Recv()
		if err != nil {
			// check if the stream is closed
			if err == io.EOF {
				return stream.SendAndClose(&proto.LogStreamResponse{
					EntriesLogged: int32(count),
				})
			}
			return err
		}

		// log message
		log.Printf("Received log [%s]: %s - %s", logEntry.GetTimestamp().AsTime(), logEntry.GetLevel().String(), logEntry.GetMessage())
		// increment count
		count++
	}

}
