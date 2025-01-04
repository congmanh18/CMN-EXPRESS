package customer

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"

	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleUpdateCustomer updates customer information
// @Summary Cập nhật thông tin khách hàng
// @Description Cập nhật chi tiết của một khách hàng cụ thể theo ID
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param   id      path      string                 true  "Customer ID"
// @Param   request body      model.UpdateCustomerReq true  "Customer Update Request"
// @Router /customers/{id} [put]
func (h *handlerImpl) HandleUpdateCustomer(c echo.Context) error {
	// 1. Kiểm tra id từ url
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	// 2. Parse dữ liệu đầu vào
	var req model.UpdateCustomerReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 3. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	// 4. mapping req thành entity
	customer := mapper.UpdateToCustomer(req)

	// 5. Thực hiện cập nhật
	if err := h.customerUsecase.UpdateInfoCustomer(c.Request().Context(), customer, &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to update: "+err.Error())
	}

	// 6. Trả về kết quả
	return response.OK(c, http.StatusOK, "Updated successfully", nil)
}
