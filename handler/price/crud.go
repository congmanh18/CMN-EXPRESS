package price

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	model "express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary      Create a new price
// @Description  Admin creates a new price entry
// @Tags         Prices
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param        body body model.PriceReq true "Price request payload"
// @Router       /prices [post]
func (h *handlerImpl) HandleCreate(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck != "admin" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	adminID := c.Get("user_id").(string)
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

// @Summary      Get all prices
// @Description  Retrieve all price entries
// @Tags         Prices
// @Accept       json
// @Produce      json
// @Router       /prices [get]
func (h *handlerImpl) HandleRead(c echo.Context) error {
	resp, err := h.priceUsecase.ReadAllPrice(c.Request().Context())
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}
	return response.OK(c, http.StatusOK, "success", resp)
}

// @Summary      Update a price
// @Description  Admin updates an existing price
// @Tags         Prices
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param        id path string true "Price ID"
// @Param        body body model.PriceReq true "Price request payload"
// @Router       /prices/{id} [put]
func (h *handlerImpl) HandleUpdate(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck != "admin" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	adminID := c.Get("user_id").(string)
	if adminID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	priceID := c.Param("id")
	if priceID == "" {
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

// @Summary      Delete a price
// @Description  Admin deletes a price entry
// @Tags         Prices
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param        id path string true "Price ID"
// @Router       /prices/{id} [delete]
func (h *handlerImpl) HandleDelete(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck != "admin" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	adminID := c.Get("user_id").(string)
	if adminID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	priceID := c.Param("id")
	if priceID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	err := h.priceUsecase.DeletePrice(c.Request().Context(), &priceID, &adminID)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)

}
