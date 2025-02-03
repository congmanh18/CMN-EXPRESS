package entity

import (
	"express_be/core/record"
)

type ActiveStatus string

const (
	OnDuty  ActiveStatus = "on_duty"
	OffDuty ActiveStatus = "off_duty"
)

type DeliveryPerson struct {
	record.BaseEntity
	Phone           *string `gorm:"not null;index"`
	SalaryRate      *float64
	ExtraWeightRate *float64
	TotalCod        *float64
	DailyIncome     *float64
	GeoHash         *string
	
	IsAvailable  bool
	ActiveStatus ActiveStatus

	Salary    []Salary               `gorm:"foreignKey:DeliveryPersonID"`
	Warehouse []WarehouseTransaction `gorm:"foreignKey:DeliveryPersonID"`
}

func (d *DeliveryPerson) TableName() string {
	return "delivery_persons"
}
