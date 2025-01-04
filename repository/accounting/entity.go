package accounting

import "express_be/core/record"

type Accounting struct {
	record.BaseEntity
	Phone    *string `json:"phone" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

func (a *Accounting) TableName() string {
	return "accountings"
}
