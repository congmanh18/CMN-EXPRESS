package message

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleGetParticipantsByConversation implements Handler.
// @Summary Get participants by conversation ID
// @Description Retrieve all participants in a specific conversation
// @Tags Conversations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param id path string true "Conversation ID"
// @Security ApiKeyAuth
// @Router /conversations/{id}/participants [get]
func (h *handlerImpl) HandleGetParticipantsByConversation(c echo.Context) error {
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

	participantID := c.Param("id")
	if participantID == "" {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}
	resp, usecaseErr := h.chatUsecase.GetParticipantsByConversation(c.Request().Context(), &participantID)
	if usecaseErr != nil {
		return response.Error(c, usecaseErr.Code, usecaseErr.Message)
	}

	return response.OK(c, http.StatusOK, "success", resp)
}
