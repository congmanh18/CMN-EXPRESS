package warehouse

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	Create(ctx context.Context, warehouse *entity.Warehouse, adminID *string) error
	ReadAllWarehouse(ctx context.Context) ([]*entity.Warehouse, error)
	Update(ctx context.Context, id *string, warehouse *entity.Warehouse, adminID *string) error
	Delete(ctx context.Context, id *string, adminID *string) error
}

type warehouseImpl struct {
	DB *postgresql.Database
}

// Create implements Repo.
func (w *warehouseImpl) Create(ctx context.Context, warehouse *entity.Warehouse, adminID *string) error {
	panic("unimplemented")
}

// Delete implements Repo.
func (w *warehouseImpl) Delete(ctx context.Context, id *string, adminID *string) error {
	panic("unimplemented")
}

// ReadAllWarehouse implements Repo.
func (w *warehouseImpl) ReadAllWarehouse(ctx context.Context) ([]*entity.Warehouse, error) {
	panic("unimplemented")
}

// Update implements Repo.
func (w *warehouseImpl) Update(ctx context.Context, id *string, warehouse *entity.Warehouse, adminID *string) error {
	panic("unimplemented")
}

func NewRepo(db *postgresql.Database) Repo {
	return &warehouseImpl{DB: db}
}
