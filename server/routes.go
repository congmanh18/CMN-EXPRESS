package server

import (
	"express_be/core/transport/http/method"
	"express_be/core/transport/http/route"
	adminHandler "express_be/handler/admin"
	customerHandler "express_be/handler/customer"
	deliveryPersonHandler "express_be/handler/delivery"

	"express_be/handler/auth"
)

func SetupRoutes(
	adminHandler adminHandler.Handler,
	customerHandler customerHandler.Handler,
	deliveryPersonHandler deliveryPersonHandler.Handler,
	authhandler auth.Handler,
	jwtSecret string,
) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/admin",
			Routes: []route.Route{
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: adminHandler.HandleRegister,
				},
				// {
				// 	Path:    "/login",
				// 	Method:  method.POST,
				// 	Handler: adminHandler.Login,
				// },
				{
					Path:    "/customers/pending",
					Method:  method.GET,
					Handler: adminHandler.HandleListPendingCustomer,
				},
				{
					Path:    "/customers/all",
					Method:  method.GET,
					Handler: adminHandler.HandleAllCustomers,
				},
				{
					Path:    "/delivery-persons/pending",
					Method:  method.GET,
					Handler: adminHandler.HandleListPendingDeliveryPerson,
				},
				{
					Path:    "/delivery-persons/all",
					Method:  method.GET,
					Handler: adminHandler.HandleAllDeliveryPersons,
				},
			},
		},
		// {
		// 	Prefix: "/auth/verify",
		// 	Middlewares: []echo.MiddlewareFunc{
		// 		authen.JWTMiddleware(jwtSecret),
		// 	},
		// 	Routes: []route.Route{
		// 		{
		// 			Path:    "/customer/:id",
		// 			Method:  method.POST,
		// 			Handler: authhandler.HandleVerifyCustomer,
		// 		},
		// 		{
		// 			Path:    "/delivery-person/:id",
		// 			Method:  method.POST,
		// 			Handler: authhandler.HandleVerifyDeliveryPerson,
		// 		},
		// 	},
		// },
		{
			Prefix: "/customers",
			Routes: []route.Route{
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: authhandler.HandleCustomerRegistration,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: customerHandler.HandleGetInfoCustomer,
				},
				{
					Path:    "/:id",
					Method:  method.DELETE,
					Handler: customerHandler.HandleDeleteCustomer,
				},
				{
					Path:    "/:id",
					Method:  method.PUT,
					Handler: customerHandler.HandleUpdateCustomer,
				},
			},
		},
		{
			Prefix: "/delivery-persons",
			Routes: []route.Route{
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: authhandler.HandleDeliveryPersonRegistration,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: deliveryPersonHandler.HandleGetInfoDeliveryPerson,
				},
				{
					Path:    "/:id",
					Method:  method.DELETE,
					Handler: deliveryPersonHandler.HandleDeleteDeliveryPerson,
				},
				{
					Path:    "/:id",
					Method:  method.PUT,
					Handler: deliveryPersonHandler.HandleUpdateDeliveryPerson,
				},
			},
		},
	}
}
