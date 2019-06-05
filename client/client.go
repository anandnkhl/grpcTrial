package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	communicationpb "grpcTrial/communicationpb/proto"
	"log"
)

func Start(port string) {
	cc, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	ctx := context.Background()
	client := communicationpb.NewReplyServiceClient(cc)

	fmt.Println("Unary Request")
	var n string
	fmt.Print("Enter Name: ")
	fmt.Scan(&n)

	resGreet, err := client.Greet(ctx, &communicationpb.CommunicationRequest{
		Msg: n,
	})

	finalMessage := resGreet
	fmt.Println(finalMessage)
}
