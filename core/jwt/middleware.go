package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Lấy Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing or invalid token"})
			}

			// Tách token từ header
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse token
			claims, err := ParseToken(tokenStr, jwtSecret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
			}

			// Gắn thông tin user vào context
			c.Set("user_id", claims.ID)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}

func ParseToken(tokenStr string, jwtSecret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra phương thức ký
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Kiểm tra và trích xuất claims
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenMalformed
}

func RoleMiddleware(requiredRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]string{"message": "role not found in context"})
			}

			for _, r := range requiredRoles {
				if role == r {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"message": "access forbidden"})
		}
	}
}
