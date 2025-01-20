package entity

import "express_be/core/record"

type BasicPrice struct {
	record.BaseEntity
	Region    *string  `json:"region"`
	Area      *string  `json:"area"`
	BasePrice *float64 `json:"base_price"`
	// SpecialRates []SpecialRate `gorm:"foreignKey:ShippingRateID" json:"special_rates"`
}

func (b *BasicPrice) TableName() string {
	return "basic_prices"
}
