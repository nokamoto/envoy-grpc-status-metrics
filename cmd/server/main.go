package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nokamoto/envoy-grpc-status-metrics/internal/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedServerServer
}

func (server) Say(_ context.Context, v *pb.Value) (*pb.Value, error) {
	log.Printf("Say(%v)", v)
	if codes.OK == codes.Code(v.GetStatus()) {
		return v, nil
	}
	return nil, status.Errorf(codes.Code(v.GetStatus()), "code=%v", v.GetStatus())
}

func main() {
	port := ":9000"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen: %v\n", port)

	s := grpc.NewServer()
	pb.RegisterServerServer(s, server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
