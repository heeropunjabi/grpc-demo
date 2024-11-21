package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-demo/pb" // Replace with the actual path to the generated proto package

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.GreetResponse{Message: "Hello, " + req.GetName()}, nil
}

func StartServer() {
	fmt.Println("Server is starting...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
