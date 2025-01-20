package message

import (
	"context"
	"express_be/core/pointer"
	"express_be/entity"

	"github.com/google/uuid"
	socketio "github.com/googollee/go-socket.io"
	"github.com/mitchellh/mapstructure"
)

func (h *handlerImpl) HandleSendMessage(s socketio.Conn, data any) {
	ctx := s.Context().(map[string]string)
	userID := ctx["user_id"] // Lấy thông tin user từ context

	var message entity.Message
	if err := mapstructure.Decode(data, &message); err != nil {
		s.Emit("error", "Invalid message format")
		return
	}

	// Gắn thêm thông tin sender vào tin nhắn
	message.UserID = &userID
	message.ID = pointer.String(uuid.New().String())
	// Lưu tin nhắn vào database
	err := h.chatUsecase.SendMessage(context.Background(), &message)
	if err != nil {
		s.Emit("error", "Failed to send message: "+err.Message)
		return
	}

	// Phát tin nhắn đến các người dùng khác trong cuộc trò chuyện
	h.socket.BroadcastToRoom("/", *message.ConversationID, "new_message", message)
	s.Emit("message_ack", "Message sent successfully")
}

func (h *handlerImpl) HandleFetchMessages(s socketio.Conn, data any) {
	ctx := s.Context().(map[string]string)
	userID := ctx["user_id"] // Use userID in the logic below if needed
	if userID == "" {
		s.Emit("error", "User ID is required")
		return
	}

	var req struct {
		ConversationID string `json:"conversation_id"`
		Page           int    `json:"page"`
		PageSize       int    `json:"page_size"`
	}
	if err := mapstructure.Decode(data, &req); err != nil {
		s.Emit("error", "Invalid request format")
		return
	}

	// Kiểm tra giá trị mặc định cho page và page_size
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}

	// Lấy danh sách tin nhắn
	messages, err := h.chatUsecase.GetMessagesByConversation(
		context.Background(),
		&req.ConversationID,
		&req.Page,
		&req.PageSize,
	)
	if err != nil {
		s.Emit("error", "Failed to fetch messages: "+err.Message)
		return
	}

	// Trả danh sách tin nhắn về client
	s.Emit("messages_list", messages)
}
