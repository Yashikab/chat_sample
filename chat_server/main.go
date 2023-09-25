package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Yashikab/chat_sample/chat_protobuf"
)

type server struct{}

func (s *server) GreetServer(ctx context.Context, p *pb.GreetRequest) (*pb.GreetMessage, error) {
	log.Printf("Request from: %s", p.Name)
	return &pb.GreetMessage{Msg: fmt.Sprintf("Hello, %s", p.Name)}, nil
}

func main() {
	port := 8400
	host := "localhost"
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("faied to listen: %v", err)
	}
	log.Printf("Run server port: %d", port)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloGrpcServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
