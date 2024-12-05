package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Received: %v\n", in)

	data := BlogItem {
		AuthorID: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatalf("Failed to insert: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			"Failed to insert: %v", err,
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to convert to OID",
		)
	}
	return &pb.BlogId{Id: oid.Hex()}, nil
}