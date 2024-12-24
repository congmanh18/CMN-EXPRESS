package server

import (
	"express_be/conf"
	"express_be/core/transport/http/route"
	"express_be/handler/auth"
	customerRepo "express_be/repository/customer"
	deliveryPersonRepo "express_be/repository/delivery"
	authUseCase "express_be/usecase/auth"

	"express_be/provider"
	"express_be/routes"
)

func InitializeAppProvider(serviceConf conf.ServiceConfig) *provider.AppProvider {
	return provider.NewAppProvider(serviceConf)
}

func InitializeRoutes(auth auth.Handler) []route.GroupRoute {
	return routes.SetupAuthRoutes(auth)
}

func InitializeHandler(appProvider *provider.AppProvider) auth.Handler {
	customerRepo := customerRepo.NewRepo(appProvider.Postgres)
	deliveryPersonRepo := deliveryPersonRepo.NewRepo(appProvider.Postgres)

	return auth.NewHandler(auth.HandlerInject{
		RegisterCustomerUseCase:       authUseCase.NewRegisterCustomerUseCase(customerRepo),
		RegisterDeliveryPersonUseCase: authUseCase.NewRegisterDeliveryPersonUseCase(deliveryPersonRepo),
	})
}
