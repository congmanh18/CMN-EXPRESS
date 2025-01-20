package message

import (
	"context"
	"encoding/json"
	handlerError "express_be/core/err"
	"express_be/core/transport/http/response"
	"express_be/model/req"
	"net/http"
	"strconv"

	"express_be/entity"
	"express_be/usecase/chat"

	"github.com/labstack/echo/v4"
	socketio "github.com/nkovacs/go-socket.io"
)

type Handler interface {
	// conversations handle
	HandleCreateConversation(c echo.Context) error
	HandleGetAllConversations(c echo.Context) error
	// HandleDeleteConversation(c echo.Context) error

	// messages handle
	// HandleDeleteMessage(c echo.Context) error

	// participants handle
	HandleGetParticipantsByConversation(c echo.Context) error

	HandleSocketEvents()
}

type handlerImpl struct {
	chatUsecase chat.ChatUsecase
	socket      *socketio.Server
}

func (h *handlerImpl) HandleSocketEvents() {
	userMap := make(map[string]string) // Map of socket ID to user

	// On connection
	h.socket.On("connection", func(so socketio.Socket) {
		println("New connection: ", so.Id())

		newMessages := make(chan string)
		s := Subscribe()

		// Handle user joining a conversation
		so.On("join", func(user string) {
			println("User joined: ", user)
			userMap[so.Id()] = user
			Join(user)

			// Send archived events to the connected user
			for _, event := range s.Archive {
				so.Emit("chat", event)
			}
		})

		// Handle user sending a message
		so.On("chat", func(msg string) {
			newMessages <- msg
		})

		// Handle user leaving a conversation
		so.On("leave", func(user string) {
			println("User leaving: ", user)
			Leave(user)
			s.Cancel()
		})

		// Handle adding a participant
		so.On("add_participant", func(participant entity.Participant) {
			err := h.chatUsecase.AddParticipant(context.Background(), &participant)
			if err != nil {
				so.Emit("error", err.Error())
				return
			}
			println("Participant added: ", participant.UserID)
		})

		// Handle removing a participant
		so.On("remove_participant", func(payload struct {
			ConversationID string `json:"conversation_id"`
			UserID         string `json:"user_id"`
		}) {
			err := h.chatUsecase.RemoveParticipant(context.Background(), &payload.ConversationID, &payload.UserID)
			if err != nil {
				so.Emit("error", err.Error())
				return
			}
			println("Participant removed: ", payload.UserID)
		})

		// Handle disconnection
		so.On("disconnection", func() {
			println("User disconnected: ", so.Id())
			user := userMap[so.Id()]
			delete(userMap, so.Id())
			Leave(user)
			s.Cancel()
		})

		go func() {
			for {
				select {
				case event := <-s.New:
					so.Emit("chat", event)
				case msg := <-newMessages:
					var newMSG entity.Message
					json.Unmarshal([]byte(msg), &newMSG)
					println("Message received: ", newMSG.Content)
					h.chatUsecase.SendMessage(context.Background(), &newMSG)
					Say(newMSG)
				}
			}
		}()
	})
}

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

type HandlerInject struct {
	ChatUsecase chat.ChatUsecase
	Socket      *socketio.Server
}

func NewHandler(inject HandlerInject) Handler {
	return &handlerImpl{
		inject.ChatUsecase,
		inject.Socket,
	}
}
