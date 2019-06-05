package server

import (
	"context"
	communicationpb "grpcTrial/communicationpb/proto"
)

type CommunicationHandler struct{}



func (ch * CommunicationHandler) Greet(ctx context.Context, request *communicationpb.CommunicationRequest) (*communicationpb.CommunicationResponse, error){
	response := &communicationpb.CommunicationResponse{}
	response.Result = "Hello "+request.GetMsg()+"! Welcome to Perennial Systems"
	return response, nil
}