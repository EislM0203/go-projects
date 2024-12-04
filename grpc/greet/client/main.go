package main

import (
	"log"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50161"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		creds, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	/*rsp, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "Markus"})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Response: %v\n", rsp.Result)*/

	//GreetManyTimes(client)
	//LongGreet(client)
	//GreetEveryone(client)
	GreetWithDeadline(client, 1)
}