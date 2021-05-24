package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "uid/proto/proto/service"
)

type UidGenerateServer struct {
	pb.UnimplementedUidGenerateServer
}

func(s *UidGenerateServer) UidGenerate(ctx context.Context, msg *pb.UidRequest) (response *pb.UidResponse, err error) {
	if msg.User1 == "" || msg.User2 == "" {
		response = &pb.UidResponse{
			Message: "",
		}
		return response, nil
	}

	user1 := msg.User1
	user2 := msg.User2
	var respMsg string

	for i:= 0; i < len(user1) || i < len(user2); i++ {
		if i < len(user1) && i < len(user2) {
			if user1[i] <= user2[i] {
				respMsg += user1[i:i+1] + user2[i:i+1]
			} else if user1[i] > user2[i] {
				respMsg += user2[i:i+1] + user1[i:i+1]
			}
		} else if i < len(user1) && i >= len(user2) {
			respMsg += user1[i:i+1]
		} else if i < len(user2) && i >= len(user1) {
			respMsg += user2[i:i+1]
		}
	}

	response = &pb.UidResponse{
		Message: respMsg,
	}
	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":38925")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUidGenerateServer(grpcServer, &UidGenerateServer{})
	grpcServer.Serve(listener)
}