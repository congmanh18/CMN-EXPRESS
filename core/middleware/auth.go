package middleware

import (
	"net/http"
	"strings"

	"express_be/core/security"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := security.ParseTokenFromString(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
		}

		c.Set("account_type", claims.AccountType)
		c.Set("user_id", claims.ID)
		return next(c)
	}
}

// UserRole
func AccountTypeMiddleware(requiredRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, ok := c.Get("account_type").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]string{"message": "role not found in context"})
			}

			// Kiểm tra nếu role của user nằm trong danh sách requiredRoles
			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"message": "access forbidden"})
		}
	}
}
