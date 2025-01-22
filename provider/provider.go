package provider

import (
	"express_be/conf"
	"express_be/core/db/postgresql"
	// "github.com/redis/go-redis/v9"
)

type AppProvider struct {
	Postgres *postgresql.Database
	// Redis    *redis.Client
}

func NewAppProvider(config conf.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
		// Redis:    ProvideRedis(config),
	}
}
