package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "uid/proto/proto/service"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterUidGenerateServer(server, &UidGenerateServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestUidServer_UidGenerate(t *testing.T) {
	tests := []struct {
		name    string
		User1  string
		User2  string
		res     *pb.UidResponse
		errMsg  string
	}{
		{
			"valid request with empty stings",
			"",
			"",
			&pb.UidResponse{Message: ""},
			"",
		},
		{
			"valid request with two equal string",
			"abc",
			"abc",
			&pb.UidResponse{Message: "aabbcc"},
			"",
		},
		{
			"valid request with user1 string < user2 string",
			"aaaa",
			"bbb",
			&pb.UidResponse{Message: "abababa"},
			"",
		},
		{
			"valid request with user1 string > user2 string",
			"bbb",
			"aaaa",
			&pb.UidResponse{Message: "abababa"},
			"",
		},
		{
			"valid request",
			"5f5h9s83fs",
			"9s85w5qfs",
			&pb.UidResponse{Message: "59fs585h9w5s8q3ffss"},
			"",
		},
		{
			"valid request",
			"9s85w5qfs",
			"5f5h9s83fs",
			&pb.UidResponse{Message: "59fs585h9w5s8q3ffss"},
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUidGenerateClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.UidRequest{User1: tt.User1, User2: tt.User2}

			response, err := client.UidGenerate(ctx, request)

			if err != nil {
				t.Fatalf("Test fail without response from server: err %v", err)
			}

			if response != nil {
				if response.Message != tt.res.Message {
					t.Error("response: expected", tt.res.Message, "received", response.Message)
				}
			}
		})
	}
}