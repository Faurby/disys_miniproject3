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

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	userID uint32
	client pb.FrontendClient
)

func main() {

	userID = uuid.New().ID()
	log.Println("Connecting to frontend on port 5000")
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect gRPC server :: %v", err)
	}
	defer conn.Close()

	client = pb.NewFrontendClient(conn)

	fmt.Println("Hello and welcome to the marvelous auction house!\n" +
		"Here you can either bid or query the result.\n" +
		"You can either type'bid amount' in the current item, or type 'result'\n")

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

		if strings.HasPrefix(clientMessage, "bid") {
			bids := strings.Split(clientMessage, " ")
			if len(bids) > 1 {
				bidValue, _ := strconv.Atoi(bids[1])
				response := bidToFrontend(int32(bidValue))
				switch response {
				case "success":
					fmt.Println("Congratulations, you have the highest bid :)")

				case "fail":
					fmt.Println("Your bid was not high enough, yikes :(")

				case "exception":
					fmt.Println("The auction has ended :/")

				}
			} else {
				fmt.Println("Please add amount")
			}
		} else if strings.EqualFold(clientMessage, "result") {
			result()
		}
	}
}

func bidToFrontend(value int32) string {
	response, _ := client.Bid(context.Background(), &pb.BidRequest{UserID: userID, Amount: value})

	return response.Ack
}

func result() {
	response, _ := client.Result(context.Background(), &pb.ResultRequest{UserID: userID})

	if response.State {
		fmt.Println("The auction is ongoing")

		if response.Leader {
			fmt.Printf("You're in the lead with current bid: %d\n", response.Result)
		} else {
			fmt.Printf("Your bid is not the highest. The highest bid is: %d\n", response.Result)
		}
	} else {
		fmt.Println("The auction is over")
		if response.Leader {
			fmt.Printf("You're the winner!! with final bid: %d\n", response.Result)
		} else {
			fmt.Printf("You didnt win the auction :( The winning bid was: %d\n", response.Result)
		}
	}
}
