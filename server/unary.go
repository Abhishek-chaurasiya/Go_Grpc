package main

import (
	"context"
	pb "go_grpc/proto"
)

func (s *server) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	// Create a response message
	response := &pb.HelloResponse{
		Message: "Hello buddy, this is a unary call !",
	}

	// Return the response
	return response, nil
}
