package provider

import (
	"express_be/conf"
	"express_be/core/db/postgresql"
)

type AppProvider struct {
	Postgres *postgresql.Database
	// Redis    *redis.Client
}

func NewAppProvider(config conf.ServiceConfig) *AppProvider {
	return &AppProvider{
		// Redis:    ProvideRedis(config),
		Postgres: ProvidePostgres(config),
	}
}
