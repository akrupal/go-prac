package main

import (
	"context"
	"grpc/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server")
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, _ := c.SayHello(ctx, &chat.Message{Body: "Hello from client"})
	log.Printf("%v", msg)
}
