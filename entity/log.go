package entity

import "express_be/core/record"

type LogPrice struct {
	record.BaseEntity
	AdminID *string // ID của admin thực hiện chỉnh sửa
	Action  *string // Hành động (VD: "Cập nhật giá", "Thêm khu vực mới")
	Details *string // Chi tiết chỉnh sửa (VD: "Thay đổi giá Nội thành từ 20,000 VNĐ thành 25,000 VNĐ")
}

func (l *LogPrice) TableName() string {
	return "log.prices"
}
