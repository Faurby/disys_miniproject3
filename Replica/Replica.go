package main

import (
	"context"
	"disys_miniproject3/pb"
	"log"
	"net"
	"os"
	"time"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedReplicaManagerServer
}

var (
	amount int32
)

func main() {

	amount = -1

	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", os.Args[1], err)
	}
	log.Printf("%s is listening", os.Args[1])

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our ChittyChatServer struct and binds it with our empty gRPC server.
	ccs := Server{}
	pb.RegisterReplicaManagerServer(grpcServer, &ccs)

	// For killing a Replica
	if os.Args[1] == ":6001" {
		go countdownToDeath()
	}
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

func (c *Server) Increment(ctx context.Context, in *pb.IncRequest) (*pb.IncResponse, error) {

	if amount == -1 {
		amount = 0
	} else {
		amount += in.Amount
	}

	return &pb.IncResponse{NewAmount: amount}, nil
}

func countdownToDeath() {
	time.Sleep(30 * time.Second)
	//os.Exit(1)
}
