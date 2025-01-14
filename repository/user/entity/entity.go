package entity

import (
	"express_be/core/record"
	accounting "express_be/repository/accounting/entity"
	admin "express_be/repository/admin/entity"
	customer "express_be/repository/customer/entity"
	delivery "express_be/repository/delivery/entity"
)

type Status string
type ApprovalStatus string
type Role string

func (r Role) String() string {
	return string(r)
}

const (
	Pending   Status = "pending"
	Verified  Status = "verified"
	Blocked   Status = "blocked"
	Active    Status = "active"
	Inactive  Status = "inactive"
	Suspended Status = "suspended"
	OnDuty    Status = "on_duty"
	OffDuty   Status = "off_duty"

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
	SpecificAddress      *string
	Ward                 *string
	District             *string
	City                 *string
	IDCard               *string
	IdentificationNumber *string
	LastName             *string
	MiddleName           *string
	FirstName            *string
	DateOfBirth          *string
	Gender               *string
	Nationality          *string
	PlaceOfOrigin        *string
	PlaceOfResidence     *string
	Role                 Role
	Status               Status
	ApprovalStatus       ApprovalStatus
	Admin                admin.Admin             `gorm:"foreignKey:Phone;references:Phone"`
	Customer             customer.Customer       `gorm:"foreignKey:Phone;references:Phone"` // Thiết lập One-to-One
	DeliveryPerson       delivery.DeliveryPerson `gorm:"foreignKey:Phone;references:Phone"`
	Accounting           accounting.Accounting   `gorm:"foreignKey:Phone;references:Phone"`
}

func (u *User) TableName() string {
	return "users"
}
