package entity

import (
	"express_be/core/record"
)

type AccountStatus string
type ApprovalStatus string
type Role string

func (r Role) String() string {
	return string(r)
}

const (
	Pending   AccountStatus = "pending"
	Verified  AccountStatus = "verified"
	Blocked   AccountStatus = "blocked"
	Suspended AccountStatus = "suspended"

	Accepted ApprovalStatus = "accept"
	Denied   ApprovalStatus = "deny"

	AdminRole          Role = "admin"
	CustomerRole       Role = "customer"
	DeliveryPersonRole Role = "delivery_person"
	AccountingRole     Role = "accounting"
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
	Status               AccountStatus
	ApprovalStatus       ApprovalStatus
	Admin                Admin          `gorm:"foreignKey:Phone;references:Phone"`
	Customer             Customer       `gorm:"foreignKey:Phone;references:Phone"`
	DeliveryPerson       DeliveryPerson `gorm:"foreignKey:Phone;references:Phone"`
	Accounting           Accounting     `gorm:"foreignKey:Phone;references:Phone"`
	Participants         []Participant  `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) TableName() string {
	return "users"
}
