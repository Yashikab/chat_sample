package main

import (
	"context"

	pb "github.com/Yashikab/chat_sample/chat_protobuf"
)

type server struct{}

func (s *server) GreetServer(ctx context.Context, p *pb.GreetRequest) *pb.GreetMessage
