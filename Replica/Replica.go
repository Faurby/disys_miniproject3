package main

import (
	"context"
	"disys_miniproject3/pb"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedReplicaServer
}

var (
	highestBid int32
)

func main() {
	highestBid = 0
	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", os.Args[1], err)
	}
	log.Printf("%s is listening", os.Args[1])

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our ChittyChatServer struct and binds it with our empty gRPC server.
	ccs := Server{}
	pb.RegisterReplicaServer(grpcServer, &ccs)

	// For killing a Replica
	// if os.Args[1] == ":6001" {
	// 	go countdownToDeath()
	// }

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

func (c *Server) Bid(ctx context.Context, in *pb.BidReplicaRequest) (*pb.BidReplicaResponse, error) {
	clientBid := in.Amount

	fmt.Printf("------\nIncoming bid: %d\nComparing to highest bid: %d\n", clientBid, highestBid)

	if clientBid > highestBid {
		highestBid = clientBid
		fmt.Printf("Success, incoming bid is higher (return success)\n\n")
		return &pb.BidReplicaResponse{Ack: "success"}, nil
	} else {
		fmt.Printf("fail, incoming bid is higher (return fail)\n\n")
		return &pb.BidReplicaResponse{Ack: "fail"}, nil
	}
}

func (c *Server) Result(ctx context.Context, in *pb.Empty) (*pb.ResultReplicaResponse, error) {

	return &pb.ResultReplicaResponse{Result: getMax()}, nil
}

func getMax() int32 {

	return highestBid
}

func countdownToDeath() {
	time.Sleep(4 * time.Second)
	os.Exit(1)
}
