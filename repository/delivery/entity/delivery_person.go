package entity

import (
	"express_be/core/record"
	indentity "express_be/repository/identity/entity"
	warehouse "express_be/repository/warehouse/entity"
)

type Status string

const (
	Pending   Status = "pending"
	Verified  Status = "verified"
	Blocked   Status = "blocked"
	Active    Status = "active"
	Inactive  Status = "inactive"
	Suspended Status = "suspended"
)

type DeliveryPerson struct {
	record.BaseEntity
	Phone          *string `gorm:"unique"`
	PasswordHash   *string
	CurrentAddress *string

	SalaryRate      *float64
	ExtraWeightRate *float64
	TotalCod        *float64
	DailyIncome     *float64
	GeoHash         *float64
	Status          Status

	BankInfos []indentity.BankInfo             `gorm:"foreignKey:DeliveryPersonID"`
	Identity  []indentity.UserIdentity         `gorm:"foreignKey:DeliveryPersonID"`
	Salary    []Salary                         `gorm:"foreignKey:DeliveryPersonID"`
	Warehouse []warehouse.WarehouseTransaction `gorm:"foreignKey:DeliveryPersonID"`
}

func (d *DeliveryPerson) TableName() string {
	return "delivery_persons"
}
