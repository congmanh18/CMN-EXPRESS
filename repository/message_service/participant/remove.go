package participant

import (
	"context"
	"express_be/entity"
)

// RemoveParticipant implements Repo.
func (p *participantImpl) RemoveParticipant(ctx context.Context, conversationID *string, userID *string) error {
	return p.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		Delete(&entity.Participant{}).
		Error
}
