package user

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleVerifyUser implements Handler.
// @Summary Update customer status
// @Description Update customer status by ID
// @Tags User-information
// @Accept json
// @Produce json

// @Param id path string true "UserID"
// @Param approval_status query string true "Trạng thái mới của khách hàng (accept, deny)"
// @Security BearerAuth
// @Router /users/{id} [patch]
func (h *handlerImpl) HandleUpdateUserStatus(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck != "admin" && roleCheck != "accounting" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	status := c.QueryParam("approval_status")
	if status == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	usecaseErr := h.userUsecase.UpdateStatus(c.Request().Context(), &id, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}
