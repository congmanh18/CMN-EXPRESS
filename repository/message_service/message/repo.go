package message

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	SaveMessage(ctx context.Context, message *entity.Message) error
	GetMessagesByConversationID(ctx context.Context, conversationID *string, page, pageSize *int) ([]entity.Message, error)
	DeleteMessage(ctx context.Context, messageID *string) error
}

type messageImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &messageImpl{DB: db}
}
