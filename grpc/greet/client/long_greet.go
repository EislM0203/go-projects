package main

import (
	"context"
	"log"
	"time"

	pb "example.com/greet-grpc/greet/proto"
)

func LongGreet(c pb.GreetServiceClient) {
	log.Println("Sending LongGreet")

	reqs := []*pb.GreetRequest{
		{FirstName: "Markus"},
		{FirstName: "Nikolaus"},
		{FirstName: "Lea"},
		{FirstName: "Peter"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Failed to call LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending: %v\n", req.FirstName)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive: %v", err)
	}
	
	log.Printf("Response: %v\n", res.Result)
}