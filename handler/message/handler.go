package message

import (
	"express_be/usecase/chat"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
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

	HandleSendMessage(s socketio.Conn, data any)
	HandleFetchMessages(s socketio.Conn, data any)
}

type handlerImpl struct {
	chatUsecase chat.ChatUsecase
	socket      *socketio.Server
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
