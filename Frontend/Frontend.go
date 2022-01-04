package main

import (
	"context"
	"disys_miniproject3/pb"
	"fmt"
	"log"
	"net"
	"os"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFrontendServer
}

var (
	replicas []pb.ReplicaManagerClient
)

func main() {
	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	log.Printf("%v is listening", os.Args[1])

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
}

func setupReplicaConnection(ip string) {
	conn, err := grpc.Dial(ip, grpc.WithInsecure())
	fmt.Println("Connected to: " + ip)

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	replicas = append(replicas, pb.NewReplicaManagerClient(conn))
	// block main
	bl := make(chan bool)
	<-bl
}

func (c *Server) Increment(ctx context.Context, in *pb.IncRequest) (*pb.IncResponse, error) {

	// Send increment request to all Replicas
	var responses []*pb.IncResponse
	for _, v := range replicas {
		response, _ := v.Increment(context.Background(), in)
		if response != nil {
			responses = append(responses, response)
		}
	}
	fmt.Printf("---- Responses from Replicates ----\n%v\n\n", responses)

	// if Replicas have agreed
	ok, amount := checkResponses(responses)
	if ok {
		return &pb.IncResponse{NewAmount: amount}, nil
	}
	return &pb.IncResponse{NewAmount: -1}, nil
}

func checkResponses(r []*pb.IncResponse) (bool, int32) {

	incResponseMap := make(map[int32]int32)

	for _, v := range r {
		incResponseMap[v.NewAmount]++
	}

	var max int32 = 0
	var answer int32
	for ans, count := range incResponseMap {
		if count > max {
			max = count
			answer = ans
		}
	}
	if max > 1 {
		return true, answer
	}
	return false, -1
}
