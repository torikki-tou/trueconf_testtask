package service

import (
	"github.com/torikki-tou/trueconf_testtask/dto"
	"github.com/torikki-tou/trueconf_testtask/entity"
	"github.com/torikki-tou/trueconf_testtask/repo"
)

type UserService interface {
	CreateUser(request dto.CreateUserRequest) (entity.User, error)
	GetUsers() (entity.User, error)
	GetUserByID(userID string) (entity.User, error)
	UpdateUser(userID string, request dto.UpdateUserRequest) error
	DeleteUser(userID string) error
}

func NewUserService(userRepo repo.UserRepositiry) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo repo.UserRepositiry
}

func (s *userService) CreateUser(request dto.CreateUserRequest) (entity.User, error) {return entity.User{}, nil}
func (s *userService) GetUsers() (entity.User, error) {return entity.User{}, nil}
func (s *userService) GetUserByID(userID string) (entity.User, error) {return entity.User{}, nil}
func (s *userService) UpdateUser(userID string, request dto.UpdateUserRequest) error {return nil}
func (s *userService) DeleteUser(userID string) error {return nil}