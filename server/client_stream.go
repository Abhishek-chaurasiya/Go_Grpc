package main

import (
	pb "go_grpc/proto"
	"io"
	"log"
)

func (s *server) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	// This function will be called for each message sent by the client
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", req.Name)
		messages = append(messages, req.Name)
	}
}
