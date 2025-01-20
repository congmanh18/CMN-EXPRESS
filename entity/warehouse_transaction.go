package entity

import (
	"express_be/core/record"
)

type WarehouseTransaction struct {
	record.BaseEntity
	OrderID          *string `gorm:"foreignKey:OrderID" json:"order_id"`
	WarehouseID      *string `gorm:"foreignKey:WarehouseID" json:"warehouse_id"`
	DeliveryPersonID *string `json:"delivery_person_id"`
	TransactionType  *string `json:"transaction_type"`
	Remarks          *string `json:"remarks"`

	Order *Order `gorm:"foreignKey:WarehouseTransactionID" json:"order"`
}

func (w *WarehouseTransaction) TableName() string {
	return "warehouse_transactions"
}
