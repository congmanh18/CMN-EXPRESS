package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/entity"
	model "express_be/model/req"

	"github.com/google/uuid"
)

func RegisterToAdmin(req model.RegisterRequest) (*entity.Admin, *entity.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())

	user := &entity.User{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone:    &req.Phone,
		Password: &hashedPassword,
		Role:     entity.AdminRole,
	}
	admin := &entity.Admin{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone: &req.Phone,
	}
	return admin, user
}

func RegisterToAccounting(req model.RegisterRequest) (*entity.Accounting, *entity.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())
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
		IdentificationNumber: &req.IdentificationNumber,
		FirstName:            &req.FirstName,
		LastName:             &req.LastName,
		MiddleName:           &req.MiddleName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
		Role:                 entity.AccountingRole,
	}
	accounting := &entity.Accounting{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone: &req.Phone,
	}

	return accounting, user
}
