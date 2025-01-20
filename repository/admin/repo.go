package admin

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, admin *entity.Admin) error
	FindByPhone(ctx context.Context, phone *string) (*entity.Admin, error)
}

type adminImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &adminImpl{DB: db}
}
