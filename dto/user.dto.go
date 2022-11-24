package dto

import (
	"net/http"
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (m *CreateUserRequest) Bind(r *http.Request) error { return nil }

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (m *UpdateUserRequest) Bind(r *http.Request) error { return nil }