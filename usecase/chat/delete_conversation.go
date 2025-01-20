package chat

import (
	"context"
	error "express_be/core/err"
)

// DeleteConversation implements ChatUseCase.
func (c *chatUseCaseImpl) DeleteConversation(ctx context.Context, id *string) *error.Err {
	err := c.conversationRepo.DeleteConversation(ctx, id)
	if err != nil {
		return error.ErrInternalServer
	}
	return nil
}
