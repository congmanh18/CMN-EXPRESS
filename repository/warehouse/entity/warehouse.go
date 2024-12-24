package entity

import "express_be/core/record"

type Warehouse struct {
	record.BaseEntity
	Address              *string                `json:"address"`
	Phone                *string                `json:"phone"`
	ManageName           *string                `json:"manage_name"`
	WarehouseTransaction []WarehouseTransaction `gorm:"foreignKey:WarehouseID" json:"warehouse_transaction"`
}

func (w *Warehouse) TableName() string {
	return "warehouses"
}
