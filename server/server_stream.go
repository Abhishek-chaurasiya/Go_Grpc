package main

import (
	pb "go_grpc/proto"
	"time"
)

func (s *server) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	for _, name := range req.Names {
		// Create a response message for each name
		response := &pb.HelloResponse{
			Message: "Hello " + name + ", this is a server streaming call!",
		}

		// Send the response to the client
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Simulate some delay between responses
	}

	return nil
}
