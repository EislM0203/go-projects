package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/greet-grpc/blog/proto"
)

var addr string = "localhost:50162"

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)
	id := createBlog(client)
	readBlog(client, id)
	//_ = readBlog(client, "id")
	updateBlog(client, id)
	listBlog(client)
	deleteBlog(client, id)
}