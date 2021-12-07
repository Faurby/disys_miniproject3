// Frontend port: 5000
// Replica#1 port: 6001
// Replica#2 port: 6002
// Replica#3 port: 6003

package main

import (
	"bufio"
	"context"
	"disys_miniproject3/pb"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

var (
	client pb.FrontendClient
)

func main() {

	log.Println("Connecting to frontend on port 5000")
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	client = pb.NewFrontendClient(conn)

	fmt.Println("Type 'increment' to increment by 1\n" +
		"Otherwise type 'increment' followed by a value.\n")

	go terminalInput()

	// block main
	bl := make(chan bool)
	<-bl
}

func terminalInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		clientMessage = strings.TrimRight(clientMessage, "\r\n")
		if err != nil {
			log.Printf("Failed to read from console : %v", err)
			continue
		}

		if strings.HasPrefix(clientMessage, "increment") {
			increment := strings.Split(clientMessage, " ")
			if len(increment) > 1 {
				bidValue, _ := strconv.Atoi(increment[1])
				response := incrementToFrontend(int32(bidValue))
				fmt.Printf("The value is this now: %d", response)
			} else {
				response := incrementToFrontend(int32(1))
				fmt.Printf("The value is this now: %d", response)
			}
		}
	}
}

func incrementToFrontend(value int32) int32 {
	response, _ := client.Increment(context.Background(), &pb.IncRequest{Amount: value})

	return response.NewAmount
}
