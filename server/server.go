package server

import (
	"express_be/conf"
	"express_be/core/config"
	"express_be/core/transport/http"
	"express_be/core/transport/http/engine"
	"express_be/core/transport/http/route"
	"express_be/migration"
	"express_be/provider"
)

const enableMigrations = true

func RunMigration(appProvider *provider.AppProvider, enableMigrate bool) {
	if enableMigrate {
		migration.Migration(appProvider.Postgres.Executor)
	}
}

func LoadServiceConfig(confPath string) conf.ServiceConfig {
	var serviceConf conf.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)
	return serviceConf
}

func NewServer(serviceConf conf.ServiceConfig, routes []route.GroupRoute) *http.Server {
	e := engine.NewEcho()
	return http.NewHttpServer(
		http.AddName(serviceConf.ServiceName),
		http.AddPort(serviceConf.ServicePort),
		http.AddEngine(e),
		http.AddGroupRoutes(routes),
	)
}

func Run(confPath string) {
	serviceConf := LoadServiceConfig(confPath)
	appProvider := InitializeAppProvider(serviceConf)
	handler := InitializeHandler(appProvider)
	routes := InitializeRoutes(handler)
	server := NewServer(serviceConf, routes)

	RunMigration(appProvider, enableMigrations)
	server.Run()
}
