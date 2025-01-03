package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Authorization header"})
			}

			// Bỏ "Bearer " nếu có
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid Authorization header"})
			}

			// Parse token
			claims := &JWTClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
			}

			// Gán thông tin user vào context
			c.Set("user", claims)
			return next(c)
		}
	}
}

// adminGroup := e.Group("/admin")
// adminGroup.Use(FirebaseAuthMiddleware(authService, "admin"))
// adminGroup.GET("/dashboard", handler.AdminDashboard)
