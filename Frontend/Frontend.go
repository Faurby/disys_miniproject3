package main

import (
	AuctionFrontend "disys_miniproject3/AuctionFrontend"
	AuctionBackend "disys_miniproject3/AuctionBackend"
	"log"
	"net"
	"context"

	grpc "google.golang.org/grpc"
)

type Server struct {
	AuctionFrontend.UnimplementedFrontendServer
}

var frontendClientToBackend AuctionBackend.ReplicaClient

//a hashmap of userid's ??? 

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen on port 5000: %v", err)
	}
	log.Println(":5000 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our Frontend Server struct and binds it with our empty gRPC server.
	ccs := Server{}
	AuctionFrontend.RegisterFrontendServer(grpcServer, &ccs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}

	//Establishes connection to a Replica server, localhost:6001?
	log.Println("Connecting to backendend on port 6001")
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	frontendClientToBackend = AuctionBackend.NewReplicaClient(conn)

}

func Bid(bidReq *AuctionFrontend.BidRequest) AuctionFrontend.BidResponse {
	
	response, _ := frontendClientToBackend.Bid(context.Background(), &AuctionBackend.BidReplicaRequest{Amount: bidReq.Amount})

	return AuctionFrontend.BidResponse{Ack:response.Ack}
}


func Result() AuctionFrontend.ResultResponse {
	response, _ := frontendClientToBackend.Result(context.Background(), &AuctionBackend.Empty{})

	return AuctionFrontend.ResultResponse{Result: response.Result}
}
