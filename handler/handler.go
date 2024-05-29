package handler

import (
	"context"

	pb "github.com/grpcProject/grpc-userservice/pb"

	"github.com/grpcProject/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range models.Users {
		if user.ID == req.Id {
			return &pb.GetUserResponse{User: &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var userList []*pb.User
	for _, id := range req.Ids {
		for _, user := range models.Users {
			if user.ID == id {
				userList = append(userList, &pb.User{
					Id:      user.ID,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	return &pb.GetUsersResponse{Users: userList}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var userList []*pb.User
	for _, user := range models.Users {
		if (req.Fname == "" || user.Fname == req.Fname) &&
			(req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Height == 0 || user.Height == req.Height) &&
			(!req.Married || user.Married == req.Married) {
			userList = append(userList, &pb.User{
				Id:      user.ID,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}
	return &pb.SearchUsersResponse{Users: userList}, nil
}
