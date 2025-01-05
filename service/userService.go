package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shashyabh/mygo/models"
	repo "github.com/shashyabh/mygo/repository"
)

type UserService interface {
	CreateUser(ctx context.Context,name string, salary int64,department,address_id string) (*models.User, error)
	GetUser(ctx context.Context,id string) (*models.User,error)
	GetAllUser(ctx context.Context) ([]models.User, error)
}

type userService struct{
	repository repo.UserRepositoryInterface
}

func NewService(r repo.UserRepositoryInterface) UserService {
	return &userService{r}
}

func (s *userService) CreateUser(ctx context.Context,name string, salary int64,department,address_id string) (*models.User, error){
	var id = uuid.NewString();
	user := &models.User{
		ID : id,
		Name: name,
		Salary: salary,
		Department: department,
		AddressId: address_id,
	}

	err := s.repository.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*models.User,error){
	user, err := s.repository.GetUserById(id);
	if err != nil {
		return nil, fmt.Errorf("failed to get user from repository: %w", err);
	}
	return user, nil
}

func (s *userService) GetAllUser(ctx context.Context) ([] models.User, error){
	//var users []models.User
	users, err := s.repository.GetAllUsers();
	if err != nil{
		return nil, fmt.Errorf("failed to get users from repository: %w", err);
	}
	return users, nil
}