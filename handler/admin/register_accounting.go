package admin

import (
	"express_be/core/security"
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"express_be/repository/accounting"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// HandleAccouting handles accounting registration
// @Summary Đăng ký kế toán
// @Description Đăng ký kế toán mới cung cấp số tối thiểu "phone" và "password"
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   request  body  model.RegisterRequest  true  "Accounting Registration Request" example({"phone": "0912345678", "password": "abc@1234", "account_type": "prepaid"})
// @Router /admin/create-accouting [post]
func (h *handlerImpl) HandleRegisterAccounting(c echo.Context) error {
	// 1. Parse dữ liệu đầu vào
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request")
	}

	// 2. Validate dữ liệu
	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	id := uuid.New().String()
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Error hashing password")
	}
	accountingEntity := accounting.Accounting{}
	accountingEntity.ID = &id
	accountingEntity.Phone = &req.Phone
	accountingEntity.Password = &hashedPassword

	usecaseErr := h.adminAccountingUsecase.CreateAccounting(c.Request().Context(), &accountingEntity)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "Register accounting successfully", nil)
}
