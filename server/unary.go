package main

import (
	"context"

	pb "github.com/kannan112/demo-grpc/proto"
)

func (s *helloserver) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
