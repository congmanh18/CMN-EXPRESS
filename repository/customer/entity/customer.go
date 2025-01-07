package entity

import (
	"express_be/core/record"
	indentity "express_be/repository/identity/entity"
	order "express_be/repository/order/entity"
)

type CustomerAccountType string

const (
	Prepaid  CustomerAccountType = "prepaid"
	Postpaid CustomerAccountType = "postpaid"
)

type Customer struct {
	record.BaseEntity
	AccountType CustomerAccountType
	Phone       *string  `gorm:"foreignKey:CustomerPhone"`
	GeoHash     *string  `json:"geo_hash"`
	Latitude    *float64 `json:"latitude"`
	Longtitude  *float64 `json:"longtitude"`

	BankInfos []indentity.BankInfo `gorm:"foreignKey:CustomerID"`
	Order     []order.Order        `gorm:"foreignKey:SenderID" json:"order"`
}

func (c *Customer) TableName() string {
	return "customers"
}
