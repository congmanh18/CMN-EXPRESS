package server

import (
	"express_be/core/transport/http/method"
	"express_be/core/transport/http/route"

	// customerHandler "express_be/handler/customer"
	// deliveryPersonHandler "express_be/handler/delivery"
	userHandler "express_be/handler/user"

	"express_be/handler/auth"
)

func SetupRoutes(
	userHandler userHandler.Handler,
	authhandler auth.Handler,
	jwtSecret string,
) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/admin",
			Routes: []route.Route{},
		},
		{
			// Middlewares: []echo.MiddlewareFunc{
			// 	authen.JWTMiddleware(jwtSecret),
			// },
			Routes: []route.Route{
				{
					Path:    "/refresh-token",
					Method:  method.POST,
					Handler: authhandler.HandleRefreshToken,
				},
				{
					Path:    "/login",
					Method:  method.POST,
					Handler: authhandler.HandleLogin,
				},
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: authhandler.HandleRegister,
				},
				{
					Path:    "/reset-password",
					Method:  method.PATCH,
					Handler: authhandler.HandleChangePassword,
				},
				{
					Path:    "/users/:id",
					Method:  method.PATCH,
					Handler: userHandler.HandleUpdateUserStatus,
				},
				{
					Path:    "/users",
					Method:  method.GET,
					Handler: userHandler.HandleListUsers,
				},
			},
		},
		// {
		// 	Prefix: "/accountings",
		// 	Routes: []route.Route{
		// 		{
		// 			Path:    "/register",
		// 			Method:  method.PATCH,
		// 			Handler: adminHandler.HandleUpdateDeliveryPersonStatus,
		// 		},
		// 	},
		// },
		// {
		// 	Prefix: "/customers",
		// 	Routes: []route.Route{
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.GET,
		// 			Handler: customerHandler.HandleGetInfoCustomer,
		// 		},
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.DELETE,
		// 			Handler: customerHandler.HandleDeleteCustomer,
		// 		},
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.PUT,
		// 			Handler: customerHandler.HandleUpdateCustomer,
		// 		},
		// 	},
		// },
		// {
		// 	Prefix: "/delivery-persons",
		// 	Routes: []route.Route{
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.GET,
		// 			Handler: deliveryPersonHandler.HandleGetInfoDeliveryPerson,
		// 		},
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.DELETE,
		// 			Handler: deliveryPersonHandler.HandleDeleteDeliveryPerson,
		// 		},
		// 		{
		// 			Path:    "/:id",
		// 			Method:  method.PUT,
		// 			Handler: deliveryPersonHandler.HandleUpdateDeliveryPerson,
		// 		},
		// 	},
		// },
	}
}
