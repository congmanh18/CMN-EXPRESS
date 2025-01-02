package customer

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleGetInfoCustomer retrieves customer information
// @Summary Nhận thông tin khách hàng
// @Description Truy xuất thông tin chi tiết về một khách hàng cụ thể theo ID
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param   id    path      string  true  "Customer ID"
// @Router /customers/{id} [get]
func (h *handlerImpl) HandleGetInfoCustomer(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	customer, usecaseErr := h.customerUsecase.GetInfoCustomer(c.Request().Context(), &id)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}

	resp := mapper.CustomerToRes(customer)
	return response.OK(c, http.StatusOK, "Customer information", resp)
}
