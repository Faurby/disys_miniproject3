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
	client1, client2 pb.FrontendClient
)

func main() {
	conn1, _ := grpc.Dial("localhost:5000", grpc.WithInsecure())
	conn2, _ := grpc.Dial("localhost:5001", grpc.WithInsecure())

	defer conn1.Close()
	defer conn2.Close()

	client1 = pb.NewFrontendClient(conn1)
	client2 = pb.NewFrontendClient(conn2)

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client1.Increment(ctx, &pb.IncRequest{Amount: value})

	if err != nil {
		var err2 error
		response, err2 = client2.Increment(ctx, &pb.IncRequest{Amount: value})

		if err2 != nil {
			fmt.Println("Both frontends down!")
			return -1
		} else {
			fmt.Println("Frontend#1 down, using frontend#2")
		}
	}

	return response.NewAmount
}
