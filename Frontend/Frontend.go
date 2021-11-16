package main

import (
	Frontend "disys_miniproject3/Frontend"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type Server struct {
	Frontend.UnimplementedFrontendServer
}

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen on port 5000: %v", err)
	}
	log.Println(":5000 is listening")

	// Creates empty gRPC server
	grpcServer := grpc.NewServer()

	// Creates instance of our ChittyChatServer struct and binds it with our empty gRPC server.
	ccs := Server{}
	Frontend.RegisterFrontendServer(grpcServer, &ccs)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to start gRPC server :: %v", err)
	}
}
