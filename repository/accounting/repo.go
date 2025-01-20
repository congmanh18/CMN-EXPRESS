package accounting

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, accounting *entity.Accounting) error
	FindByPhone(ctx context.Context, phone *string) (*entity.Accounting, error)
}

type accountingImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &accountingImpl{DB: db}
}
