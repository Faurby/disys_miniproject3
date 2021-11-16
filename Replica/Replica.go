package main

import (
	AuctionBackend "disys_miniproject3/AuctionBackend"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type Server struct {
	AuctionBackend.UnimplementedReplicaServer
}

var maxBid int32 = 0

func main() {
	listen, err := net.Listen("tcp", ":6001")
	if err != nil {
		log.Fatalf("Failed to listen on port 6001: %v", err)
	}
	log.Println(":6001 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our Backend Server struct and binds it with our empty gRPC server.
	ccs := Server{}
	AuctionBackend.RegisterReplicaServer(grpcServer, &ccs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}

func Bid(bidReq *AuctionBackend.BidReplicaRequest) AuctionBackend.BidReplicaResponse {
	if bidReq.Amount > maxBid {
		maxBid = bidReq.Amount
		return AuctionBackend.BidReplicaResponse{Ack: "success"}
	} else if bidReq.Amount <= maxBid {
		return AuctionBackend.BidReplicaResponse{Ack: "fail"}
	}
	return AuctionBackend.BidReplicaResponse{Ack: "exception"}
}

func Result() AuctionBackend.ResultReplicaResponse {

	return AuctionBackend.ResultReplicaResponse{Result: maxBid}
}
