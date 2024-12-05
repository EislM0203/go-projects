package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, blogID string) *pb.Blog{
	log.Println("Reading the blog...")

	res, err := c.ReadBlog(context.Background(), &pb.BlogId{Id: blogID})
	if err != nil {
		log.Fatalf("Failed to read blog: %v\n", err)
	}

	log.Printf("Blog has been read: %v\n", res.String())
	return res
}