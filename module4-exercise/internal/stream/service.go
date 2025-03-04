package stream

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/wgsaxton/go-grpc-class/module4-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedFileUploadServiceServer
}

func (s Service) DownloadFile(req *proto.DownloadFileRequest, stream grpc.ServerStreamingServer[proto.DownloadFileResponse]) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "filename cannot be empty")
	}

	// open file on the server
	file, err := os.Open(req.GetName())
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file not found")
		}
		return err
	}
	defer file.Close()

	const bufferSize = 5 * 1024 // send in 5KB chuncks
	buff := make([]byte, bufferSize)
	for {
		// read bytes
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				return nil // end of file, close the stream
			}
			return status.Error(codes.Internal, "error reading file")
		}

		// stream bytes to client
		if err = stream.Send(&proto.DownloadFileResponse{Content: buff[:bytesRead]}); err != nil {
			return status.Error(codes.Internal, "error streaming file")
		}
	}
}

func (s Service) UploadFile(stream grpc.ClientStreamingServer[proto.UploadFileRequest, proto.UploadFileResponse]) error {
	// generate a file name
	fileName := fmt.Sprintf("%s.png", uuid.New().String())

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		return status.Error(codes.Internal, "error creating file")
	}
	defer file.Close()

	for {
		// receive chunk
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// close stream and send response to client
				return stream.SendAndClose(&proto.UploadFileResponse{Name: fileName})
			}
			return err
		}

		// write chunk to file
		bw, err := file.Write(res.GetContent())
		if err != nil {
			return status.Errorf(codes.Internal, "error writing to file. tried to write %d bytes", bw)
		}

		log.Printf("bytes written: %d", bw)
	}
}
