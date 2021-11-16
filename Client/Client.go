// Frontend port: 5000
// Replica#1 port: 6001
// Replica#2 port: 6002
// Replica#3 port: 6003

package main

import (
	"bufio"
	"context"
	AuctionFrontend "disys_miniproject3/AuctionFrontend"
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
	client AuctionFrontend.FrontendClient
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

	client = AuctionFrontend.NewFrontendClient(conn)

	fmt.Println("Hello and welcome to the marvelous auction house!\n" +
		"Here you can either bid or query the result.\n" +
		"You can either type'bid amount' in the current item, or type 'result'\n")

	terminalInput()
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

				if bidValue > int(maxBid) {
					maxBid = int32(bidValue)

					response := bidToFrontend(int32(bidValue))
					if strings.EqualFold(response, "success") {
						fmt.Println("Congratulations, you have the highest bid :)")
					} else if strings.EqualFold(response, "fail") {
						fmt.Println("Your bid was not high enough, yikes :(")
					} else if strings.EqualFold(response, "exception") {
						fmt.Println("Uff, something went wrong :/ Please bid an integer.")
					}
				} else {
					fmt.Println("You have to bid higher than your previous bid.")
				}
			}
		} else if strings.EqualFold(clientMessage, "result") {
			fmt.Printf("The highest bidder / result : %d", resultFromFrontend())
		}
	}
}

func bidToFrontend(value int32) string {
	response, _ := client.Bid(context.Background(), &AuctionFrontend.BidRequest{UserID: userID, Amount: value})

	return response.Ack
}

func resultFromFrontend() int32 {
	response, _ := client.Result(context.Background(), &AuctionFrontend.Empty{})

	return response.Result
}
