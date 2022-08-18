package main

import (
	"context"
	"fmt"
	pb "github.com/DanilaVyatkin/gRPC_test.gi/gen/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewTestApiClient(conn)
	response, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(response, response.Msg)

}
