package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	accountingEntity "express_be/repository/accounting/entity"
	adminEntity "express_be/repository/admin/entity"
	userEntity "express_be/repository/user/entity"

	"github.com/google/uuid"
)

func RegisterToAdmin(req model.RegisterRequest) (*adminEntity.Admin, *userEntity.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())

	user := &userEntity.User{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Password: &hashedPassword,
		Role:     userEntity.Admin,
	}
	admin := &adminEntity.Admin{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone: &req.Phone,
	}
	return admin, user
}

func RegisterToAccounting(req model.RegisterRequest) (*accountingEntity.Accounting, *userEntity.User) {
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}
	id := pointer.String(uuid.New().String())
	user := &userEntity.User{
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
		Role:                 userEntity.Accounting,
	}
	accounting := &accountingEntity.Accounting{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		Phone: &req.Phone,
	}

	return accounting, user
}
