package order

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	model "express_be/model/res"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleListOrder godoc
// @Summary List orders
// @Description Get a list of orders with pagination
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Security BearerAuth
// @Router /orders [get]
func (h *handlerImpl) HandleListOrder(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}
	id, ok := c.Get("user_id").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}
	if roleCheck == "" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	orders, err := h.orderUsecase.GetOrderList(c.Request().Context(), &id, &roleCheck, &page, &pageSize)

	resp := model.PaginationResponse{
		Page:     page,
		PageSize: pageSize,
		Orders:   orders,
	}

	return response.OK(c, http.StatusOK, "success", resp)
}
