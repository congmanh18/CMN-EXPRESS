package req

import (
	"express_be/core/pointer"
	"express_be/core/record"
	"express_be/entity"
	"express_be/model/req"
	"time"

	"github.com/google/uuid"
)

func ReqToMessage(req req.MessageReq, messageType entity.MessageType, userID, conversationID *string) *entity.Message {
	return &entity.Message{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		ConversationID: conversationID,
		UserID:         userID,
		Content:        req.Content,
		MessageType:    messageType,
	}
}

func ReqToConversationParticipants(req req.NewConversationReq, senderID *string) (*entity.Conversation, *entity.Participant, *entity.Participant) {
	conversationID := uuid.New().String()
	conversation := &entity.Conversation{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(conversationID),
		},
		Name: req.Name,
	}

	sender := &entity.Participant{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		UserID:         senderID,
		JoinedAt:       pointer.Time(time.Now()),
		ConversationID: &conversationID,
	}

	receiver := &entity.Participant{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		UserID:         req.ReceiverID,
		JoinedAt:       pointer.Time(time.Now()),
		ConversationID: &conversationID,
	}

	return conversation, sender, receiver
}
