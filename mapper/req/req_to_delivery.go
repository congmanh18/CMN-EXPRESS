package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	"express_be/repository/delivery/entity"
	user "express_be/repository/user/entity"

	"github.com/google/uuid"
)

func RegisterToDeliveryPerson(req model.RegisterRequest) (*entity.DeliveryPerson, *user.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())

	deliveryPerson := &entity.DeliveryPerson{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone: &req.Phone,
	}
	user := &user.User{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone:                &req.Phone,
		Password:             &hashedPassword,
		CurrentAddress:       &req.CurrentAddress,
		IdentificationNumber: &req.IdentificationNumber,
		FullName:             &req.FullName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
		Status:               user.Pending,
		Role:                 user.DeliveryPerson,
	}
	return deliveryPerson, user
}

func UpdateToDeliveryPerson(req model.UpdateDeliveryPersonReq) (*entity.DeliveryPerson, *user.User) {
	user := &user.User{
		CurrentAddress:       &req.CurrentAddress,
		IdentificationNumber: &req.IdentificationNumber,
		FullName:             &req.FullName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
	}
	deliveryPerson := &entity.DeliveryPerson{}
	return deliveryPerson, user
}
