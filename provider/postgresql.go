package provider

import (
	"express_be/conf"
	"fmt"

	"express_be/core/db/postgresql"
)

func ProvidePostgres(config conf.ServiceConfig) *postgresql.Database {
	fmt.Println("Connecting to PostgreSQL...")
	return postgresql.New(postgresql.Connection{
		Host:     config.DBHost,
		Port:     config.DBPort,
		Database: config.DBName,
		User:     config.DBUser,
		Password: config.DBPwd,
		SSLMode:  postgresql.Disable,
	})
}

// Disable - Require
