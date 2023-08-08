package main

import (
	"log"

	pb "github.com/kannan112/demo-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect 8080 %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// name := &pb.NameList{
	// 	Names:[]string{"Akil", "Alice", "Bob"},
	// }
	CallSayHello(client)
}
