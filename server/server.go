package server

import (
	"express_be/conf"
	"express_be/core/config"
	httpServer "express_be/core/transport/http"
	"express_be/core/transport/http/engine"
	"express_be/core/transport/http/route"
	"express_be/docs"
	"express_be/migration"
	"express_be/provider"

	authHandler "express_be/handler/auth"
	orderHandler "express_be/handler/order"
	priceHandler "express_be/handler/price"
	userHandler "express_be/handler/user"

	accountingRepo "express_be/repository/accounting"
	adminRepo "express_be/repository/admin"
	customerRepo "express_be/repository/customer"
	deliveryPersonRepo "express_be/repository/delivery"
	orderRepo "express_be/repository/order"
	priceRepo "express_be/repository/price"
	"express_be/repository/token"
	userRepo "express_be/repository/user"
	"express_be/usecase/auth"

	"express_be/usecase/order"
	"express_be/usecase/price"
	"express_be/usecase/user"

	echoSwagger "github.com/swaggo/echo-swagger"
)

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

func NewServer(serviceConf conf.ServiceConfig, routes []route.GroupRoute) *httpServer.Server {
	e := engine.NewEcho()
	if serviceConf.Environment == "production" {
		docs.SwaggerInfo.Host = "203.145.47.225" // Host cho production
	} else {
		docs.SwaggerInfo.Host = "localhost:4579" // Host cho local
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return httpServer.NewHttpServer(
		httpServer.AddName(serviceConf.ServiceName),
		httpServer.AddPort(serviceConf.ServicePort),
		httpServer.AddEngine(e),
		httpServer.AddGroupRoutes(routes),
	)
}

func Run(confPath string) {

	serviceConf := LoadServiceConfig(confPath)
	appProvider := provider.NewAppProvider(serviceConf)
	RunMigration(appProvider, serviceConf.EnableMigrations)

	// Khởi tạo Repo
	userRepo := userRepo.NewRepo(appProvider.Postgres)
	orderRepo := orderRepo.NewRepo(appProvider.Postgres)
	adminRepo := adminRepo.NewRepo(appProvider.Postgres)
	accountingRepo := accountingRepo.NewRepo(appProvider.Postgres)
	custRepo := customerRepo.NewRepo(appProvider.Postgres)
	deliRepo := deliveryPersonRepo.NewRepo(appProvider.Postgres)
	tokenRepo := token.NewRepo(appProvider.Postgres)
	priceRepo := priceRepo.NewRepo(appProvider.Postgres)

	// Khởi tạo Usecase
	authUsecase := auth.NewAuthUsecase(userRepo, adminRepo, custRepo, deliRepo, accountingRepo, tokenRepo, serviceConf.JwtSecretKey)
	orderUsecase := order.NewOrderUsecase(orderRepo)
	userUsecase := user.NewUserUsecase(userRepo, custRepo, deliRepo)
	priceUsecase := price.NewPriceUsecase(priceRepo)

	// Khởi tạo handler
	userHandl := userHandler.NewHandler(userHandler.HandlerInject{
		UserUsecase: userUsecase,
	})
	authHandl := authHandler.NewHandler(authHandler.HandlerInject{
		AuthUsecase: authUsecase,
	})
	orderHandl := orderHandler.NewHandler(orderHandler.HandlerInject{
		OrderUsecase: orderUsecase,
	})

	priceHandl := priceHandler.NewHandler(priceHandler.HandlerInject{
		PriceUsecase: priceUsecase,
	})

	httpRoutes := SetupRoutes(userHandl, authHandl, priceHandl, orderHandl, serviceConf.JwtSecretKey)

	s := NewServer(serviceConf, httpRoutes)
	s.Run()
}
