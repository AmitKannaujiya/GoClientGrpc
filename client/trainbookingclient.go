package client

import (
	"context"
	pb "grpc-client/proto/api.v1"
	u "grpc-client/utill"
	"time"

	"log"

	"google.golang.org/grpc"
)


func SetupTrainBookingClient() {
	connection, err := grpc.Dial(u.BASE_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect :  %v", err)
	}
	defer connection.Close()
	client := pb.NewTrainBookingClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	respBooking, er := client.BookTicket(ctx, &pb.TicketRequest{Departure : "London", Destination : "France", 
	FirstName : "Amit", LastName : "Kumar", Price : 20.50, Email : "amit@abc.com"})

	if er != nil {
		log.Fatalf("Cound not book the ticket : %v", er)
	}
	log.Println("booking confirmation %s, Message : %s", respBooking.GetConfirmation(), respBooking.GetMessage())

	// check booking status 
	status, e :=  client.GetBookingStatus(ctx, &pb.BookingStatusRequest{Confirmation : respBooking.GetConfirmation() })

	if e != nil {
		log.Fatalf("could not get booking status : %v", e)
	}
	log.Printf("Booking status : %s", status.GetStatus())
}
