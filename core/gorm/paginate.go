package gorm

import (
	"express_be/core/record"

	"gorm.io/gorm"
)

func Paginate(pagination *record.Pagination, txCountTotalRows *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return nil
}
