package main

import (
	"context"
	"log"

	pb "udemy-projects.com/grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Mark",
	})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("greeting: %s\n", res.Result)
}
