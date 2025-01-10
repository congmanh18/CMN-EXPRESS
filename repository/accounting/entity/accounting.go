package entity

import "express_be/core/record"

type Accounting struct {
	record.BaseEntity
	Phone *string `gorm:"not null;index"`
}

func (a *Accounting) TableName() string {
	return "accountings"
}
