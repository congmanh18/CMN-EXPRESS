package entity

type CustomerDetails struct {
	User
	AccountType *string  `json:"account_type" gorm:"column:account_type"`
	Latitude    *float64 `json:"latitude" gorm:"column:latitude"`
	Longtitude  *float64 `json:"longtitude" gorm:"column:longtitude"`
}

type DeliveryPersonDetails struct {
	User
	SalaryRate *float64 `json:"salary_rate" gorm:"column:salary_rate"`
}

type UserDetails struct {
	User
	AccountType *string `json:"account_type" gorm:"column:account_type"`
}
