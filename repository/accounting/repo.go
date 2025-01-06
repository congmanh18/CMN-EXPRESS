package accounting

import (
	"context"
	"express_be/core/db/postgresql"
)

type Repo interface {
	CreateAccounting(ctx context.Context, accounting *Accounting) error
	FindByPhone(ctx context.Context, phone *string) (*Accounting, error)
	ChangePassword(ctx context.Context, phone *string, newPassword *string) error
}

type accountingImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &accountingImpl{DB: db}
}
