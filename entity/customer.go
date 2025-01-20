package entity

import (
	"express_be/core/record"
)

type CustomerAccountType string

const (
	Prepaid  CustomerAccountType = "prepaid"
	Postpaid CustomerAccountType = "postpaid"
)

// Shop/Sender/Customer
type Customer struct {
	record.BaseEntity
	AccountType CustomerAccountType
	ShopName    *string
	Phone       *string  `gorm:"not null;index"`
	GeoHash     *string  `json:"geo_hash"`
	Latitude    *float64 `json:"latitude"`
	Longtitude  *float64 `json:"longtitude"`
	Order       []Order  `gorm:"foreignKey:SenderID" json:"order"`
}

func (c *Customer) TableName() string {
	return "customers"
}
