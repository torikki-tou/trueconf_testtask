package service

import (
	"strconv"
	"time"

	"github.com/torikki-tou/trueconf_testtask/dto"
	"github.com/torikki-tou/trueconf_testtask/entity"
	"github.com/torikki-tou/trueconf_testtask/repo"
)

type UserService interface {
	CreateUser(request dto.CreateUserRequest) (int, error)
	GetUsers() (entity.UserList, error)
	GetUserByID(userID string) (entity.User, error)
	UpdateUser(userID string, request dto.UpdateUserRequest) error
	DeleteUser(userID string) error
}

func NewUserService(userRepo repo.UserRepository, queueRepo repo.RabbitRepository) UserService {
	return &userService{
		userRepo: userRepo,
		queueRepo: queueRepo,
	}
}

type userService struct {
	userRepo repo.UserRepository
	queueRepo repo.RabbitRepository
}

func (s *userService) CreateUser(request dto.CreateUserRequest) (int, error) {
	user := entity.User{
		CreatedAt: time.Now(),
		DisplayName: request.DisplayName,
		Email: request.Email,
	}

	userID, err := s.userRepo.InsertUser(user)
	if err != nil {
		return 0, err
	}

	err = s.queueRepo.ProduceMessage(dto.Notification{
		Type: "user_created",
		UserID: userID,
	})
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *userService) GetUsers() (entity.UserList, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return entity.UserList{}, err
	}
	return users, nil
}

func (s *userService) GetUserByID(userID string) (entity.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *userService) UpdateUser(userID string, request dto.UpdateUserRequest) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.DisplayName = request.DisplayName

	err = s.userRepo.UpdateUser(userID, user)
	if err != nil {
		return err
	}

	intID, _ := strconv.Atoi(userID)
	err = s.queueRepo.ProduceMessage(dto.Notification{
		Type: "user_updated",
		UserID: intID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(userID string) error {
	err := s.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	intID, _ := strconv.Atoi(userID)
	err = s.queueRepo.ProduceMessage(dto.Notification{
		Type: "user_deleted",
		UserID: intID,
	})
	if err != nil {
		return err
	}

	return nil
}