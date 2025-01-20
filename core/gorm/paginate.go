package gorm

import (
	"express_be/core/record"

	"gorm.io/gorm"
)

func Paginate(pagination *record.Pagination, txCountTotalRows *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return nil
}

func CalculateOffset(page, limit int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * limit
}
