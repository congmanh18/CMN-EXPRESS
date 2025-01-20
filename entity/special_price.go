package entity

import (
	"express_be/core/record"
	"time"
)

type SpecialRate struct {
	record.BaseEntity
	ShippingRateID      *string   // Liên kết đến bảng "ShippingRate"
	DiscountPercentage  *float64  // Phần trăm giảm giá (VD: 10.00 cho 10%)
	SurchargePercentage *float64  // Phần trăm phụ phí (VD: 20.00 cho 20%)
	StartDate           time.Time // Ngày bắt đầu áp dụng
	EndDate             time.Time // Ngày kết thúc áp dụng
	Description         *string   // Mô tả (VD: "Phụ phí ngày lễ", "Giảm giá Tết")
}
