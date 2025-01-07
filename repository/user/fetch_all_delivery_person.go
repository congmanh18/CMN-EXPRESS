package user

import (
	"context"
	"express_be/repository/user/entity"
)

func (c *userImpl) FetchAllDeliveryPerson(ctx context.Context, page, pageSize *int) ([]entity.DeliveryPersonDetails, error) {
	var result []entity.DeliveryPersonDetails
	offset := (*page - 1) * *pageSize

	query := c.DB.Executor.
		WithContext(ctx).
		Table("users").
		Select(`
			users.*,
			delivery_persons.salary_rate
		`).
		Joins("JOIN delivery_persons ON delivery_persons.phone = users.phone"). // Liên kết bảng
		Where("users.role = ?", entity.DeliveryPerson).
		Order("users.full_name ASC").
		Offset(offset).
		Limit(*pageSize).
		Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
