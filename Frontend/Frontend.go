package main

import (
	"context"
	"disys_miniproject3/pb"
	"fmt"
	"log"
	"net"
	"reflect"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFrontendServer
}

var (
	replicas []pb.ReplicaClient
)

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen on port 5000: %v", err)
	}
	log.Println(":5000 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our struct and binds it with our empty gRPC server.
	ccs := Server{}
	pb.RegisterFrontendServer(grpcServer, &ccs)

	// Set up client connection to Replicas
	go setupReplicaConnection("localhost:6001")
	go setupReplicaConnection("localhost:6002")
	go setupReplicaConnection("localhost:6003")

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

	replicas = append(replicas, pb.NewReplicaClient(conn))
	// block main
	bl := make(chan bool)
	<-bl
}

func (c *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {

	// Send bid to all Replicas
	var responses []*pb.BidReplicaResponse
	for _, v := range replicas {
		response, _ := v.Bid(context.Background(), in)
		if response != nil {
			responses = append(responses, response)
		}
	}
	fmt.Printf("---- Responses from Replicates ----\n%v\n\n", responses)

	// if Replicas have agreed
	ok, ack := checkResponses(responses)
	if ok {
		return &pb.BidResponse{Ack: ack}, nil
	}
	return &pb.BidResponse{Ack: "exception"}, nil
}

func (c *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {
	// Send bid to all Replicas
	var responses []*pb.ResultResponse
	for _, v := range replicas {
		response, _ := v.Result(context.Background(), in)
		if response != nil {
			responses = append(responses, response)
		}
	}

	fmt.Printf("---- Responses from Replicates ----\n%v\n\n", responses)

	// Replicas have agreed
	return checkResultResponses(responses), nil
}

func checkResultResponses(r []*pb.ResultResponse) *pb.ResultResponse {

	if reflect.DeepEqual(r[0], r[1]) {
		fmt.Println("2 replicas agree out of 2")
		return r[0]

	} else if len(r) > 2 {
		if reflect.DeepEqual(r[1], r[2]) || reflect.DeepEqual(r[0], r[2]) {
			fmt.Println("2 replicas agree out of 3")
			return r[2]
		}
	}
	fmt.Println("No replicas agree man :(")
	return &pb.ResultResponse{Result: -1}
}

func checkResponses(r []*pb.BidReplicaResponse) (bool, string) {

	bidResponseMap := make(map[string]int32)

	for _, v := range r {
		bidResponseMap[v.Ack]++
	}

	var max int32 = 0
	var answer string
	for ack, count := range bidResponseMap {
		if count > max {
			max = count
			answer = ack
		}
	}
	if max > 1 {
		return true, answer
	}
	return false, "exception"
}
