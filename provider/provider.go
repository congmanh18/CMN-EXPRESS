package provider

import (
	"express_be/conf"
	"express_be/core/db/postgresql"
)

type AppProvider struct {
	Postgres *postgresql.Database
}

func NewAppProvider(config conf.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
	}
}
