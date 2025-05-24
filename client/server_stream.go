package main

import (
	"context"
	pb "go_grpc/proto"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
    log.Println("streaming started...")
	// Call the SayHelloServerStreaming method
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("erro sending names: %v", err)
	}

	// Receive messages from the server
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			// End of the stream
			log.Println("stream ended.")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		log.Printf("Received message: %s", message.Message)
	}
}
