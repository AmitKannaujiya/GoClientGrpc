package client

import (
	"context"
	pb "grpc-client/proto/api.v1"
	u "grpc-client/utill"

	"log"

	"google.golang.org/grpc"
)

type RequestData struct {
	name string
	message string
	id      int
}


func SetupClient() {
	connection, err := grpc.Dial(u.BASE_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect :  %v", err)
	}
	defer connection.Close()
	client := pb.NewComServiceClient(connection)
	resp, er := client.ProcessMessage(context.Background(),& pb.Request{Name : "Cleint:09", Message : "Hello from 09", Id : 123})

	if er != nil {
		log.Fatalf("error sending request : %v", er)
	}
	log.Printf("Response from server : %v", resp)
}
