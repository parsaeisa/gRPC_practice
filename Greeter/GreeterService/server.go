package main

import (
	"net"
	"log"
	"context"

	grpc "google.golang.org/grpc"
	pb "github.com/parsaeisa/gRPC_practice/Greeter"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello (ctx context.Context, helloRequest *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message : "Hello" ,
	} , nil
}

func main () {
	lis , err := net.Listen("tcp" , ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v" , err)
	}

	s := grpc.NewServer()
	// registering out server as a grpc server 
	pb.RegisterGreeterServer(s , &Server{})
	log.Printf("server listening at 50051")

	if err := s.Serve(lis) ; err != nil {
		log.Fatal("failed serving . ")
	}
}
