package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	model "express_be/model/req"
	"express_be/repository/customer/entity"
	user "express_be/repository/user/entity"

	"github.com/google/uuid"
	"github.com/mmcloughlin/geohash"
)

func RegisterToCustomer(req model.RegisterRequest) (*entity.Customer, *user.User) {
	// Xử lý geohash luôn nếu cần
	geohash := geohash.Encode(req.Latitude, req.Longtitude)
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, nil
	}

	var status user.Status
	var approvalStatus user.ApprovalStatus
	if req.CustomerAccountType == string(entity.Postpaid) {
		status = user.Verified
		approvalStatus = user.Accepted
	} else {
		status = user.Pending
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
		Longtitude:  &req.Longtitude,
	}

	user := &user.User{
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
		Role:                 user.Customer,
	}
	return customer, user
}

func UpdateToCustomer(req model.UpdateUserReq) (*entity.Customer, *user.User) {
	geohash := geohash.Encode(req.Latitude, req.Longtitude)
	user := &user.User{
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
		GeoHash:    &geohash,
		Latitude:   &req.Latitude,
		Longtitude: &req.Longtitude,
	}
	return customer, user
}
