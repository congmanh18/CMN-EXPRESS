package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	"express_be/repository/customer/entity"

	"github.com/google/uuid"
	"github.com/mmcloughlin/geohash"
)

func ReqToCustomer(req model.RegisterRequest) *entity.Customer {
	// Xử lý geohash luôn nếu cần
	geohash := geohash.Encode(req.Latitude, req.Longtitude)
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil
	}

	return &entity.Customer{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		AccountType:          entity.CustomerAccountType(req.CustomerAccountType),
		Phone:                &req.Phone,
		PasswordHash:         &hashedPassword,
		CurrentAddress:       &req.CurrentAddress,
		GeoHash:              &geohash,
		Latitude:             &req.Latitude,
		Longtitude:           &req.Longtitude,
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

func UpdateToCustomer(req model.UpdateCustomerReq) *entity.Customer {
	geohash := geohash.Encode(req.Latitude, req.Longtitude)
	return &entity.Customer{
		AccountType:          entity.CustomerAccountType(req.CustomerAccountType),
		CurrentAddress:       &req.CurrentAddress,
		IdentificationNumber: &req.IdentificationNumber,
		FullName:             &req.FullName,
		DateOfBirth:          &req.DateOfBirth,
		Gender:               &req.Gender,
		Nationality:          &req.Nationality,
		PlaceOfOrigin:        &req.PlaceOfOrigin,
		PlaceOfResidence:     &req.PlaceOfResidence,
		GeoHash:              &geohash,
		Latitude:             &req.Latitude,
		Longtitude:           &req.Longtitude,
	}
}
