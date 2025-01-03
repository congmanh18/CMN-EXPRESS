package migration

import (
	"express_be/repository/admin"
	customer "express_be/repository/customer/entity"
	delivery "express_be/repository/delivery/entity"
	identity "express_be/repository/identity/entity"
	order "express_be/repository/order/entity"
	warehouse "express_be/repository/warehouse/entity"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Printf("Migration.....")
	err := db.AutoMigrate(
		&admin.Admin{},
		&customer.Customer{},
		&delivery.DeliveryPerson{},
		&delivery.Salary{},

		&identity.BankInfo{},

		&order.Order{},
		&order.OrderStatus{},
		&order.CODTransaction{},
		&order.CODReconciliation{},
		&order.AdditionalService{},

		&warehouse.Warehouse{},
		&warehouse.WarehouseTransaction{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
