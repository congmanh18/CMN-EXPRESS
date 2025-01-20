package migration

import (
	"express_be/entity"

	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Printf("Migration.....")
	db.Exec("CREATE SCHEMA IF NOT EXISTS log")
	err := db.AutoMigrate(
		&entity.RefreshToken{},
		&entity.User{},
		&entity.Admin{},
		&entity.Accounting{},
		&entity.Customer{},
		&entity.DeliveryPerson{},
		&entity.Salary{},

		&entity.Conversation{},
		&entity.Message{},
		&entity.Participant{},

		// &user.BankInfo{},

		&entity.Order{},
		// &order.OrderStatus{},
		// &order.CODTransaction{},
		// &order.CODReconciliation{},
		// &order.AdditionalService{},

		// &warehouse.Warehouse{},
		// &warehouse.WarehouseTransaction{},
		&entity.BasicPrice{},
		&entity.LogPrice{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
