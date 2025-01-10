package customer

import (
	"context"
)

func (c *customerImpl) FetchPhoneWithCache(ctx context.Context, phone *string) (bool, error) {
	// if phone == nil {
	// 	return false, errors.New("phone cannot be nil")
	// }

	// cacheKey := fmt.Sprintf("phone:%s", *phone)

	// // Kiểm tra cache
	// cachedValue, err := c.Cache.Get(ctx, cacheKey).Result()
	// if err == nil {
	// 	if cachedValue == "exists" {
	// 		return true, nil
	// 	} else if cachedValue == "not_exists" {
	// 		return false, nil
	// 	}
	// } else if err != redis.Nil {
	// 	log.Printf("Cache error: %v", err)
	// }

	// // Truy vấn cơ sở dữ liệu
	// var count int64
	// query := c.DB.Executor.WithContext(ctx).
	// 	Model(&entity.Customer{}).
	// 	Where("phone = ?", *phone).
	// 	Count(&count)

	// if query.Error != nil {
	// 	return false, query.Error
	// }

	// exists := count > 0

	// // Lưu kết quả vào cache
	// if exists {
	// 	c.Cache.Set(ctx, cacheKey, "exists", 10*time.Minute)
	// } else {
	// 	c.Cache.Set(ctx, cacheKey, "not_exists", 10*time.Minute)
	// }

	panic("unreachable")
}
