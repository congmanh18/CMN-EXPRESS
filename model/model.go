package model

type LoginRequest struct {
	Phone        *string `json:"phone" validate:"required"`
	PasswordHash *string `json:"password" validate:"required"`
}
