package admin

import "express_be/core/record"

type Admin struct {
	record.BaseEntity
	Phone *string `gorm:"foreignKey:AdminPhone"`
}

func (a *Admin) TableName() string {
	return "admins"
}
