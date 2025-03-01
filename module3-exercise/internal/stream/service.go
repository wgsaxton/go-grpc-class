package stream

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/wgsaxton/go-grpc-class/module3-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	// below cannot be a pointer or it errors out
	proto.UnimplementedFileUploadServiceServer
}

func (s Service) DownloadFile(req *proto.DownloadFileRequest, stream grpc.ServerStreamingServer[proto.DownloadFileResponse]) error {
	// check if the name is empty in the request
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "file name cannot be empty")
	}

	// open file
	file, err := os.Open(req.GetName())
	if err != nil {
		// check if the file is found. if not, return not found
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file not found")
		}
		return err
	}
	// read the file in chunks of 5kb
	const bufferSize = 5 * 1024
	buff := make([]byte, bufferSize)
	for {
		bytes, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				// done reading file. close the server stream
				return nil
			}
			return status.Error(codes.Internal, "Error reading file")
		}

		// stream the chunk to the client
		err = stream.Send(&proto.DownloadFileResponse{Content: buff[:bytes]})
		if err != nil {
			return err
		}
	}

}

func (s Service) UploadFile(stream grpc.ClientStreamingServer[proto.UploadFileRequest, proto.UploadFileResponse]) error {
	// generate the filename
	fileName := fmt.Sprintf("%s.png", uuid.New().String())

	// create a file
	file, err := os.Create(fileName)
	if err != nil {
		return status.Error(codes.Internal, "error creating file")
	}
	defer file.Close()

	// receive chunks from client
	for {
		res, err := stream.Recv()
		if err != nil {
			// client has closed the stream, end of file
			if err == io.EOF {
				return stream.SendAndClose(&proto.UploadFileResponse{
					Name: fileName,
				})
			}
			return err
		}

		// write chunk to file
		bw, err := file.Write(res.GetContent())
		if err != nil {
			return status.Error(codes.Internal, "error writing to file")
		}

		log.Printf("bytes written: %d", bw)
	}

}
