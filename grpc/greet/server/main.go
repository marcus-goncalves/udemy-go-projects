package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "udemy-projects.com/grpc/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to list on %v\n", addr)
	}

	fmt.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed on serve: %v\n", err)
	}
}
