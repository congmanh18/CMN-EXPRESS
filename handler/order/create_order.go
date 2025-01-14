package order

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"

	model "express_be/model/req"

	"github.com/labstack/echo/v4"
)

func (h *handlerImpl) HandleCreate(c echo.Context) error {
	var req model.CreateOrderReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	if err := model.Validate.Struct(req); err != nil {
		return response.Error(c, handlerError.ErrInvalidFormat.Code, handlerError.ErrInvalidFormat.Message)
	}

	err := h.orderUsecase.CreateOrder(c.Request().Context(), req)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}
	return response.OK(c, http.StatusOK, "success", nil)

}
