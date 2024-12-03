package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "example.com/greet-grpc/greet/proto"
)

func GreetEveryone(c pb.GreetServiceClient) {
	log.Printf("GreetEveryone was invoked\n")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Failed to greet everyone: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Markus"},
		{FirstName: "Eva"},
		{FirstName: "John"},
		{FirstName: "Jane"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending: %v\n", req.FirstName)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
				close(waitc)
				return
			}
			log.Printf("Received: %v\n", res.Result)
		}
	}()

	<-waitc
}