package customer

import (
	"context"

	"gorm.io/gorm"
)

func (c *customerImpl) FindByPhone(ctx context.Context, phone *string) (*string, error) {
	var result string
	err := c.DB.Executor.WithContext(ctx).
		Table("customers").
		Select("phone").
		Where("phone = ? AND deleted_at IS NULL", phone).
		Limit(1).
		Scan(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { // Nếu không tìm thấy bản ghi
			return nil, nil
		}
		return nil, err // Lỗi khác
	}

	if result == "" {
		return nil, nil // Trả về nil nếu chuỗi rỗng
	}

	return &result, nil

}
