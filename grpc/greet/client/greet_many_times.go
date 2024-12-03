package main

import (
	"context"
	"io"
	"log"

	pb "example.com/greet-grpc/greet/proto"
)

func GreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("GreetManyTimes was invoked\n")

	req := &pb.GreetRequest{FirstName: "Markus"}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet many times: %v", err)
	}
	
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read: %v", err)
		}
		log.Printf("Response: %v\n", msg.Result)
	}
}