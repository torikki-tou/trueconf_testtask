package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/torikki-tou/trueconf_testtask/common/custom_error"
	"github.com/torikki-tou/trueconf_testtask/common/response"
	"github.com/torikki-tou/trueconf_testtask/dto"
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

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := dto.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	user_id, err := h.userService.CreateUser(request)
	if err != nil {
		render.Render(w, r, response.ErrIternal(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{"user_id": user_id})
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()
	if err != nil {
		render.Render(w, r, response.ErrIternal(err))
		return
	}

	render.JSON(w, r, users)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, &custom_error.ErrObjectNotFound{}){
			render.Render(w, r, response.ErrNotFound(err))
		} else {
			render.Render(w, r, response.ErrIternal(err))
		}
		return
	}

	render.JSON(w, r, user)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	request := dto.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	err := h.userService.UpdateUser(userID, request)
	if err != nil {
		if errors.Is(err, &custom_error.ErrObjectNotFound{}){
			render.Render(w, r, response.ErrNotFound(err))
		} else {
			render.Render(w, r, response.ErrIternal(err))
		}
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	err := h.userService.DeleteUser(userID)
	if err != nil {
		if errors.Is(err, &custom_error.ErrObjectNotFound{}){
			render.Render(w, r, response.ErrNotFound(err))
		} else {
			render.Render(w, r, response.ErrIternal(err))
		}
		return
	}

	render.Status(r, http.StatusNoContent)
}