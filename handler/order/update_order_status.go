package order

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleUpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order_status query string true "New status of the order (picking_up, picked_up, pickup_canceled, return_to_hub, at_hub, out_for_delivery, delivered, returned, delivery_canceled)"
// @Security BearerAuth
// @Router /orders/{id} [patch]
func (h *handlerImpl) HandleUpdateOrderStatus(c echo.Context) error {
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

	status := c.QueryParam("order_status")
	if status == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	usecaseErr := h.orderUsecase.UpdateOrderStatus(c.Request().Context(), &id, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}
