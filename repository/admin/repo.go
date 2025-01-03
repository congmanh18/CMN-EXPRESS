package admin

import (
	"context"
	"express_be/core/db/postgresql"
)

type Repo interface {
	CreateAdmin(ctx context.Context, admin *Admin) error
	FindByPhone(ctx context.Context, phone *string) (*Admin, error)
}

type adminImple struct {
	DB *postgresql.Database
}

// CreateAdmin implements Repo.
func (a *adminImple) CreateAdmin(ctx context.Context, admin *Admin) error {
	return a.DB.Executor.WithContext(ctx).Debug().Create(&admin).Error
}

// FindByPhone implements Repo.
func (a *adminImple) FindByPhone(ctx context.Context, phone *string) (*Admin, error) {
	var result *Admin
	query := a.DB.Executor.WithContext(ctx).
		Model(&Admin{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func NewRepo(db *postgresql.Database) Repo {
	return &adminImple{DB: db}
}
