package auth

type RegisterRequest struct {
	// 1. Default field
	Phone          string `json:"phone" validate:"required"`
	Password       string `json:"password" validate:"required"`
	CurrentAddress string `json:"current_address"`

	// 2. Field for customer
	CustomerAccountType string  `json:"account_type"`
	Latitude            float64 `json:"latitude"`
	Longtitude          float64 `json:"longtitude"`

	// 3. Field for delivery person
}
