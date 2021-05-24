package client_test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"uid/client"

	pb "uid/proto/proto/service"
)

type mockUidGenerate struct {
	pb.UnimplementedUidGenerateServer
}

func (*mockUidGenerate) UidGenerate(ctx context.Context, req *pb.UidRequest) (response *pb.UidResponse, err error) {
	if req.User1 == "" || req.User2 == "" {
		response = &pb.UidResponse{
			Message: "",
		}
		return response, nil
	}

	user1 := req.User1
	user2 := req.User2
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

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterUidGenerateServer(server, &mockUidGenerate{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestUidGenerate_Uid(t *testing.T) {
	tests := []struct {
		name    string
		User1  string
		User2  string
		res     *pb.UidResponse
		StatusCode  int
	}{
		{
			"valid request with empty stings",
			"",
			"",
			&pb.UidResponse{Message: ""},
			400,
		},
		{
			"valid request with two equal string",
			"abc",
			"abc",
			&pb.UidResponse{Message: "aabbcc"},
			200,
		},
		{
			"valid request with user1 string < user2 string",
			"aaaa",
			"bbb",
			&pb.UidResponse{Message: "abababa"},
			200,
		},
		{
			"valid request with user1 string > user2 string",
			"bbb",
			"aaaa",
			&pb.UidResponse{Message: "abababa"},
			200,
		},
		{
			"valid request",
			"5f5h9s83fs",
			"9s85w5qfs",
			&pb.UidResponse{Message: "59fs585h9w5s8q3ffss"},
			200,
		},
		{
			"valid request",
			"9s85w5qfs",
			"5f5h9s83fs",
			&pb.UidResponse{Message: "59fs585h9w5s8q3ffss"},
			200,
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/get_method", nil)
			if err != nil {
				t.Fatal(err)
			}
			user1 := req.URL.Query()
			user1.Add("user1", tt.User1)
			req.URL.RawQuery = user1.Encode()
			user2 := req.URL.Query()
			user2.Add("user2", tt.User2)
			req.URL.RawQuery = user2.Encode()
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(client.GetUsersId)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK && status != tt.StatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			fmt.Printf("a")
		})
	}
}