package main

import (
	"context"
	"io"
	"log"

	pb "example.com/greet-grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("Listing the blogs...")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Failed to list blog: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v\n", err)
		}

		log.Printf("Blog has been read: %v\n", res.String())
	}
}