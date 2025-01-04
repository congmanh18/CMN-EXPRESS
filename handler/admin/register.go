package admin

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/req"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleRegistration handles admin registration
// @Summary Đăng ký quản trị viên
// @Description Đăng ký quản trị viên mới cung cấp tối thiểu "phone" và "password"
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   request  body  model.BaseRegisterRequest  true  "Register successfully"
// @Router /admin/register [post]
func (h *handlerImpl) HandleRegister(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.BaseRegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Xử lý đăng ký
	admin := mapper.ReqToAdmin(req)

	// 3. Gọi đến usecase để xử lý logic
	err := h.adminUsecase.CreateAdmin(c.Request().Context(), admin)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Internal server error")
	}

	return response.OK(c, http.StatusOK, "Register successfully", nil)
}
