package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, blogID string) {
	log.Println("Deleting the blog...")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: blogID})
	if err != nil {
		log.Fatalf("Failed to delete blog: %v\n", err)
	}

	log.Printf("Blog has been deleted %v\n", blogID)
}