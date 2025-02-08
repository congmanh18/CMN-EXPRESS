package order

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"
	"strconv"

	model "express_be/model/req"

	"github.com/labstack/echo/v4"
)

// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body model.CreateOrderReq true "Create Order Request"
// @Security BearerAuth
// @Router /orders [post]
func (h *handlerImpl) HandleCreate(c echo.Context) error {
	id, ok := c.Get("user_id").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	var req model.CreateOrderReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	err := h.orderUsecase.CreateOrder(c.Request().Context(), req, &id)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}
	return response.OK(c, http.StatusOK, "success", nil)
}

// @Summary Get the nearest delivery person
// @Description Retrieve the closest delivery person to the provided latitude and longitude
// @Tags Orders
// @Accept json
// @Produce json
// @Param lat query string true "Latitude of the location"
// @Param lon query string true "Longitude of the location"
// @Param status query string false "Filter by delivery person's status (default: off_duty)"
// @Security BearerAuth
// @Router /orders/nearest [get]
func (h *handlerImpl) HandleNearestDeliveryPerson(c echo.Context) error {
	latStr := c.QueryParam("lat")
	lonStr := c.QueryParam("lon")
	if latStr == "" || lonStr == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	status := c.QueryParam("status")
	if status == "" {
		status = "off_duty"
	}

	nearestDeliveryPerson, usecaseErr := h.orderUsecase.FindNearestDeliveryPerson(c.Request().Context(), &lat, &lon, &status)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}
	return response.OK(c, http.StatusOK, "success", nearestDeliveryPerson)
}
