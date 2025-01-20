package participant

import (
	"context"
	"express_be/entity"
)

// AddParticipant implements Repo.
func (p *participantImpl) AddParticipant(ctx context.Context, participant *entity.Participant) error {
	return p.DB.Executor.
		WithContext(ctx).
		Debug().
		Create(participant).
		Error
}
