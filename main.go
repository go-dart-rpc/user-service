package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-dart-rpc/user-model/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UserServiceServer
}

func (*server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	fmt.Println("Create user request")
	res := &proto.CreateUserResponse{
		User: &proto.User{
			Id:        "1",
			FirstName: "Bob",
			LastName:  "Marley",
			Email:     "bob.marley@gmail.com",
			Gender:    "M",
		},
	}
	return res, nil
}

func (*server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	fmt.Println("Delete user request")
	res := &proto.DeleteUserResponse{
		UserId: "1",
	}
	return res, nil
}

func (*server) ListUser(_ *proto.ListUserRequest, stream proto.UserService_ListUserServer) error {
	fmt.Println("List user request")
	res := &proto.ListUserResponse{
		User: &proto.User{
			Id:        "1",
			FirstName: "Bob",
			LastName:  "Marley",
			Email:     "bob.marley@gmail.com",
			Gender:    "M",
		},
	}
	return stream.Send(res)
}

func (*server) ReadUser(ctx context.Context, req *proto.ReadUserRequest) (*proto.ReadUserResponse, error) {
	fmt.Println("Read user request")
	res := &proto.ReadUserResponse{
		User: &proto.User{
			Id:        "1",
			FirstName: "Bob",
			LastName:  "Marley",
			Email:     "bob.marley@gmail.com",
			Gender:    "M",
		},
	}
	return res, nil
}

func (*server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	fmt.Println("Update user request")
	res := &proto.UpdateUserResponse{
		User: &proto.User{
			Id:        "1",
			FirstName: "Bob",
			LastName:  "Marley",
			Email:     "bob.marley@gmail.com",
			Gender:    "M",
		},
	}
	return res, nil
}

func main() {
	fmt.Println("User service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
