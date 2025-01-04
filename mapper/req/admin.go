package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	entity "express_be/repository/admin"

	"github.com/google/uuid"
)

func ReqToAdmin(admin model.BaseRegisterRequest) *entity.Admin {
	hashedPassword, err := security.HashPassword(admin.Password)
	if err != nil {
		return nil
	}
	return &entity.Admin{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Phone:    &admin.Phone,
		Password: &hashedPassword,
	}
}
