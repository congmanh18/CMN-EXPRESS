package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"

	"express_be/entity"

	"github.com/google/uuid"
	"github.com/mmcloughlin/geohash"
)

func RegisterToCustomer(req model.RegisterRequest) (*entity.Customer, *entity.User) {
	// Xử lý geohash luôn nếu cần
	geohash := geohash.Encode(req.Latitude, req.Longitude)
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}

	var status entity.AccountStatus
	var approvalStatus entity.ApprovalStatus
	if req.CustomerAccountType == string(entity.Postpaid) {
		status = entity.Verified
		approvalStatus = entity.Accepted
	} else {
		status = entity.Pending
		approvalStatus = ""
	}

	id := pointer.String(uuid.New().String())
	customer := &entity.Customer{
		BaseEntity: record.BaseEntity{
			ID: id,
		},
		AccountType: entity.CustomerAccountType(req.CustomerAccountType),
		Phone:       &req.Phone,
		GeoHash:     &geohash,
		Latitude:    &req.Latitude,
		Longitude:   &req.Longitude,
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
		Status:               status,
		ApprovalStatus:       approvalStatus,
		Role:                 entity.CustomerRole,
	}
	return customer, user
}

func UpdateToCustomer(req model.UpdateUserReq) (*entity.Customer, *entity.User) {
	geohash := geohash.Encode(req.Latitude, req.Longitude)
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
	customer := &entity.Customer{
		GeoHash:   &geohash,
		Latitude:  &req.Latitude,
		Longitude: &req.Longitude,
	}
	return customer, user
}
