package main

import (
	"context"
	"log"

	"github.com/cndoit18/practice-grpc/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	client := types.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &types.HelloRequest{
		Name: "name",
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(resp.GetMessage())

	resp, err = client.SayHelloAgain(context.Background(), &types.HelloRequest{
		Name: "name",
	})
	if err != nil {
		log.Panic(err)
	}

	log.Println(resp.GetMessage())
}
