package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	"express_be/repository/delivery/entity"

	"github.com/google/uuid"
)

func RegisterToDeliveryPerson(req model.RegisterRequest) *entity.DeliveryPerson {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil
	}
	return &entity.DeliveryPerson{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Phone:                &req.Phone,
		PasswordHash:         &hashedPassword,
		CurrentAddress:       &req.CurrentAddress,
		Status:               entity.Pending,
		IdentificationNumber: &req.IdentificationNumber,
		FullName:             &req.FullName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
	}
}

func UpdateToDeliveryPerson(req model.UpdateDeliveryPersonReq) *entity.DeliveryPerson {
	return &entity.DeliveryPerson{
		CurrentAddress:       &req.CurrentAddress,
		IdentificationNumber: &req.IdentificationNumber,
		FullName:             &req.FullName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
	}
}
