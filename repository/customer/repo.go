package customer

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/repository/customer/entity"
)

type Repo interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	FindByPhone(ctx context.Context, phone *string) (*string, error) // Chưa phân trang
}

type customerImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &customerImpl{
		DB: db,
	}
}
