package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/niteshchandra7/grpc-go/calculator/proto"
)

func doCalculation(c pb.CalculateServiceClient) {
	res, err := c.Calculate(context.Background(), &pb.CalculateRequest{
		FirstNum:  3,
		SecondNum: 10,
	})
	if err != nil {
		log.Printf("Calculate Failed: %v\n", err)
	}
	fmt.Printf("Result is: %v\n", res.Result)
}
