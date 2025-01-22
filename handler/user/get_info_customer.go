package user

import (
	handlerError "express_be/core/err"

	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Get user information
// @Description  Fetch the details of a user based on their ID. The user can be either a customer or a delivery person.
// @Tags         User-information
// @Accept       json
// @Produce      json

// @Param id path string true "UserID"
// @Security BearerAuth
// @Router       /users/{id} [get]
func (h *handlerImpl) HandleGetInfoUser(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck == "" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	customer, deliveryPerson, usecaseErr := h.userUsecase.GetInfoUser(c.Request().Context(), &id)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}
	if customer != nil && deliveryPerson == nil {
		return response.OK(c, http.StatusOK, "success", customer)
	} else {
		return response.OK(c, http.StatusOK, "success", deliveryPerson)
	}
}
