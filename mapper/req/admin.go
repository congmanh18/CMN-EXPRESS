package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	accountingEntity "express_be/repository/accounting"
	adminEntity "express_be/repository/admin"

	"github.com/google/uuid"
)

func ReqToAdmin(admin model.RegisterRequest) *adminEntity.Admin {
	hashedPassword, err := security.HashPassword(admin.Password)
	if err != nil {
		return nil
	}
	return &adminEntity.Admin{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Phone:    &admin.Phone,
		Password: &hashedPassword,
	}
}

func ReqToAccounting(accounting model.RegisterRequest) *accountingEntity.Accounting {
	hashedPassword, err := security.HashPassword(accounting.Password)
	if err != nil {
		return nil
	}
	return &accountingEntity.Accounting{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Phone:    &accounting.Phone,
		Password: &hashedPassword,
	}
}
