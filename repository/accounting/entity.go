package accounting

import "express_be/core/record"

type Accounting struct {
	record.BaseEntity
	Phone *string `gorm:"foreignKey:AccountingPhone"`
}

func (a *Accounting) TableName() string {
	return "accountings"
}
