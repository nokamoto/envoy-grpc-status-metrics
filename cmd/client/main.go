package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nokamoto/envoy-grpc-status-metrics/internal/protobuf"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:9000"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Say(ctx, &pb.Value{})
	if err != nil {
		log.Fatalf("could not say: %v", err)
	}
	log.Printf("Say: %v", res)
}
