package server

import (
	"express_be/conf"
	"express_be/core/config"
	httpServer "express_be/core/transport/http"
	"express_be/core/transport/http/engine"
	"express_be/core/transport/http/route"
	"express_be/migration"
	"express_be/provider"

	adminHandler "express_be/handler/admin"
	authHandler "express_be/handler/auth"
	customerHandler "express_be/handler/customer"
	deliveryPersonHandler "express_be/handler/delivery"

	accountingRepo "express_be/repository/accounting"
	adminRepo "express_be/repository/admin"
	customerRepo "express_be/repository/customer"
	deliveryPersonRepo "express_be/repository/delivery"
	"express_be/repository/token"

	"express_be/usecase/accounting"
	"express_be/usecase/admin"
	"express_be/usecase/auth"
	"express_be/usecase/customer"
	"express_be/usecase/delivery"

	echoSwagger "github.com/swaggo/echo-swagger"
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

func NewServer(serviceConf conf.ServiceConfig, routes []route.GroupRoute) *httpServer.Server {
	e := engine.NewEcho()
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
	RunMigration(appProvider, enableMigrations)

	// firebasAuth, err := auth.NewFirebaseAuthService("./conf/express-2227f-firebase-adminsdk-vbu9s-83580ceea5.json")
	// if err != nil {
	// 	panic(err)
	// }

	jwtSecret := serviceConf.JwtSecretKey

	// Khởi tạo Repo
	custRepo := customerRepo.NewRepo(appProvider.Postgres)
	deliRepo := deliveryPersonRepo.NewRepo(appProvider.Postgres)
	adminRepo := adminRepo.NewRepo(appProvider.Postgres)
	tokenRepo := token.NewRepo(appProvider.Postgres)
	accountingRepo := accountingRepo.NewRepo(appProvider.Postgres)

	// Khởi tạo Usecase
	authUsecase := auth.NewAuthUsecase(adminRepo, custRepo, deliRepo, accountingRepo, tokenRepo)
	adminUsecase := admin.NewAdminUsecase(adminRepo)
	adminCustomerUsecase := customer.NewAdminUsecase(custRepo)
	adminDeliveryPersonUsecase := delivery.NewAdminUsecase(deliRepo)
	adminAccountingUsecase := accounting.NewAdminUsecase(accountingRepo)

	customerUsecase := customer.NewCustomerUsecase(custRepo)
	deliveryUsecase := delivery.NewDeliveryPersonUsecase(deliRepo)

	// Khởi tạo handler
	adminHandl := adminHandler.NewHandler(adminHandler.HandlerInject{
		AdminCustomerUsecase:       adminCustomerUsecase,
		AdminDeliveryPersonUsecase: adminDeliveryPersonUsecase,
		AdminAccountingUsecase:     adminAccountingUsecase,
		AdminUsecase:               adminUsecase,
	})
	customerHandl := customerHandler.NewHandler(customerHandler.HandlerInject{
		CustomerUsecase: customerUsecase,
	})
	deliveryHandl := deliveryPersonHandler.NewHandler(deliveryPersonHandler.HandlerInject{
		DeliveryPersonUsecase: deliveryUsecase,
	})
	authHandl := authHandler.NewHandler(authHandler.HandlerInject{
		AuthUsecase: authUsecase,
	})

	// Khởi tạo routes
	routes := SetupRoutes(adminHandl, customerHandl, deliveryHandl, authHandl, jwtSecret)

	s := NewServer(serviceConf, routes)
	s.Run()
}
