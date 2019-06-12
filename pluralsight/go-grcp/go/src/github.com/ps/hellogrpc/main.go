package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"pluralsight/go-grcp/go/src/github.com/ps/hellogrpc/messages"
)

const port = ":50000"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	messages.RegisterHelloServiceServer(s, &server{})
	fmt.Printf("rpc server on %s...\n", port)
	s.Serve(listen)
}
