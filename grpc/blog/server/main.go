package main

import (
	"context"
	"log"
	"net"

	pb "example.com/greet-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "localhost:50162"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}