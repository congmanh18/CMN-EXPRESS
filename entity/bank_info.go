package entity

import "express_be/core/record"

type BankInfo struct {
	record.BaseEntity
	UserID           *string `gorm:"foreignKey:UserID" json:"user_id"`
	DeliveryPersonID *string `gorm:"foreignKey:DeliveryPersonID" json:"delivery_person_id"`
	CustomerID       *string `gorm:"foreignKey:CustomerID" json:"customer_id"`

	BankName          *string `json:"bank_name"`
	AccountNumber     *string `json:"account_number"`
	AccountHolderName *string `json:"account_holder_name"`
	BankBranch        *string `json:"bank_branch"`
}

func (b *BankInfo) TableName() string {
	return "bank_infos"
}
