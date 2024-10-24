package main

import (
	"fmt"
	client "grpc-client/client"
)

func main()  {
	fmt.Println("Hello World")
	client.SetupClient()
	client.SetupTrainBookingClient()
}