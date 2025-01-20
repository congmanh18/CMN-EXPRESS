package user

import (
	"context"
	"express_be/entity"
)

func (c *userImpl) FindByPhone(ctx context.Context, phone *string) (*entity.User, error) {
	var result *entity.User
	query := c.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Where("phone = ?", phone).
		First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
