package main

import (
	"context"
	pb "go_grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started...")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not start bidirectional streaming: %v", err)
	}

	waitc := make(chan struct{})

	// Receive responses from the server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error receiving response: %v", err)
				break
			}
			log.Printf("Received response: %s", res.Message)
		}
		close(waitc)
	}()

	// Send names to the server
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send name %s: %v", name, err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}
	<-waitc
	log.Printf("Bidirectional streaming finished.")
}
