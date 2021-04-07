package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nokamoto/envoy-grpc-status-metrics/internal/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	address := "localhost:9001"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, code := range []codes.Code{codes.OK, codes.AlreadyExists, codes.Internal, codes.InvalidArgument} {
		res, err := c.Say(ctx, &pb.Value{
			Status: int32(code),
		})
		log.Printf("res=%v, err=%v\n", res, err)
	}
}
