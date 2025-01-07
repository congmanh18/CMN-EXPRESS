package auth

import (
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handle Change Password
// @Summary Đổi mật khẩu
// @Description Đổi mật khẩu người dùng bằng số điện thoại
// @Description Example reset password payload:
//
//	@Description ``` {
//	@Description	"confirm_password": "strongpassword123",
//	@Description	"new_password": "strongpassword321",
//	@Description	"phone": "0977683511",
//	@Description	"role": "customer"
//	@Description } ```
//
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body model.ResetPasswordRequest true "Reset Password Request"
// @Router /reset-password [patch]
func (h *handlerImpl) HandleChangePassword(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid request")
	}

	// 2. So sánh password và confirm password
	if req.NewPassword == nil || req.ConfirmPassword == nil {
		return response.Error(c, http.StatusBadRequest, "password or confirm password cannot be nil")
	}

	if *req.NewPassword != *req.ConfirmPassword {
		return response.Error(c, http.StatusBadRequest, "passwords do not match")
	}

	// 3. Xử lý đổi mật khẩu
	err := h.authUsecase.ChangePassword(c.Request().Context(), req.Phone, req.NewPassword)
	if err != nil {
		return response.Error(c, err.Code, "Failed to change password: "+err.Message)
	}

	// 4. Trả về kết quả
	return response.OK(c, http.StatusOK, "changed password successfully", nil)
}
