package client

import (
	"context"
	pb "grpc-client/proto/api.v1"
	"log"
	"time"

	"google.golang.org/grpc"
	conf "grpc-client/cmd/config"
	u "grpc-client/utill"
)

func Execute(Config *conf.Config) {
	connection, err := grpc.Dial(Config.App.Host+":"+Config.App.Port, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect :  %v", err)
	}
	defer connection.Close()
	client := pb.NewTrainTicketingClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Purchase Ticket
	respBooking, er := client.PurchaseTicket(ctx, &pb.PurchaseRequest{From: "London", To: "France",
		User: &pb.User{
			FirstName: "Amit", LastName: "Kumar",
			Email: "amit@abc.com",
		},
		Price: 20.50, When: "2024-10-27", Section: "A"})

	if er != nil {
		log.Fatalf("Cound not book the ticket : %v", er)
	}
	log.Printf("booking confirmation %s, Message : %s\n", respBooking.GetReceiptId(), respBooking.GetDetail())

	// check booking status
	receiptResponse, egr := client.GetReceipt(ctx, &pb.ReceiptRequest{ReceiptId: respBooking.GetReceiptId()})

	if egr != nil {
		log.Fatalf("could not get booking status : %v", egr)
	}
	log.Printf("Booking status : %t\n, From : %s, To : %s, SeatNo : %s, User : %s ", receiptResponse.GetSuccess(), receiptResponse.From, receiptResponse.To, receiptResponse.Seat, receiptResponse.GetUser().String())

	//  Get Users Details based on section A
	sectionResponse, eus := client.GetUsersBySection(ctx, &pb.SectionRequest{Section: u.SECTION_A})

	if eus != nil {
		log.Fatalf("could not get users by section A : %v", eus)
	}
	for _, userSeat := range sectionResponse.Users {
		log.Printf("User  : %s, Seat : %s\n", userSeat.User.String(), userSeat.Seat)
	}
	//  Modify Seat for user
	firstUser := sectionResponse.Users[0]
	modifySeatResponse, ems := client.ModifySeat(ctx, &pb.ModifySeatRequest{Email: firstUser.User.Email, NewSeat: u.SEAT_CHANGE_REQUEST_NEW_SEAT_NO})
	if eus != nil {
		log.Fatalf("could not modify seat : %v", ems)
	}
	log.Printf("Modified seat for the user : %s , seat no : %s , response success : %t\n", firstUser.User.String(), u.SEAT_CHANGE_REQUEST_NEW_SEAT_NO, modifySeatResponse.Success)

	//  Get Users Details based on section B
	sectionResponse2, eus2 := client.GetUsersBySection(ctx, &pb.SectionRequest{Section: u.SECTION_B})

	if eus2 != nil {
		log.Fatalf("could not get users by section A : %v", eus2)
	}
	for _, userSeat := range sectionResponse2.Users {
		log.Printf("User  : %s, Seat : %s\n", userSeat.User.String(), userSeat.Seat)
	}
	// purchase ticket again
	respBooking2, er2 := client.PurchaseTicket(ctx, &pb.PurchaseRequest{From: "London", To: "France",
		User: &pb.User{
			FirstName: "Amit1", LastName: "Kumar1",
			Email: "amit1@abc.com",
		},
		Price: 50.50, When: "2024-10-28", Section: "B"})

	if er2 != nil {
		log.Fatalf("Cound not book the ticket : %v", er2)
	}
	log.Printf("booking confirmation %s, Message : %s\n", respBooking2.GetReceiptId(), respBooking2.GetDetail())
	// Remove User from Train
	removeResponse2, eru1 := client.RemoveUser(ctx, &pb.RemoveRequest{Email: "amit1@abc.com"})
	if eru1 != nil {
		log.Fatalf("could not remove user : %v", eru1)
	}
	log.Printf("booking removed for user email : amit1@abc.com, Success : %t\n", removeResponse2.Success)
}
