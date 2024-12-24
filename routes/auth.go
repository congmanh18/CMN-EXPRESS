package routes

import (
	"express_be/core/transport/http/method"
	"express_be/core/transport/http/route"
	"express_be/handler/auth"
)

func SetupAuthRoutes(handler auth.Handler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/auth",
			Routes: []route.Route{
				{
					Path:    "/customer/register",
					Method:  method.POST,
					Handler: handler.HandleCustomerRegistration,
				},
				{
					Path:    "/delivery/register",
					Method:  method.POST,
					Handler: handler.HandleDeliveryPersonRegistration,
				},
			},
		},
	}
}
