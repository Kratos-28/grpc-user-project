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

func StartServer() (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return nil, nil, err
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServiceServer{})
	log.Println("gRPC server is running on port :50051")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return s, lis, nil
}

func main() {
	_, _, err := StartServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
