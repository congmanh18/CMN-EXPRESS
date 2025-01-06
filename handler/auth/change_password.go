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
	switch *req.Role {
	case "admin":
		err := h.authUsecase.ChangePasswordAdmin(c.Request().Context(), req.Phone, req.NewPassword)
		if err != nil {
			return response.Error(c, err.Code, "Failed to change password: "+err.Message)
		}
	case "accounting":
		err := h.authUsecase.ChangePasswordAccounting(c.Request().Context(), req.Phone, req.NewPassword)
		if err != nil {
			return response.Error(c, err.Code, "Failed to change password: "+err.Message)
		}
	case "customer":
		err := h.authUsecase.ChangePasswordCustomer(c.Request().Context(), req.Phone, req.NewPassword)
		if err != nil {
			return response.Error(c, err.Code, "Failed to change password: "+err.Message)
		}
	case "deliver_person":
		err := h.authUsecase.ChangePasswordDeliveryPerson(c.Request().Context(), req.Phone, req.NewPassword)
		if err != nil {
			return response.Error(c, err.Code, "Failed to change password: "+err.Message)
		}
	default:
		return response.Error(c, http.StatusUnauthorized, "Invalid role")
	}

	// 4. Trả về kết quả
	return response.OK(c, http.StatusOK, "changed password successfully", nil)
}
