package main

import (
	"io"
	"log"

	pb "example.com/greet-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Received LongGreet")

	result := "Hello, "

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: result})
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
			break
		}

		result += req.FirstName + "! "
	}

	return stream.SendAndClose(&pb.GreetResponse{Result: result})
}