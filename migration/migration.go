package migration

import (
	accounting "express_be/repository/accounting/entity"
	admin "express_be/repository/admin/entity"
	customer "express_be/repository/customer/entity"
	delivery "express_be/repository/delivery/entity"

	// identity "express_be/repository/identity/entity"
	order "express_be/repository/order/entity"
	price "express_be/repository/price/entity"

	"express_be/repository/token"
	user "express_be/repository/user/entity"

	// warehouse "express_be/repository/warehouse/entity"

	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Printf("Migration.....")
	err := db.AutoMigrate(
		&token.RefreshToken{},
		&user.User{},
		&admin.Admin{},
		&accounting.Accounting{},
		&customer.Customer{},
		&delivery.DeliveryPerson{},
		&delivery.Salary{},

		// &identity.BankInfo{},

		&order.Order{},
		// &order.OrderStatus{},
		// &order.CODTransaction{},
		// &order.CODReconciliation{},
		// &order.AdditionalService{},

		// &warehouse.Warehouse{},
		// &warehouse.WarehouseTransaction{},
		&price.BasicPrice{},
		&price.LogPrice{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
