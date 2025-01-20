package accounting

import (
	"context"
	"express_be/entity"
)

// FindByPhone implements Repo.
func (a *accountingImpl) FindByPhone(ctx context.Context, phone *string) (*entity.Accounting, error) {
	var result *entity.Accounting
	query := a.DB.Executor.WithContext(ctx).
		Model(&entity.Accounting{}).
		Where("phone =?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
