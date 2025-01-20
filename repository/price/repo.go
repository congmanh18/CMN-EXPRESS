package price

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/core/gorm"
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/entity"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, price *entity.BasicPrice, adminID *string) error
	ReadAllPrice(ctx context.Context) ([]entity.BasicPrice, error)
	Update(ctx context.Context, id *string, price *entity.BasicPrice, adminID *string) error
	Delete(ctx context.Context, id *string, adminID *string) error
}

type priceImpl struct {
	DB *postgresql.Database
}

// Create implements Repo.
func (p *priceImpl) Create(ctx context.Context, price *entity.BasicPrice, adminID *string) error {
	// Tạo giá mới
	if err := p.DB.Executor.WithContext(ctx).Create(price).Error; err != nil {
		return fmt.Errorf("failed to create new price: %w", err)
	}

	// Ghi log thao tác
	log := entity.LogPrice{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		AdminID: adminID,
		Action:  pointer.String("Create"),
		Details: pointer.String(fmt.Sprintf("Created new price: %+v", price)),
	}

	if err := p.DB.Executor.WithContext(ctx).Create(&log).Error; err != nil {
		return fmt.Errorf("failed to log create action: %w", err)
	}

	return nil
}

// ReadAllPrice implements Repo.
func (p *priceImpl) ReadAllPrice(ctx context.Context) ([]entity.BasicPrice, error) {
	var result []entity.BasicPrice
	if err := p.DB.Executor.WithContext(ctx).
		Order("region ASC").
		Order("base_price DESC"). // Sắp xếp theo giá giảm dần nếu region trùng
		Find(&result).Error; err != nil {
		return nil, fmt.Errorf("failed to read all prices: %w", err)
	}

	return result, nil
}

// Update implements Repo.
func (p *priceImpl) Update(ctx context.Context, id *string, price *entity.BasicPrice, adminID *string) error {
	// Bỏ qua các field có giá trị zero
	omitFields := gorm.OmitFields(price, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero()
	})

	// Cập nhật giá
	err := p.DB.Executor.WithContext(ctx).
		Model(&entity.BasicPrice{}).
		Where("id = ?", *id).
		Omit(omitFields...).
		Updates(price).Error

	if err != nil {
		return fmt.Errorf("failed to update price (ID: %s): %w", *id, err)
	}

	// Ghi log thao tác
	log := entity.LogPrice{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		AdminID: adminID,
		Action:  pointer.String("Update"),
		Details: pointer.String(fmt.Sprintf("Updated price (ID: %s) with data: %+v", *id, price)),
	}

	if err := p.DB.Executor.WithContext(ctx).Create(&log).Error; err != nil {
		return fmt.Errorf("failed to log update action: %w", err)
	}

	return nil
}

// Delete implements Repo.
func (p *priceImpl) Delete(ctx context.Context, id *string, adminID *string) error {
	// Xóa giá theo ID
	err := p.DB.Executor.WithContext(ctx).
		Where("id = ?", *id).
		Delete(&entity.BasicPrice{}).Error

	if err != nil {
		return fmt.Errorf("failed to delete price (ID: %s): %w", *id, err)
	}

	// Ghi log thao tác
	log := entity.LogPrice{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		AdminID: adminID,
		Action:  pointer.String("Delete"),
		Details: pointer.String(fmt.Sprintf("Deleted price with ID: %s", *id)),
	}

	if err := p.DB.Executor.WithContext(ctx).Create(&log).Error; err != nil {
		return fmt.Errorf("failed to log delete action: %w", err)
	}

	return nil
}
func NewRepo(db *postgresql.Database) Repo {
	return &priceImpl{DB: db}
}
