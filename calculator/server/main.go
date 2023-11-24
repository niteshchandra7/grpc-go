package main

import (
	"log"
	"net"

	pb "github.com/niteshchandra7/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

type Server struct {
	pb.CalculateServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculateServiceServer(s, &Server{})
	log.Printf("Listening on %v\n", addr)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
}
