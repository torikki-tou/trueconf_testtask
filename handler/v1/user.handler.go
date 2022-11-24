package v1

import (
	"net/http"

	"github.com/torikki-tou/trueconf_testtask/service"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

type userHandler struct {
	userService service.UserService
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}
func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {}
func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {}
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
