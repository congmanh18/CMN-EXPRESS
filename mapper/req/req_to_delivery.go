package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/entity"
	model "express_be/model/req"

	"github.com/google/uuid"
)

func RegisterToDeliveryPerson(req model.RegisterRequest) (*entity.DeliveryPerson, *entity.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())

	deliveryPerson := &entity.DeliveryPerson{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		FullName: pointer.String(req.LastName + " " + req.MiddleName + " " + req.FirstName),
		Phone:    &req.Phone,
	}
	user := &entity.User{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone:                &req.Phone,
		Password:             &hashedPassword,
		SpecificAddress:      &req.SpecificAddress,
		Ward:                 &req.Ward,
		District:             &req.District,
		City:                 &req.City,
		IDCard:               &req.IDCard,
		IdentificationNumber: &req.IdentificationNumber,
		LastName:             &req.LastName,
		MiddleName:           &req.MiddleName,
		FirstName:            &req.FirstName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
		Status:               entity.Pending,
		Role:                 entity.DeliveryPersonRole,
	}
	return deliveryPerson, user
}

func UpdateToDeliveryPerson(req model.UpdateUserReq) (*entity.DeliveryPerson, *entity.User) {
	user := &entity.User{
		SpecificAddress:      &req.SpecificAddress,
		Ward:                 &req.Ward,
		District:             &req.District,
		City:                 &req.City,
		IDCard:               &req.IDCard,
		IdentificationNumber: &req.IdentificationNumber,
		LastName:             &req.LastName,
		MiddleName:           &req.MiddleName,
		FirstName:            &req.FirstName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
	}
	deliveryPerson := &entity.DeliveryPerson{}
	return deliveryPerson, user
}
