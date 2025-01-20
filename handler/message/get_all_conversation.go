package message

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// HandleGetAllConversations implements Handler.
// @Summary Get all conversations
// @Description Retrieve all conversations for a user
// @Tags Conversations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Security ApiKeyAuth
// @Router /conversations [get]
func (h *handlerImpl) HandleGetAllConversations(c echo.Context) error {
	roleCheck, ok := c.Get("role").(string)
	if !ok {
		return response.Error(c, handlerError.ErrTokenMissing.Code, handlerError.ErrTokenMissing.Message)
	}

	if roleCheck == "" {
		return response.Error(c, handlerError.ErrAccessDenied.Code, handlerError.ErrAccessDenied.Message)
	}

	userID := c.Get("user_id").(string)
	if userID == "" {
		return response.Error(c, handlerError.ErrInvalidToken.Code, handlerError.ErrInvalidToken.Message)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	resp, uscaseErr := h.chatUsecase.GetAllConversations(c.Request().Context(), &userID, &page, &pageSize)
	if err != nil {
		return response.Error(c, uscaseErr.Code, uscaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "success", resp)
}
