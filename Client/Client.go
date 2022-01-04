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
	"time"

	"google.golang.org/grpc"
)

var (
	client pb.FrontendClient
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect gRPC server 5001: %v \n Trying to connect to 5001", err)
		conn, _ = grpc.Dial("localhost:5001", grpc.WithInsecure())
		log.Println("Connected to frontend on port 5001")

	} else {
		log.Println("Connected to frontend on port 5000")

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
				fmt.Printf("The value is this now: %d\n", response)
			} else {
				response := incrementToFrontend(int32(1))
				fmt.Printf("The value is this now: %d\n", response)
			}
		}
	}
}

func incrementToFrontend(value int32) int32 {
	response, _ := client.Increment(context.Background(), &pb.IncRequest{Amount: value})

	return response.NewAmount
}
