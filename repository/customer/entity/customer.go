package entity

import (
	"express_be/core/record"
	order "express_be/repository/order/entity"
)

type CustomerAccountType string
type Status string

const (
	Prepaid  CustomerAccountType = "prepaid"
	Postpaid CustomerAccountType = "postpaid"

	Pending   Status = "pending"
	Verified  Status = "verified"
	Blocked   Status = "blocked"
	Active    Status = "active"
	Inactive  Status = "inactive"
	Suspended Status = "suspended"
)

type Customer struct {
	record.BaseEntity
	AccountType    CustomerAccountType `json:"account_type"`
	Phone          *string             `gorm:"unique"`
	PasswordHash   *string
	CurrentAddress *string

	GeoHash    *float64 `json:"geo_hash"`
	Status     Status   `json:"status"`
	Latitude   *float64 `json:"latitude"`
	Longtitude *float64 `json:"longtitude"`

	Order []order.Order `gorm:"foreignKey:SenderID" json:"order"`
}

func (c *Customer) TableName() string {
	return "customers"
}
