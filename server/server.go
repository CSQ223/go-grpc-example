package main

import (
	"context"
	"log"
	"net"

	pb "grpc-example/StopTimeOpt"

	"google.golang.org/grpc"
)

const (
	port = ":4999"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedMetroOptServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) StopTimeOpt(ctx context.Context, in *pb.StopTimeOptRequest) (*pb.StopTimeOptResponse, error) {
	log.Printf("Received:  min time = %v, max time = %v", in.GetMinTime(), in.GetMaxTime())

	return &pb.StopTimeOptResponse{ResultCode: 0, StopTime: 25}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMetroOptServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
