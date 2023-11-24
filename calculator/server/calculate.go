package main

import (
	"context"
	"log"

	pb "github.com/niteshchandra7/grpc-go/calculator/proto"
)

func (*Server) Calculate(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Println("Calculate function was invoked")
	return &pb.CalculateResponse{
		Result: in.FirstNum + in.SecondNum,
	}, nil
}
