package user

import (
	"context"
	"express_be/repository/user/entity"
)

// CheckRoleByID implements Repo.
func (c *userImpl) CheckRoleByID(ctx context.Context, id *string) (*string, error) {
	var role string
	query := c.DB.Executor.WithContext(ctx).
		Model(&entity.User{}).
		Select("role").
		Where("id = ?", id).
		First(&role)

	if query.Error != nil {
		return nil, query.Error
	}
	return &role, nil
}
