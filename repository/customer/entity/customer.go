package entity

import (
	"express_be/core/record"
	indentity "express_be/repository/identity/entity"
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
	AccountType    CustomerAccountType
	Phone          *string `gorm:"uniqueIndex:idx_phone_deleted_at"`
	PasswordHash   *string
	CurrentAddress *string

	GeoHash    *string  `json:"geo_hash"`
	Status     Status   `json:"status"`
	Latitude   *float64 `json:"latitude"`
	Longtitude *float64 `json:"longtitude"`

	IdentificationNumber *string `json:"identification_number"`
	FullName             *string `json:"full_name"`
	DateOfBirth          *string `json:"date_of_birth"`
	Gender               *string `json:"gender"`
	Nationality          *string `json:"nationality"`
	PlaceOfOrigin        *string `json:"place_of_origin"`
	PlaceOfResidence     *string `json:"place_of_residence"`

	BankInfos []indentity.BankInfo `gorm:"foreignKey:CustomerID"`
	Order     []order.Order        `gorm:"foreignKey:SenderID" json:"order"`
}

func (c *Customer) TableName() string {
	return "customers"
}
