package main

import (
	"context"
	"log"
	"time"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GreetWithDeadline(c pb.GreetServiceClient, timeout int64) {
	log.Printf("GreetWithDeadline was invoked\n")

	req := &pb.GreetRequest{FirstName: "Markus"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout was hit! Deadline was exceeded\n")
			} else {
				log.Fatalf("Error %v: %v\n", e.Code(), e.Message())
			}
		} else {
			log.Fatalf("Error: %v\n", err)
		}
		return
	}

	log.Printf("Response: %v\n", res.Result)
}