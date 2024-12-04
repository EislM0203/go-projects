package main

import (
	"context"
	"log"
	"time"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v\n", in.FirstName)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}
		time.Sleep(1 * time.Second)
		log.Printf("Processing %v", i)
	}

	return &pb.GreetResponse{Result: "Hello, " + in.FirstName}, nil
}
