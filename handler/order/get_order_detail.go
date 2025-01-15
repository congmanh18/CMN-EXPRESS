package order

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandlerOrderDetail godoc
// @Summary Get order detail
// @Description Retrieve detailed information about an order based on its ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param id path string true "Order ID"
// @Security ApiKeyAuth
// @Router /orders/{id} [get]
func (h *handlerImpl) HandlerOrderDetail(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck == "" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	order, err := h.orderUsecase.GetOrderDetail(c.Request().Context(), &id)
	if err != nil {
		return response.Error(c, handlerError.ErrInternalServer.Code, handlerError.ErrInternalServer.Message)
	}

	return response.OK(c, http.StatusOK, "success", order)

}
