package admin

import "express_be/core/record"

type Admin struct {
	record.BaseEntity
	Phone    *string `json:"phone" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

func (a *Admin) TableName() string {
	return "admins"
}
