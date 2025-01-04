package middleware

import (
	"context"
	"fmt"
	"net/http"

	"express_be/core/transport/http/response"

	"github.com/labstack/echo/v4"
)

// Define một interface chung cho các use case tìm kiếm
type FindByIDUseCase[T any] interface {
	Execute(ctx context.Context, id string) (T, error)
}

// Generic middleware tìm kiếm theo ID
func FindEntityByID[T any](useCase FindByIDUseCase[T], paramName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Lấy ID từ param
			entityID := c.Param(paramName)
			if entityID == "" {
				return response.Error(c, http.StatusBadRequest, fmt.Sprintf("%s ID is required", paramName))
			}

			// Tìm entity bằng use case
			entity, err := useCase.Execute(c.Request().Context(), entityID)
			if err != nil {
				return response.Error(c, http.StatusNotFound, fmt.Sprintf("%s not found", paramName))
			}

			// Gắn entity vào context để handler sau dùng
			c.Set(paramName, entity)

			return next(c) // Chuyển sang handler tiếp theo
		}
	}
}
