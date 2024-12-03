package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v\n", in.FirstName)
	return &pb.GreetResponse{Result: "Hello, " + in.FirstName}, nil
}