package res

import (
	"express_be/entity"
	"express_be/model/res"
)

func ConversationToRes(conversation *entity.Conversation) *res.ConversationRes {
	return &res.ConversationRes{
		ConversationID: conversation.ID,
		Name:           conversation.Name,
		UpdatedAt:      conversation.UpdatedAt,
	}
}
