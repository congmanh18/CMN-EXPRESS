package participant

import (
	"context"
	"errors"
	"express_be/entity"

	"gorm.io/gorm"
)

// GetParticipantsByConversationID implements Repo.
func (p *participantImpl) GetParticipantsByConversationID(ctx context.Context, conversationID *string) ([]entity.Participant, error) {
	if conversationID == nil {
		return nil, errors.New("conversationID cannot be nil")
	}

	var participants []entity.Participant
	err := p.DB.Executor.
		WithContext(ctx).
		Debug().
		Where("conversation_id = ?", conversationID).
		Find(&participants).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Trả về danh sách rỗng nếu không có bản ghi
		}
		return nil, err
	}

	return participants, nil
}
