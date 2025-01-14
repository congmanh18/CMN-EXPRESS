package req

import "github.com/go-playground/validator/v10"

var Validate = validator.New(validator.WithRequiredStructEnabled())

type LoginRequest struct {
	Phone    *string `json:"phone" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	// 1. Default field
	Phone                string `json:"phone" validate:"required"`
	Password             string `json:"password" validate:"required"`
	SpecificAddress      string `json:"specific_address"`
	Ward                 string `json:"ward"`
	District             string `json:"district"`
	City                 string `json:"city"`
	IDCard               string `json:"id_card"`
	IdentificationNumber string `json:"identification_number,omitempty"`
	LastName             string `json:"last_name,omitempty"`
	MiddleName           string `json:"middle_name,omitempty"`
	FirstName            string `json:"first_name,omitempty"`
	DateOfBirth          string `json:"date_of_birth,omitempty"`
	Gender               string `json:"gender,omitempty"`
	Nationality          string `json:"nationality,omitempty"`
	PlaceOfOrigin        string `json:"place_of_origin"`
	PlaceOfResidence     string `json:"place_of_residence"`

	// 2. Field for customer
	CustomerAccountType string  `json:"account_type"`
	Latitude            float64 `json:"latitude,omitempty"`
	Longtitude          float64 `json:"longtitude,omitempty"`

	Role string `json:"role" validate:"required"`
}

type ForgotPasswordRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type ResetPasswordRequest struct {
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
