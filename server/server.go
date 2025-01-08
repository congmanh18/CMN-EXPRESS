package server

import (
	"express_be/conf"
	"express_be/core/config"
	httpServer "express_be/core/transport/http"
	"express_be/core/transport/http/engine"
	"express_be/core/transport/http/route"
	"express_be/migration"
	"express_be/provider"

	authHandler "express_be/handler/auth"
	userHandler "express_be/handler/user"

	accountingRepo "express_be/repository/accounting"
	adminRepo "express_be/repository/admin"
	customerRepo "express_be/repository/customer"
	deliveryPersonRepo "express_be/repository/delivery"
	"express_be/repository/token"
	userRepo "express_be/repository/user"

	"express_be/usecase/auth"
	"express_be/usecase/user"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const enableMigrations = false

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
	userRepo := userRepo.NewRepo(appProvider.Postgres)
	adminRepo := adminRepo.NewRepo(appProvider.Postgres)
	accountingRepo := accountingRepo.NewRepo(appProvider.Postgres)
	custRepo := customerRepo.NewRepo(appProvider.Postgres)
	deliRepo := deliveryPersonRepo.NewRepo(appProvider.Postgres)
	tokenRepo := token.NewRepo(appProvider.Postgres)

	// Khởi tạo Usecase
	authUsecase := auth.NewAuthUsecase(userRepo, adminRepo, custRepo, deliRepo, accountingRepo, tokenRepo)
	userUsecase := user.NewUserUsecase(userRepo, custRepo, deliRepo)

	// Khởi tạo handler
	userHandl := userHandler.NewHandler(userHandler.HandlerInject{
		UserUsecase: userUsecase,
	})
	// customerHandl := customerHandler.NewHandler(customerHandler.HandlerInject{
	// 	CustomerUsecase: customerUsecase,
	// })
	// deliveryHandl := deliveryPersonHandler.NewHandler(deliveryPersonHandler.HandlerInject{
	// 	DeliveryPersonUsecase: deliveryUsecase,
	// })
	authHandl := authHandler.NewHandler(authHandler.HandlerInject{
		AuthUsecase: authUsecase,
	})

	// Khởi tạo routes
	routes := SetupRoutes(userHandl, authHandl, jwtSecret)

	s := NewServer(serviceConf, routes)
	s.Run()
}
