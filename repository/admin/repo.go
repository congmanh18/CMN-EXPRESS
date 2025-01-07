package admin

import (
	"context"
	"express_be/core/db/postgresql"
)

type Repo interface {
	CreateAdmin(ctx context.Context, admin *Admin) error
	FindByPhone(ctx context.Context, phone *string) (*Admin, error)
}

type adminImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &adminImpl{DB: db}
}
