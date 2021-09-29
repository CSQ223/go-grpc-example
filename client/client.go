package main

import (
	"context"
	"log"
	"time"

	pb "grpc-example/StopTimeOpt"

	"google.golang.org/grpc"
)

const (
	address  = "localhost:4999"
	min_time = 25
	max_time = 35
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMetroOptClient(conn)

	// Contact the server and print out its response.
	min_time := min_time

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopTimeOpt(ctx, &pb.StopTimeOptRequest{MinTime: int32(min_time)})
	if err != nil {
		log.Fatalf("could not opt stop time: %v", err)
	}
	log.Printf("Optimization Code: %v, Result: %v", r.GetResultCode(), r.GetStopTime())
}
