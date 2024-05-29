package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpcProject/grpc-userservice/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server using grpc.DialContext
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	// Test GetUser
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", r.User)

	// Test GetUsers
	ids := []int32{1, 2}
	r2, err := c.GetUsers(ctx, &pb.GetUsersRequest{Ids: ids})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("Users: %v", r2.Users)

	// Test SearchUsers
	searchReq := &pb.SearchUsersRequest{
		Fname:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Married: true,
	}
	r3, err := c.SearchUsers(ctx, searchReq)
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	log.Printf("Search Result: %v", r3.Users)
}
