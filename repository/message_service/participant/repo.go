package participant

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	AddParticipant(ctx context.Context, participant *entity.Participant) error
	RemoveParticipant(ctx context.Context, conversationID *string, userID *string) error
	GetParticipantsByConversationID(ctx context.Context, conversationID *string) ([]entity.Participant, error)
}

type participantImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &participantImpl{DB: db}
}
