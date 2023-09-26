package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/Yashikab/chat_sample/chat_protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8400", "the address to connect to")
	// name = flag.String("name", "Yashi", "Name to greet")
)

func main() {
	flag.Parse()
	// set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var name string
	fmt.Print("Type your Name > ")
	_, err = fmt.Scan(&name)
	if err != nil {
		log.Fatalf("could not type input: %v", err)
	}

	r, err := c.GreetServer(ctx, &pb.GreetRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Msg)
}
