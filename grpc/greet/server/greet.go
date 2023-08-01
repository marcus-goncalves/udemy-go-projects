package main

import (
	"context"
	"log"

	pb "udemy-projects.com/grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
