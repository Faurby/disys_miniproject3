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
	clients       []pb.ReplicaClient
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

	// Set up client connection to Replicas
	go setupReplicaConnection("localhost:6001")
	go setupReplicaConnection("localhost:6002")
	go setupReplicaConnection("localhost:6003")

	go countdown()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
	fmt.Println("testtesttest this is after server skrrttt")
}

func setupReplicaConnection(ip string) {
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	fmt.Println("Connected to: " + ip)

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	clients = append(clients, pb.NewReplicaClient(conn))
	// block main
	bl := make(chan bool)
	<-bl
}

func (c *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	clientID := in.UserID
	clientBid := in.Amount
	output := &pb.BidResponse{}

	if !auctionActive {
		return &pb.BidResponse{Ack: "exception"}, nil
	}

	// Send bid to all Replicas
	var responses []*pb.BidReplicaResponse
	for _, v := range clients {
		response, _ := v.Bid(context.Background(), &pb.BidReplicaRequest{Amount: clientBid})
		responses = append(responses, response)
	}
	fmt.Printf("---- Responses from Replicates ----\n%v\n\n", responses)

	// Replicas have agreed
	if ok, ack := checkResponses(responses); ok {
		output = &pb.BidResponse{Ack: ack}
	}

	// if client exists in array
	if _, ok := HighestBidMap[clientID]; ok {

		fmt.Println(HighestBidMap)
		HighestBidMap[clientID] = clientBid
		return output, nil
		// else add client to map with clientBid
	} else {

		HighestBidMap[clientID] = clientBid
		fmt.Println(HighestBidMap)

		return output, nil
	}
}

func (c *Server) Result(ctx context.Context, in *pb.Empty) (*pb.ResultResponse, error) {
	// Send bid to all Replicas
	var responses []*pb.ResultReplicaResponse
	for _, v := range clients {
		response, _ := v.Result(context.Background(), &pb.Empty{})
		responses = append(responses, response)
	}
	fmt.Printf("---- Responses from Replicates ----\n%v\n\n", responses)

	// Replicas have agreed
	if ok, result := checkResultResponses(responses); ok {
		return &pb.ResultResponse{Result: result}, nil
	}
	return &pb.ResultResponse{Result: -1}, nil
}

func checkResultResponses(r []*pb.ResultReplicaResponse) (bool, int32) {

	// If three are equal
	if r[0].Result == r[1].Result && r[1].Result == r[2].Result {
		return true, r[0].Result
	}

	// If two are equal
	if r[0].Result == r[1].Result {
		return true, r[0].Result

	} else if r[1].Result == r[2].Result {
		return true, r[1].Result

	} else if r[0].Result == r[2].Result {
		return true, r[2].Result

	}
	// None are equal, return false
	return false, -1
}

func checkResponses(r []*pb.BidReplicaResponse) (bool, string) {

	// If three are equal
	if r[0].Ack == r[1].Ack && r[1].Ack == r[2].Ack {
		return true, r[0].Ack
	}

	// If two are equal
	if r[0].Ack == r[1].Ack {
		return true, r[0].Ack

	} else if r[1].Ack == r[2].Ack {
		return true, r[1].Ack

	} else if r[0].Ack == r[2].Ack {
		return true, r[2].Ack

	}
	// None are equal, return false
	return false, "exception"
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

func countdown() {
	time.Sleep(20 * time.Second)
	auctionActive = false
	return
}
