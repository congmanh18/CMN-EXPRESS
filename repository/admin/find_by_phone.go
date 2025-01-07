package admin

import (
	"context"
	"express_be/repository/admin/entity"
)

// FindByPhone implements Repo.
func (a *adminImpl) FindByPhone(ctx context.Context, phone *string) (*entity.Admin, error) {
	var result entity.Admin
	query := a.DB.Executor.WithContext(ctx).
		Model(&entity.Admin{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
