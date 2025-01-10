package auth

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	model "express_be/model/req"

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
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	if *req.NewPassword != *req.ConfirmPassword {
		return response.Error(c, handlerError.ErrWrongPasswordMatchConfirm.Code, handlerError.ErrWrongPasswordMatchConfirm.Description)
	}

	// 3. Xử lý đổi mật khẩu
	err := h.authUsecase.ChangePassword(c.Request().Context(), req.Phone, req.NewPassword)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	// 4. Trả về kết quả
	return response.OK(c, 200, "success", nil)
}
