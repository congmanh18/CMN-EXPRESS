package customer

import (
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleDeleteCustomer deletes a customer (soft delete)
// @Summary Xóa tài khoản khách hàng
// @Description Thực hiện xóa khách hàng theo ID
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param   id    path      string  true  "Customer ID"
// @Router /customers/{id} [delete]
func (h *handlerImpl) HandleDeleteCustomer(c echo.Context) error {
	// 1. Kiểm tra id từ url
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	// 2. Xóa khách hàng từ cơ sở dữ liệu
	if err := h.customerUsecase.SoftDeleteCustomer(c.Request().Context(), &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to delete: "+err.Error())
	}

	// 3. Trả về kết quả thành công
	return response.OK(c, http.StatusOK, "Deleted successfully", nil)
}
