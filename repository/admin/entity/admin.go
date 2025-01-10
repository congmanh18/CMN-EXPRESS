package entity

import "express_be/core/record"

type Admin struct {
	record.BaseEntity
	Phone *string `gorm:"not null;index"`
}

func (a *Admin) TableName() string {
	return "admins"
}
