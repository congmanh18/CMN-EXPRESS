package entity

import (
	"express_be/core/record"
	"express_be/repository/accounting"
	"express_be/repository/admin"
	customer "express_be/repository/customer/entity"
	delivery "express_be/repository/delivery/entity"
)

type Status string
type ApprovalStatus string
type Role string

const (
	Pending   Status = "pending"
	Verified  Status = "verified"
	Blocked   Status = "blocked"
	Active    Status = "active"
	Inactive  Status = "inactive"
	Suspended Status = "suspended"
	OnDuty    Status = "on-duty"
	OffDuty   Status = "off-duty"

	Accepted ApprovalStatus = "accept"
	Denied   ApprovalStatus = "deny"

	Admin          Role = "admin"
	Customer       Role = "customer"
	DeliveryPerson Role = "delivery_person"
	Accounting     Role = "accounting"
)

type User struct {
	record.BaseEntity
	Phone                *string `gorm:"uniqueIndex:idx_phone_deleted_at"`
	Password             *string
	CurrentAddress       *string
	IdentificationNumber *string
	FullName             *string
	DateOfBirth          *string
	Gender               *string
	Nationality          *string
	PlaceOfOrigin        *string
	PlaceOfResidence     *string
	Role                 Role
	Status               Status
	ApprovalStatus       ApprovalStatus
	Admin                admin.Admin             `gorm:"foreignKey:Phone"`
	Customer             customer.Customer       `gorm:"foreignKey:Phone"`
	DeliveryPerson       delivery.DeliveryPerson `gorm:"foreignKey:Phone"`
	Accounting           accounting.Accounting   `gorm:"foreignKey:Phone"`
}

func (u *User) TableName() string {
	return "users"
}
