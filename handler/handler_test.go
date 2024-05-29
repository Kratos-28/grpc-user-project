package handler

import (
	"context"
	"testing"

	pb "github.com/grpcProject/grpc-userservice/pb"
	"github.com/grpcProject/models"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Mock data
var mockUsers = []models.User{
	{ID: 1, Fname: "John", City: "New York", Phone: 1234567890, Height: 180, Married: false},
	{ID: 2, Fname: "Jane", City: "Los Angeles", Phone: 2345678901, Height: 170, Married: true},
	{ID: 3, Fname: "Doe", City: "Chicago", Phone: 3456789012, Height: 160, Married: false},
}

func setup() {
	// Assign mock data to the global Users slice
	models.Users = mockUsers
}

func TestGetUser(t *testing.T) {
	setup()
	server := &UserServiceServer{}

	tests := []struct {
		request  *pb.GetUserRequest
		expected *pb.GetUserResponse
		errCode  codes.Code
	}{
		{&pb.GetUserRequest{Id: 1}, &pb.GetUserResponse{User: &pb.User{Id: 1, Fname: "John", City: "New York", Phone: 1234567890, Height: 180, Married: false}}, codes.OK},
		{&pb.GetUserRequest{Id: 4}, nil, codes.NotFound},
	}

	for _, test := range tests {
		resp, err := server.GetUser(context.Background(), test.request)
		if err != nil {
			st, _ := status.FromError(err)
			assert.Equal(t, test.errCode, st.Code())
		} else {
			assert.Nil(t, err)
			assert.Equal(t, test.expected, resp)
		}
	}
}

func TestGetUsers(t *testing.T) {
	setup()
	server := &UserServiceServer{}

	tests := []struct {
		request  *pb.GetUsersRequest
		expected *pb.GetUsersResponse
	}{
		{&pb.GetUsersRequest{Ids: []int32{1, 2}}, &pb.GetUsersResponse{Users: []*pb.User{
			{Id: 1, Fname: "John", City: "New York", Phone: 1234567890, Height: 180, Married: false},
			{Id: 2, Fname: "Jane", City: "Los Angeles", Phone: 2345678901, Height: 170, Married: true},
		}}},
		{&pb.GetUsersRequest{Ids: []int32{3, 4}}, &pb.GetUsersResponse{Users: []*pb.User{
			{Id: 3, Fname: "Doe", City: "Chicago", Phone: 3456789012, Height: 160, Married: false},
		}}},
	}

	for _, test := range tests {
		resp, err := server.GetUsers(context.Background(), test.request)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, resp)
	}
}

func TestSearchUsers(t *testing.T) {
	setup()
	server := &UserServiceServer{}

	tests := []struct {
		request  *pb.SearchUsersRequest
		expected *pb.SearchUsersResponse
	}{
		{&pb.SearchUsersRequest{Fname: "John"}, &pb.SearchUsersResponse{Users: []*pb.User{
			{Id: 1, Fname: "John", City: "New York", Phone: 1234567890, Height: 180, Married: false},
		}}},
		{&pb.SearchUsersRequest{City: "Chicago"}, &pb.SearchUsersResponse{Users: []*pb.User{
			{Id: 3, Fname: "Doe", City: "Chicago", Phone: 3456789012, Height: 160, Married: false},
		}}},
		{&pb.SearchUsersRequest{Married: true}, &pb.SearchUsersResponse{Users: []*pb.User{
			{Id: 2, Fname: "Jane", City: "Los Angeles", Phone: 2345678901, Height: 170, Married: true},
		}}},
		{&pb.SearchUsersRequest{Height: 160}, &pb.SearchUsersResponse{Users: []*pb.User{
			{Id: 3, Fname: "Doe", City: "Chicago", Phone: 3456789012, Height: 160, Married: false},
		}}},
	}

	for _, test := range tests {
		resp, err := server.SearchUsers(context.Background(), test.request)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, resp)
	}
}
