package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog...")

	blog := &pb.Blog{
		AuthorId: "Markus",
		Title: "My First Blog",
		Content: "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Failed to create blog: %v", err)
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}