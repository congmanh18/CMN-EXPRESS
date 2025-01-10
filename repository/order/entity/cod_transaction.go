package entity

import "express_be/core/record"

type CODTransaction struct {
	record.BaseEntity
	OrderID          *string  `gorm:"foreignKey:OrderID" json:"order_id"`
	DeliveryPersonID *string  `json:"delivery_person_id"`
	CustomerID       *string  `json:"customer_id"`
	Amount           *float64 `json:"amount"`
	TransactionDate  *string  `json:"transactionDate"`
	Status           *string  `json:"status"`

	CODReconciliation []CODReconciliation `gorm:"foreignKey:TransactionID" json:"cod_reconciliation"`
}

func (c *CODTransaction) TableName() string {
	return "cod_transactions"
}
