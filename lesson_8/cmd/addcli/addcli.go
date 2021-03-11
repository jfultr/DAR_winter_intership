package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/jfultr/DAR_winter_intership/lesson_8/pb"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// CreateUser
	request1 := &pb.CreateUserRequest{
		Name: "Farkhat",
	}
	response1, err := client.CreateUser(context.Background(), request1)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response1.Ok)

	// GetUser
	request2 := &pb.GetUserRequest{
		ID: "55c0b697-329a-40d6-b11f-9bb79c28eff5",
	}
	response2, err := client.GetUser(context.Background(), request2)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response2.Name)

}
