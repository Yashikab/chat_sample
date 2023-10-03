package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Yashikab/chat_sample/chat_protobuf"
)

type server struct {
	contentMap map[string](chan pb.Message)
	pb.UnimplementedHelloGrpcServer
}

func (s *server) GreetServer(ctx context.Context, p *pb.GreetRequest) (*pb.GreetMessage, error) {
	log.Printf("Request from: %s", p.Name)
	return &pb.GreetMessage{Msg: fmt.Sprintf("Hello, %s", p.Name)}, nil
}

func (s *server) SendMessage(stream pb.HelloGrpc_SendMessageServer) error {
	for {
		m, err := stream.Recv()
		log.Printf("Recieve message >> [%s] %s", m.User, m.Content)
		if _, ok := s.contentMap[m.Id]; ! ok {
			s.contentMap[m.Id] = make(chan pb.Message)
		}
		s.contentMap[m.Id] <- pb.Message{Id: m.Id, User: m.User, Content: m.Content}
		if err == io.EOF {
			return stream.SendAndClose(&pb.SendResult{Result: "true"})
		}
		if err != nil {
			return err
		}
		if m.Content == "/exit" {
			return stream.SendAndClose(&pb.SendResult{Result: "true"})
		}
	}
}

func (s *server) GetMessage(p *pb.MessagesRequest, stream pb.HelloGrpc_GetMessageServer) error {

	for {
		msg := <- s.contentMap[p.Id]
		if err := stream.Send(&pb.Message{Id: msg.Id, User: msg.User, Content: msg.Content}); err != nil {
			return err
		}
	}
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
	pb.RegisterHelloGrpcServer(grpcServer, &server{contentMap: make(map[string](chan pb.Message))})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
