package price

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlerImpl) HandleCreate(c echo.Context) error {
	adminID := c.Param("id")
	if adminID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	var req model.PriceReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.priceUsecase.CreateNewPrice(c.Request().Context(), req, &adminID)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}

// HandleRead implements Handler.
func (h *handlerImpl) HandleRead(c echo.Context) error {
	resp, err := h.priceUsecase.ReadAllPrice(c.Request().Context())
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}
	return response.OK(c, http.StatusOK, "success", resp)
}

// HandleUpdate implements Handler.
func (h *handlerImpl) HandleUpdate(c echo.Context) error {
	priceID := c.Param("id")
	if priceID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}
	adminID := c.Param("id")
	if adminID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	var req model.PriceReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.priceUsecase.UpdatePrice(c.Request().Context(), &priceID, req, &adminID)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}

// HandleDelete implements Handler.
func (h *handlerImpl) HandleDelete(c echo.Context) error {
	priceID := c.Param("id")
	if priceID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	adminID := c.Param("id")
	if adminID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	err := h.priceUsecase.DeletePrice(c.Request().Context(), &priceID, &adminID)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)

}
