package message

import (
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleCreateConversation implements Handler.
// @Summary Create a new conversation
// @Description Create a new conversation for a user
// @Tags Conversations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token" default(Bearer <access-token>)
// @Param request body req.NewConversationReq true "Create Conversation Request"
// @Security ApiKeyAuth
// @Router /conversations [post]
func (h *handlerImpl) HandleCreateConversation(c echo.Context) error {
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

	var req req.NewConversationReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, handlerError.ErrMissingField.Code, handlerError.ErrMissingField.Message)
	}

	err := h.chatUsecase.CreateConversation(c.Request().Context(), req, &userID)
	if err != nil {
		return response.Error(c, err.Code, err.Message)
	}

	return response.OK(c, http.StatusOK, "success", nil)
}
