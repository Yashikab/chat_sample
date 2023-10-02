package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"

	pb "github.com/Yashikab/chat_sample/chat_protobuf"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:8400", "the address to connect to")
	// name = flag.String("name", "Yashi", "Name to greet")
)

func sendMessage(c pb.HelloGrpcClient, id string, user string) error {
	stdin := bufio.NewScanner(os.Stdin)
	stream, err := c.SendMessage(context.Background())
	if err != nil {
		return err
	}
	for {
		stdin.Scan()
		text := stdin.Text()
		if err := stream.Send(&pb.SendRequest{Id: id, User: user, Content: text}); err != nil {
			log.Fatalf("Send failed: %v", err)
		}
		if text == "/exit" {
			break
		}
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}

func getMessage(c pb.HelloGrpcClient, id string) error {
	req := &pb.MessagesRequest{Id: id}
	stream, err := c.GetMessage(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("[%s] %s", msg.User, msg.Content)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	// set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloGrpcClient(conn)

	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "send":
			err := sendMessage(c, os.Args[2], os.Args[3])
			if err != nil {
				log.Fatalf("Couldn't execute: %v", err)
			}
		case "stream":
			err := getMessage(c, os.Args[2])
			if err != nil {
				log.Fatalf("Couldn't execute: %v", err)
			}
		case "chat":
			go getMessage(c, os.Args[2])
			err := sendMessage(c, os.Args[2], os.Args[3])
			if err != nil {
				log.Fatalf("Couldn't execute: %v", err)
			}
		default:
			log.Fatalf("Unknown command.")
		}
	} else {
		log.Fatalf("Need arguments.")
	}
}
