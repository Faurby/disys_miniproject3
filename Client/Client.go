// Frontend port: 5000
// Replica#1 port: 6001
// Replica#2 port: 6002
// Replica#3 port: 6003

package main

import (
	"bufio"
	"context"
	"disys_miniproject3/Frontend"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	maxBid int32
	userID uint32
	client Frontend.FrontendClient
)

func main() {
	maxBid = 0

	userID = uuid.New().ID()
	log.Println("Connecting to frontend on port 5000")
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	client = Frontend.NewFrontendClient(conn)

	fmt.Println("Hello and welcome to the marvelous auction house!\n" +
		"Here you can either bid or query the result.\n" +
		"You can either type'bid amount' in the current item, or type 'result'\n")

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

		if strings.HasPrefix(clientMessage, "bid") {
			bids := strings.Split(clientMessage, " ")
			if len(bids) > 1 {
				bidValue, _ := strconv.Atoi(bids[1])
				response := bidToFrontend(int32(bidValue))
				if strings.EqualFold(response, "success") {
					fmt.Println("Congratulations, you have the highest bid :)")
				} else if strings.EqualFold(response, "fail") {
					fmt.Println("Your bid was not high enough, yikes :(")
				} else if strings.EqualFold(response, "exception") {
					fmt.Println("Uff, something went wrong :/")
				}
			}
		} else if strings.EqualFold(clientMessage, "result") {
			fmt.Printf("The highest bidder / result : %d", result())
		}
	}
}

func bidToFrontend(value int32) string {
	response, _ := client.Bid(context.Background(), &Frontend.BidRequest{UserID: userID, Amount: value})

	return response.Ack
}

func result() int32 {
	response, _ := client.Result(context.Background(), &Frontend.Empty{})

	return response.Result
}
