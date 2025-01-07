package entity

import (
	"express_be/core/record"
	indentity "express_be/repository/identity/entity"
	warehouse "express_be/repository/warehouse/entity"
)

type DeliveryPerson struct {
	record.BaseEntity
	Phone           *string `gorm:"foreignKey:DeliveryPersonPhone"`
	SalaryRate      *float64
	ExtraWeightRate *float64
	TotalCod        *float64
	DailyIncome     *float64
	GeoHash         *string

	BankInfos []indentity.BankInfo             `gorm:"foreignKey:DeliveryPersonID"`
	Salary    []Salary                         `gorm:"foreignKey:DeliveryPersonID"`
	Warehouse []warehouse.WarehouseTransaction `gorm:"foreignKey:DeliveryPersonID"`
}

func (d *DeliveryPerson) TableName() string {
	return "delivery_persons"
}
