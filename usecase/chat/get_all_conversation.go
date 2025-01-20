package chat

import (
	"context"
	error "express_be/core/err"
	mapper "express_be/mapper/res"

	"express_be/model/res"
)

// GetAllConversations implements ChatUseCase.
func (c *chatUseCaseImpl) GetAllConversations(ctx context.Context, userID *string, page, pageSize *int) ([]res.ConversationRes, *error.Err) {
	conversations, err := c.conversationRepo.GetAllConversations(ctx, userID, page, pageSize)
	if err != nil {
		return nil, error.ErrInternalServer
	}
	var conversationResponses []res.ConversationRes
	for _, conversation := range conversations {
		conversationResponse := mapper.ConversationToRes(&conversation)
		conversationResponses = append(conversationResponses, *conversationResponse)
	}

	return conversationResponses, nil
}
