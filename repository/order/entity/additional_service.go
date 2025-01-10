package entity

import "express_be/core/record"

type AdditionalService struct {
	record.BaseEntity
	OrderID     *string `gorm:"foreignKey:OrderID" json:"order_id"`
	ServiceName *string `json:"service_name"`
	ServiceFee  *string `json:"service_fee"`
}
