package main

import (
	"context"
	bp "github.com/DanilaVyatkin/gRPC_test.gi/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type testApiServer struct {
	bp.UnimplementedTestApiServer
}

func (t *testApiServer) GetUser(ctx context.Context, req *bp.UserRequest) (*bp.UserResponse, error) {
	return &bp.UserResponse{}, nil
}

func (t *testApiServer) Echo(ctx context.Context, req *bp.ResponseRequest) (*bp.ResponseRequest, error) {
	return req, nil
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	bp.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}
}
