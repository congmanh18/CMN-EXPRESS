package entity

import "express_be/core/record"

type OrderJourney struct {
	record.BaseEntity
	OrderID     *string `json:"order_id"`
	Status      Status  `json:"status"`
	UpdatedBy   *string `json:"created_by"`
	Description *string `json:"description"`
}
