package entity

import (
	"express_be/core/record"
	"time"
)

type CODReconciliation struct {
	record.BaseEntity
	TransactionID      *string    `gorm:"foreignKey:TransactionID" json:"transaction_id"`
	Amount             *float64   `json:"amount"`
	Status             *string    `json:"status"`
	ReconciliationDate *time.Time `json:"reconciliation"`
}

func (c *CODReconciliation) TableName() string {
	return "cod_reconciliations"
}
