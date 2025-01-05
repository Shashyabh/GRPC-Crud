package client

import (
	"context"
	"fmt"
	"log"

	"github.com/shashyabh/mygo/models"
	pb "github.com/shashyabh/mygo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.UserGrpcServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to create client")
		return nil, err
	}

	ser := pb.NewUserGrpcServiceClient(conn)

	return &Client{conn, ser},nil
}

func (c *Client) CreateUser(ctx context.Context,id,name string, salary int64,department,address_id string)(*models.User,error){
	user, err := c.service.CreateUser(
		ctx,
		&pb.CreateUserRequest{
			Name: name,
			Salary: salary,
			Department: department,
			Address:nil ,
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Name: user.User.Name,
		Salary: user.User.Salary,
		Department: user.User.Department,
		AddressId: "", 
	},nil
}

func (c *Client) GetUser(ctx context.Context, id string) (*models.User,error){
	user, err := c.service.GetUser(
		ctx,
		&pb.GetUserRequest{Id: id},
	);
	if err != nil {
		return nil, fmt.Errorf("failed to get user from repository: %w", err);
	}
	return &models.User{
		Name: user.User.Name,
		Salary: user.User.Salary,
		Department: user.User.Department,
		AddressId: "", 
	},nil
}

func (c *Client) GetAllUser(ctx context.Context) ([] models.User, error){
	//var users []models.User
	user, err := c.service.GetAllUser(ctx,
		&pb.GetAllUsersRequest{},
	)
	if err != nil{
		return nil, fmt.Errorf("failed to get users from repository: %w", err);
	}

	users := []models.User{}

	for _,u := range user.User{
		users = append(users,models.User{
			Name: u.Name,
			Salary: u.Salary,
			Department: u.Department,
			AddressId: "", 
		})
	}
	return users, nil
}
