package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go_grpc/proto"
)

const (
	port = ":8080"
)

func main() {
	// Create a connection to the server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}
	// callSayHello(client) - Unary call
	callSayHelloServerStreaming(client,names)
}
