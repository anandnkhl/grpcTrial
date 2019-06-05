package server

import (
	"google.golang.org/grpc"
	communicationpb "grpcTrial/communicationpb/proto"
	"log"
	"net"
)

func Start (port string){
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil{
		log.Fatal(err)
	}

	s := grpc.NewServer()
	communicationpb.RegisterReplyServiceServer(s, &CommunicationHandler{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}