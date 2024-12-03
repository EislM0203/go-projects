package main

import (
	"log"

	pb "example.com/greet-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Received: %v\n", in.FirstName)

	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.GreetResponse{Result: "Hello, " + in.FirstName}); err != nil {
			return err
		}
	}

	return nil
}