package main

import (
	"log"

	pb "example.com/greet-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50161"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

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