package main

import (
	"log"
	"net"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}