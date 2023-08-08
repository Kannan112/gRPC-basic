package main

import (
	"log"
	"net"

	pb "github.com/kannan112/demo-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloserver struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to port 8080 %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start :%v", err)
	}
}
