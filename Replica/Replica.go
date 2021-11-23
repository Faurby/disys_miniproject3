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
	HighestBidMap map[uint32]int32
	auctionActive bool
)

func main() {
	auctionActive = true
	HighestBidMap = make(map[uint32]int32)
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
	if os.Args[1] == ":6001" {
		go countdownToDeath()
	}

	go countdown()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

func (c *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidReplicaResponse, error) {

	if !auctionActive {
		return &pb.BidReplicaResponse{Ack: "exception"}, nil
	}
	// Find current max bid of all participants
	_, maxBid := getMax()

	// If clients incoming bid is higher than previous bid, set currentbid to new max
	currentBid := HighestBidMap[in.UserID]
	if currentBid < in.Amount {
		HighestBidMap[in.UserID] = in.Amount
	}

	fmt.Printf("------\nIncoming bid: %d\nComparing to highest bid: %d\n", in.Amount, maxBid)

	if in.Amount > maxBid {
		fmt.Printf("SUCCESS, incoming bid is higher than current max\n\n")
		return &pb.BidReplicaResponse{Ack: "success"}, nil
	} else {
		fmt.Printf("FAIL, incoming bid is NOT higher than current max\n\n")
		return &pb.BidReplicaResponse{Ack: "fail"}, nil
	}
}

func (c *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {

	userID, max := getMax()

	var leader, state bool

	if in.UserID == userID {
		leader = true
	}

	if auctionActive {
		state = true
	}

	return &pb.ResultResponse{Result: max, Leader: leader, State: state}, nil
}

func getMax() (userID uint32, max int32) {
	for user, v := range HighestBidMap {
		if v > max {
			max = v
			userID = user
		}
	}
	return
}

func countdownToDeath() {
	time.Sleep(10 * time.Second)
	os.Exit(1)
}

func countdown() {
	time.Sleep(10 * time.Second)
	auctionActive = false
	return
}
