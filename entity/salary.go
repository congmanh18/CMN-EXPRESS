package entity

import "express_be/core/record"

type Salary struct {
	record.BaseEntity
	DeliveryPersonID  *string  `gorm:"foreignKey:DeliveryPersonID" json:"delivery_person_id"`
	Period            *string  `json:"period"`
	TotalOrders       *int     `json:"total_orders"`
	BaseRatePerOrder  *int     `json:"base_rate_per_order"`
	TotalBaseSalary   *float64 `json:"total_base_salary"`
	TotalWeight       *float64 `json:"total_weight"`
	ExtraWeightSalary *float64 `json:"extra_weight_salary"`
	TotalSalary       *float64 `json:"total_salary"`
	CalculatedDate    *float64 `json:"calculated_date"`
}

func (s *Salary) TableName() string {
	return "salaries"
}
