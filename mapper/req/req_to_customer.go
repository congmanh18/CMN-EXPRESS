package req

import (
	"express_be/auth"
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/core/security"
	"express_be/repository/customer/entity"

	"github.com/google/uuid"
)

func RegisterToCustomer(req auth.RegisterRequest) *entity.Customer {
	var status entity.Status
	// Xử lý geohash luôn nếu cần
	// geohash := geohash.Encode(req.Latitude, req.Longtitude)
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil
	}
	if req.CustomerAccountType == string(entity.Prepaid) {
		status = entity.Pending
	} else if req.CustomerAccountType == string(entity.Postpaid) {
		status = entity.Verified
	}

	return &entity.Customer{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		AccountType:    entity.CustomerAccountType(req.CustomerAccountType),
		Phone:          &req.Phone,
		PasswordHash:   &hashedPassword,
		CurrentAddress: &req.CurrentAddress,
		GeoHash:        nil,
		Latitude:       &req.Latitude,
		Longtitude:     &req.Longtitude,
		Status:         status,
	}
}
