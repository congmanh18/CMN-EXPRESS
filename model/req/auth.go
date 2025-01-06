package req

import "github.com/go-playground/validator/v10"

var Validate = validator.New(validator.WithRequiredStructEnabled())

type LoginRequest struct {
	Phone    *string `json:"phone" validate:"required"`
	Password *string `json:"password" validate:"required"`
	Role     *string `json:"role" validate:"required"`
}

type RegisterRequest struct {
	// 1. Default field
	Phone                string `json:"phone" validate:"required"`
	Password             string `json:"password" validate:"required"`
	CurrentAddress       string `json:"current_address,omitempty"`
	IdentificationNumber string `json:"identification_number,omitempty"`
	FullName             string `json:"full_name,omitempty"`
	DateOfBirth          string `json:"date_of_birth,omitempty"`
	Gender               string `json:"gender,omitempty"`
	Nationality          string `json:"nationality,omitempty"`
	PlaceOfOrigin        string `json:"place_of_origin"`
	PlaceOfResidence     string `json:"place_of_residence"`

	// 2. Field for customer
	CustomerAccountType string  `json:"account_type"`
	Latitude            float64 `json:"latitude,omitempty"`
	Longtitude          float64 `json:"longtitude,omitempty"`

	// 3. Field for delivery person
	Role string `json:"role,omitempty"`
}

type ForgotPasswordRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type ResetPasswordRequest struct {
	Role            *string `json:"role" validate:"required"`
	Phone           *string `json:"phone" validate:"required"`
	NewPassword     *string `json:"new_password" validate:"required"`
	ConfirmPassword *string `json:"confirm_password" validate:"required"`
}

// type LoginResponse struct{
// }

type UIDRequest struct {
	ID int `json:"id" validate:"required"`
}

type BaseRegisterRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
