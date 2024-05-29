package main

import (
	"log"
	"net"

	pb "github.com/grpcProject/grpc-userservice/pb"

	"google.golang.org/grpc"
)

// UserServiceServer implements the user.UserServiceServer interface
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServiceServer{})
	log.Println("gRPC server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
