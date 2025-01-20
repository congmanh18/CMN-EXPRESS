package user

import (
	"context"
	"express_be/entity"
)

func (d *userImpl) FindByID(ctx context.Context, id *string) (*entity.User, error) {
	var result entity.User
	query := d.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", id).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return &result, nil
}
