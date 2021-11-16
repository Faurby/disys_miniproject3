package main

import (
	"context"
	"disys_miniproject3/pb"
	"fmt"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFrontendServer
}

var (
	HighestBidMap map[uint32]int32
	auctionActive bool
)

func main() {
	HighestBidMap = make(map[uint32]int32)
	auctionActive = true
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen on port 5000: %v", err)
	}
	log.Println(":5000 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our ChittyChatServer struct and binds it with our empty gRPC server.
	ccs := Server{}
	pb.RegisterFrontendServer(grpcServer, &ccs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
	go countdown()
}

func (c *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	clientID := in.UserID
	clientBid := in.Amount

	if !auctionActive {
		return &pb.BidResponse{Ack: "exception"}, nil
	}

	// if client exists in array, check if current value is higher than incoming bid
	if val, ok := HighestBidMap[clientID]; ok {
		if val < clientBid {
			fmt.Println(HighestBidMap)
			// Send bid to all Replicas

			// wait for akc from all Replicas. If recievede 2 >= ack, then send return response
			_, max := getMax()
			HighestBidMap[clientID] = clientBid

			if clientBid > max {
				return &pb.BidResponse{Ack: "success"}, nil
			} else {
				return &pb.BidResponse{Ack: "fail"}, nil
			}
		} else {
			// Client input is not higher than last time
			return &pb.BidResponse{Ack: "fail"}, nil
		}
		// else add client to map with clientBid
	} else {
		HighestBidMap[clientID] = clientBid
		fmt.Println(HighestBidMap)
		// Send bid to all Replicas

		// wait for akc from all Replicas. If recievede 2 >= ack, then send return response
		return &pb.BidResponse{Ack: "success"}, nil
	}
}

func getMax() (uint32, int32) {
	var userID uint32 = 0
	var max int32 = 0
	for user, v := range HighestBidMap {
		if v > max {
			max = v
			userID = user
		}
	}
	return userID, max
}

func countdown() {
	time.Sleep(20 * time.Second)
	auctionActive = false
	return
}
