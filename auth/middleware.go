package auth

import (
	"context"
	"net/http"

	"express_be/core/transport/http/response"

	"github.com/labstack/echo/v4"
)

func FirebaseAuthMiddleware(authService AuthService, requiredType string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return response.Error(c, http.StatusUnauthorized, "Missing Authorization header")
			}

			token := ""
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				token = authHeader[7:]
			} else {
				return response.Error(c, http.StatusUnauthorized, "Invalid Authorization header format")
			}

			ctx := context.Background()
			firebaseToken, err := authService.VerifyToken(ctx, token)
			if err != nil {
				return response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			}

			// Xác thực account type
			if requiredType != "" {
				err := authService.CheckAccountType(firebaseToken, requiredType)
				if err != nil {
					return response.Error(c, http.StatusForbidden, err.Error())
				}
			}

			// Lưu thông tin user vào context
			c.Set("user", firebaseToken)
			return next(c)
		}
	}
}

// Cách sử dụng ở Routes
// e.GET("/view", func(c echo.Context) error {
// 	user := c.Get("user").(*auth.Token)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Welcome Viewer",
// 		"uid":     user.UID,
// 	})
// }, FirebaseAuthMiddleware(authService, string(ViewOnly)))
