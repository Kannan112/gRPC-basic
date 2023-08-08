package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kannan112/demo-grpc/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal("could not greet:%v", err)
	}
	log.Printf("%s", res.Message)
}
