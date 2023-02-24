package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cndoit18/practice-grpc/types"
	"google.golang.org/grpc"
)

var _ types.GreeterServer = &server{}

type server struct {
	types.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, req *types.HelloRequest) (*types.HelloReply, error) {
	return &types.HelloReply{Message: fmt.Sprintf("echo: %s", req.Name)}, nil
}

// Sends another greeting
func (s *server) SayHelloAgain(_ context.Context, req *types.HelloRequest) (*types.HelloReply, error) {
	return &types.HelloReply{Message: fmt.Sprintf("echo again: %s", req.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Listen %s", lis.Addr().String())

	srv := grpc.NewServer()
	types.RegisterGreeterServer(srv, &server{})

	if err := srv.Serve(lis); err != nil {
		log.Panic(err)
	}
}
