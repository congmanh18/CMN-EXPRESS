package entity

import "express_be/core/record"

type UserIdentity struct {
	record.BaseEntity
	CustomerID       *string `gorm:"foreignKey:CustomerID" json:"customer_id"`
	DeliveryPersonID *string `gorm:"foreignKey:DeliveryPersonID" json:"delivery_person_id"`

	IdentificationNumber *string `json:"identification_number"`
	FullName             *string `json:"full_name"`
	DateOfBirth          *string `json:"date_of_birth"`
	Gender               *string `json:"gender"`
	Nationality          *string `json:"nationality"`
	PermanentAddress     *string `json:"permanent_address"`
	TemporaryAddress     *string `json:"temporary_address"`
}

func (u *UserIdentity) TableName() string {
	return "user_identities"
}
