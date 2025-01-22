package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Connection struct {
	Addr     string
	Password string
	DB       int
}

func New(conn Connection) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conn.Addr,
		Password:     conn.Password,
		DB:           conn.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Kiểm tra kết nối
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis successfully")
	return rdb, nil
}
