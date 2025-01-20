package chat

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/req"
	"express_be/model/req"
)

// CreateConversation implements ChatUseCase.
func (c *chatUseCaseImpl) CreateConversation(ctx context.Context, req req.NewConversationReq, senderID *string) *error.Err {
	conversation, sender, receiver := mapper.ReqToConversationParticipants(req, senderID)
	err := c.conversationRepo.CreateConversation(ctx, conversation)
	if err != nil {
		return error.ErrInternalServer
	}

	err = c.participantRepo.AddParticipant(ctx, sender)
	if err != nil {
		return error.ErrInternalServer
	}

	err = c.participantRepo.AddParticipant(ctx, receiver)
	if err != nil {
		return error.ErrInternalServer
	}

	return nil
}
