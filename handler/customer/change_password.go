package customer

import (
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handle Change Password
// @Summary Đổi mật khẩu
// @Description Đổi mật khẩu khách hàng bằng số điện thoại
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param request body model.ResetPasswordRequest true "Reset Password Request"
// @Router /customers/reset-password [patch]
func (h *handlerImpl) HandleChangePassword(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid request")
	}

	// 2. So sánh password và confirm password
	if req.NewPassword != req.ConfirmPassword {
		return response.Error(c, http.StatusBadRequest, "passwords do not match")
	}

	// 3. Xử lý đổi mật khẩu
	err := h.customerUsecase.ChangePassword(c.Request().Context(), &req.Phone, &req.NewPassword)
	if err != nil {
		return response.Error(c, err.Code, "Failed to change password: "+err.Message)
	}

	// 4. Trả về kết quả
	return response.OK(c, http.StatusOK, "changed password successfully", nil)
}
