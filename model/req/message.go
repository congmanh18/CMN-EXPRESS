package req

type MessageReq struct {
	Content     *string `json:"content"`
	MessageType *string `json:"message_type"`
}

type NewConversationReq struct {
	Name       *string `json:"name"`
	ReceiverID *string `json:"receiver_id" validate:"required"`
}
