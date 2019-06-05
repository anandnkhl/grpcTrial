package main

import (
	"fmt"
	"grpcTrial/client"
	"grpcTrial/server"
)

func main() {
	go func() {
		server.Start("50051")
		fmt.Println("Server is Running")
	}()

	client.Start("50051")
}