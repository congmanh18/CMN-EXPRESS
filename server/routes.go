package server

import (
	"express_be/core/jwt"
	"express_be/core/transport/http/method"
	"express_be/core/transport/http/route"

	// customerHandler "express_be/handler/customer"
	// deliveryPersonHandler "express_be/handler/delivery"
	orderHandler "express_be/handler/order"
	priceHandler "express_be/handler/price"
	userHandler "express_be/handler/user"

	"express_be/handler/auth"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(
	userHandler userHandler.Handler,
	authhandler auth.Handler,
	priceHandler priceHandler.Handler,
	orderHandler orderHandler.Handler,
	jwtSecret string,
) []route.GroupRoute {
	return []route.GroupRoute{
		// Nhóm công khai (Không cần xác thực)
		{
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
					Path:    "/prices",
					Method:  method.GET,
					Handler: priceHandler.HandleRead,
				},
			},
		},

		// Nhóm cần xác thực (AuthMiddleware)
		{
			Middlewares: []echo.MiddlewareFunc{
				jwt.AuthMiddleware(jwtSecret),
			},
			Routes: []route.Route{
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
				{
					Path:    "/user-info",
					Method:  method.GET,
					Handler: userHandler.HandleGetInfoUser,
				},
				{
					Path:    "/users/:id",
					Method:  method.PUT,
					Handler: userHandler.HandleUpdateInfo,
				},
				{
					Path:    "/search",
					Method:  method.GET,
					Handler: userHandler.HandleSearch,
				},
			},
		},
		{
			Prefix: "/admin/services/prices",
			Middlewares: []echo.MiddlewareFunc{
				jwt.AuthMiddleware(jwtSecret),
				jwt.RoleMiddleware("admin"),
			},
			Routes: []route.Route{
				{
					Path:    "",
					Method:  method.POST,
					Handler: priceHandler.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.PUT,
					Handler: priceHandler.HandleUpdate,
				},
				{
					Path:    "/:id",
					Method:  method.DELETE,
					Handler: priceHandler.HandleDelete,
				},
			},
		},
		{
			Prefix: "/orders",
			Middlewares: []echo.MiddlewareFunc{
				jwt.AuthMiddleware(jwtSecret),
			},
			Routes: []route.Route{
				{
					Path:    "",
					Method:  method.POST,
					Handler: orderHandler.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.PATCH,
					Handler: orderHandler.HandleUpdateOrderStatus,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: orderHandler.HandlerOrderDetail,
				},
				{
					Path:    "",
					Method:  method.GET,
					Handler: orderHandler.HandleListOrder,
				},
			},
		},
	}
}
