package server

import (
	"context"
	"log"
	"net"

	pb "github.com/shashyabh/mygo/proto"
	ser "github.com/shashyabh/mygo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedUserGrpcServiceServer
	service ser.UserService
}

func ListenGRPC (s ser.UserService, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("Error listening in grpc server")
		return err
	} 

	serv := grpc.NewServer()

	reflection.Register(serv)

	pb.RegisterUserGrpcServiceServer(serv, &grpcServer{
		UnimplementedUserGrpcServiceServer: pb.UnimplementedUserGrpcServiceServer{},
		service: s,
	} )
	return serv.Serve(lis)
}

func (s *grpcServer) CreateUser(ctx context.Context, r *pb.CreateUserRequest)(*pb.CreateUserResponse,error){
	user, err := s.service.CreateUser(
		ctx, r.Name, r.Salary, r.Department,"")
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User : &pb.User{
			Id: user.ID,
			Name: user.Name,
			Salary: user.Salary,
			Department: user.Department,
			Address: nil, 
		},
	},nil
}

func (s *grpcServer) GetUser(ctx context.Context, r *pb.GetUserRequest)(*pb.GetUserResponse,error){
	user, err := s.service.GetUser(
		ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User : &pb.User{
			Id: user.ID,
			Name: user.Name,
			Salary: user.Salary,
			Department: user.Department,
			Address: nil, 
		},
	},nil
}

func (s *grpcServer) GetAllUser(ctx context.Context, r *pb.GetAllUsersRequest)(*pb.GetAllUsersResponse,error){
	res, err := s.service.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	users := []*pb.User{}

	for _, u := range res{
		users = append(users,&pb.User{
				Id: u.ID,
				Name: u.Name,
				Salary: u.Salary,
				Department: u.Department,
				Address: nil,
			})
	}

	return &pb.GetAllUsersResponse{
		User: users,
	},nil
}