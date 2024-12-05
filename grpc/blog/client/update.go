package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("Updating the blog...")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Markus",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	log.Println("Updating blog with id: ", id)
	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Failed to update blog: %v", err)
	}
}