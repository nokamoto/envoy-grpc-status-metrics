package main

import (
	"log"
	"net"

	pb "github.com/nokamoto/envoy-grpc-status-metrics/internal/protobuf"
	"google.golang.org/grpc"
)

func main() {
	port := ":9000"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen: %v\n", port)

	s := grpc.NewServer()
	pb.RegisterServerServer(s, &pb.UnimplementedServerServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
