package dto

import (
	"net/http"

	"github.com/torikki-tou/trueconf_testtask/config"
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
}

func (m *CreateUserRequest) Bind(r *http.Request) error {
	return config.Validator.Struct(m)
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (m *UpdateUserRequest) Bind(r *http.Request) error { return nil }