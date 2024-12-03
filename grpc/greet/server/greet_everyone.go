package main

import (
	"fmt"
	"io"
	"log"

	pb "example.com/greet-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Received GreetEveryone")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
			return err
		}
		
		res := fmt.Sprintf("Hello, %v!", req.FirstName)
		if err := stream.Send(&pb.GreetResponse{Result: res}); err != nil {
			log.Fatalf("Failed to send: %v", err)
			return err
		}
	}

}