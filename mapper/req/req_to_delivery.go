package req

import (
	"express_be/auth"
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/repository/delivery/entity"

	"github.com/google/uuid"
)

func RegisterToDeliveryPerson(req auth.RegisterRequest) *entity.DeliveryPerson {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil
	}
	return &entity.DeliveryPerson{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Phone:          &req.Phone,
		PasswordHash:   &hashedPassword,
		CurrentAddress: &req.CurrentAddress,
		Status:         entity.Pending,
	}
}
