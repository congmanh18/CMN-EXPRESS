package delivery

import (
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handle Change Password
// @Summary Đổi mật khẩu
// @Description Đổi mật khẩu khách hàng bằng số điện thoại
// @Tags DeliveryPersons
// @Accept  json
// @Produce  json
// @Param request body model.ResetPasswordRequest true "Reset Password Request"
// @Router /delivery-persons/reset-password [patch]
func (h *handlerImpl) HandleChangePassword(c echo.Context) error {
	var req model.ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid request")
	}

	// 2. So sánh password và confirm password
	if req.NewPassword != req.ConfirmPassword {
		return response.Error(c, http.StatusBadRequest, "passwords do not match")
	}

	// 3. Xử lý đổi mật khẩu
	err := h.deliveryPersonUsecase.ChangePassword(c.Request().Context(), &req.Phone, &req.NewPassword)
	if err != nil {
		return response.Error(c, err.Code, "Failed to change password: "+err.Message)
	}

	return response.OK(c, http.StatusOK, "changed password successfully", nil)

}
