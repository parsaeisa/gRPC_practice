package main

import (
	"time"
	"context"
	"log"

	grpc "google.golang.org/grpc"
	pb "github.com/parsaeisa/gRPC_practice/Greeter"
)

const (
	address ="localhost:50051"
)

func main () {
	//dial a connection to grpc server 
	// WithBlock means that this function will not return untill the connection is made 
	conn , err := grpc.Dial(address , grpc.WithInsecure() , grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %" , err)
	}

	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel ()

	HelloReply , err := client.SayHello(ctx , &pb.HelloRequest{Name : "Hello , I'm client "})
	if err != nil {
		log.Fatalf("error while greeting %v" , err)
	}

	log.Print(HelloReply.Message)
}