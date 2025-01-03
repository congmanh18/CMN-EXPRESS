package entity

import "express_be/core/record"

type OrderStatus struct {
	record.BaseEntity
	OrderID    *string `json:"order_id" gorm:"foreignKey:OrderID"` // Chỉ rõ khóa ngoại ở bảng con
	Status     *string `json:"status"`
	Reason     *string `json:"reason"`
	PhotoURL   *string `json:"photo_url"`
	FailedDate *string `json:"failed_date"`
}

func (o *OrderStatus) TableName() string {
	return "order_status"
}
