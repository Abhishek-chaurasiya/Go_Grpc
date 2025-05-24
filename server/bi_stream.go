package main

import (
	pb "go_grpc/proto"
	"io"
	"log"
)

func (s *server) SayHelloBidirectionalStreaming(
	stream pb.GreetService_SayHelloBidirectionalStreamingServer,
) error {
	for {
		// Receive a message from the client
		name, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received name: %s", name.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + name.Name,
		}

		if err := stream.Send(res); err != nil {
			log.Printf("Error sending response: %v", err)
			return err
		}

	}
}
