package customer

import (
	"context"
	"errors"
	"express_be/entity"

	"gorm.io/gorm"
)

func (c *customerImpl) FindByID(ctx context.Context, id *string) (*entity.Customer, error) {
	var result entity.Customer
	query := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("id = ?", *id).
		First(&result)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, query.Statement.Error
	}

	return &result, nil
}
