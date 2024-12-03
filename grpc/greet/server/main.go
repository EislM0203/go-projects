package main

import (
	"log"
	"net"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr = "localhost:50161"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}