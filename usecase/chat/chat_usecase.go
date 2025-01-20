package chat

import (
	"context"
	error "express_be/core/err"
	"express_be/entity"
	"express_be/model/req"
	"express_be/model/res"

	"express_be/repository/message_service/conversation"
	"express_be/repository/message_service/message"
	"express_be/repository/message_service/participant"
)

type ChatUsecase interface {
	// Conversation-related methods
	CreateConversation(ctx context.Context, req req.NewConversationReq, senderID *string) *error.Err
	GetAllConversations(ctx context.Context, userID *string, page, pageSize *int) ([]res.ConversationRes, *error.Err)
	DeleteConversation(ctx context.Context, id *string) *error.Err

	// Message-related methods
	SendMessage(ctx context.Context, message *entity.Message) *error.Err
	GetMessagesByConversation(ctx context.Context, conversationID *string, page, pageSize *int) ([]entity.Message, *error.Err)
	DeleteMessage(ctx context.Context, messageID *string) *error.Err

	// Participant-related methods
	AddParticipant(ctx context.Context, participant *entity.Participant) *error.Err
	RemoveParticipant(ctx context.Context, conversationID, userID *string) *error.Err
	GetParticipantsByConversation(ctx context.Context, conversationID *string) ([]string, *error.Err)
}

type chatUseCaseImpl struct {
	conversationRepo conversation.Repo
	messageRepo      message.Repo
	participantRepo  participant.Repo
}

// RemoveParticipant implements ChatUseCase.
func (c *chatUseCaseImpl) RemoveParticipant(ctx context.Context, conversationID *string, userID *string) *error.Err {
	err := c.participantRepo.RemoveParticipant(ctx, conversationID, userID)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}

// AddParticipant implements ChatUseCase.
func (c *chatUseCaseImpl) AddParticipant(ctx context.Context, participant *entity.Participant) *error.Err {
	err := c.participantRepo.AddParticipant(ctx, participant)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}

// DeleteMessage implements ChatUseCase.
func (c *chatUseCaseImpl) DeleteMessage(ctx context.Context, messageID *string) *error.Err {
	err := c.messageRepo.DeleteMessage(ctx, messageID)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}

// GetMessagesByConversation implements ChatUseCase.
func (c *chatUseCaseImpl) GetMessagesByConversation(ctx context.Context, conversationID *string, page *int, pageSize *int) ([]entity.Message, *error.Err) {
	messages, err := c.messageRepo.GetMessagesByConversationID(ctx, conversationID, page, pageSize)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	return messages, nil
}

// GetParticipantsByConversation implements ChatUseCase.
func (c *chatUseCaseImpl) GetParticipantsByConversation(ctx context.Context, conversationID *string) ([]string, *error.Err) {
	participants, err := c.participantRepo.GetParticipantsByConversationID(ctx, conversationID)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	var userIDs []string
	for _, participant := range participants {
		userIDs = append(userIDs, *participant.UserID)
	}
	return userIDs, nil
}

// SendMessage implements ChatUseCase.
func (c *chatUseCaseImpl) SendMessage(ctx context.Context, message *entity.Message) *error.Err {
	err := c.messageRepo.SaveMessage(ctx, message)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}

func NewChatUsecase(
	conversationRepo conversation.Repo,
	messageRepo message.Repo,
	participantRepo participant.Repo,
) ChatUsecase {
	return &chatUseCaseImpl{
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
		participantRepo:  participantRepo,
	}
}
