package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status")

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Received ReadBlog request for blog %v\n", in.Id)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with specified ID",
		)
	}
	return documentToBlog(data), nil
}
