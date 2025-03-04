package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/wgsaxton/go-grpc-class/module4-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	clientCert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		log.Fatalf("failed to log client certs: %w", err)
	}

	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA cert: %w", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		log.Fatal("failed to append CA cert to pool")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	creds := credentials.NewTLS(tlsConfig)

	// initialise gRPC connection
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewFileUploadServiceClient(conn)

	http.HandleFunc("/", downloadHander(client))

	log.Printf("starting http server on address: %s", "localhost:8080")

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}

}

func downloadHander(client proto.FileUploadServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// make request to grpc server and initialize server stream
		stream, err := client.DownloadFile(ctx, &proto.DownloadFileRequest{Name: "gopher.png"})
		if err != nil {
			// check the status code received from server
			st := status.Convert(err)
			switch st.Code() {
			case codes.NotFound:
				http.Error(w, "File not found", http.StatusNotFound)
				return
			case codes.InvalidArgument:
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("server stream started")

		// create a slice of file contents
		var fileContents []byte

		for {
			// receive file chunk
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// break out of loop. stream is done
					break
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Println("receive file chunk")

			// append the file chunk to slice
			fileContents = append(fileContents, res.GetContent()...)
		}

		log.Println("server stream done")

		// return file contents to user (web page)
		if _, err := w.Write(fileContents); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
