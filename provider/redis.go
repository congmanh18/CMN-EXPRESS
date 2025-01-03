package provider

import (
	"express_be/conf"
	core "express_be/core/db/redis"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ProvideRedis(config conf.ServiceConfig) *redis.Client {
	fmt.Println("Connecting to Redis...")
	return core.New(core.Connection{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       config.RedisPort,
	})
}
