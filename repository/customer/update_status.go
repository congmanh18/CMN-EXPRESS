package customer

import (
	"context"
	"errors"
	"express_be/repository/customer/entity"
)

func (c *customerImpl) UpdateStatus(ctx context.Context, id *string, status *entity.Status) error {
	if id == nil {
		return errors.New("id cannot be nil")
	}

	// Cập nhật trực tiếp status mà không cần struct
	err := c.DB.Executor.WithContext(ctx).
		Model(&entity.Customer{}).
		Where("id = ?", id).
		Update("status", status).Error

	if err != nil {
		return err // Trả lỗi nếu có vấn đề
	}

	return nil // Thành công
}
