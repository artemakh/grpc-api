package client

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net/http"
	pb "uid/proto/proto/service"
)

func GetUsersId(w http.ResponseWriter, r *http.Request) {
	user1 := r.URL.Query().Get("user1")
	if user1 == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'user1' not specified in url.", 400)
		return
	}
	user2 := r.URL.Query().Get("user2")
	if user2 == "" {
		//Get not set, send a 400 bad request
		http.Error(w, "Get 'user2' not specified in url.", 400)
		return
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:38925", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewUidGenerateClient(conn)
	uidRequest := &pb.UidRequest{
		User1: user1,
		User2: user2,
	}
	response, err := client.UidGenerate(context.Background(), uidRequest)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	req := map[string]string{
		"uid": response.Message,
	}

	request, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("json marshal err: %v", err)
	}

	w.Write(request)
}
