package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Connection struct {
	Addr     string
	Password string
	DB       int
}

func New(conn Connection) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr,     // Địa chỉ Redis
		Password: conn.Password, // Mật khẩu Redis (để trống nếu không có)
		DB:       conn.DB,       // Sử dụng database mặc định
	})

	// Kiểm tra kết nối
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb
}
