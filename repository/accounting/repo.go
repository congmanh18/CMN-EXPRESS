package accounting

import (
	"context"
	"express_be/core/db/postgresql"
)

type Repo interface {
	CreateAccounting(ctx context.Context, accounting *Accounting) error
	FindByPhone(ctx context.Context, phone *string) (*Accounting, error)
}

type accountingImpl struct {
	DB *postgresql.Database
}

// Create implements Repo.
func (a *accountingImpl) CreateAccounting(ctx context.Context, accounting *Accounting) error {
	return a.DB.Executor.WithContext(ctx).Debug().Create(&accounting).Error
}

// FindByPhone implements Repo.
func (a *accountingImpl) FindByPhone(ctx context.Context, phone *string) (*Accounting, error) {
	var result *Accounting
	query := a.DB.Executor.WithContext(ctx).
		Model(&Accounting{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func NewRepo(db *postgresql.Database) Repo {
	return &accountingImpl{DB: db}
}
