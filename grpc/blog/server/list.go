package main

import (
	"context"
	"log"

	pb "example.com/greet-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("Received ListBlog")

	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Failed to fetch blogs: %v", err,
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Error while decoding stuff from mongodb %v", err,
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error %v", err,
		)
	}

	return nil
}