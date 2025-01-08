package dashboard

import (
	"express_be/core/transport/http/response"
	mapper "express_be/mapper/res"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Get user information
// @Description  Fetch the details of a user based on their ID. The user can be either a customer or a delivery person.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Router       /users/{id} [get]
func (h *handlerImpl) HandleGetInfoUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "id is required")
	}

	customer, deliveryPerson, usecaseErr := h.userUsecase.GetInfoUser(c.Request().Context(), &id)
	if usecaseErr != nil {
		return response.Error(
			c,
			usecaseErr.Code,
			usecaseErr.Message,
		)
	}
	if customer != nil {
		resp := mapper.CustomerToRes(customer)
		return response.OK(c, http.StatusOK, "User information", resp)
	} else {
		resp := mapper.DeliveryPersonToRes(deliveryPerson)
		return response.OK(c, http.StatusOK, "User information", resp)
	}
}
