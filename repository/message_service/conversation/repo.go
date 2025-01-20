package conversation

import (
	"context"
	"express_be/core/db/postgresql"
	"express_be/entity"
)

type Repo interface {
	CreateConversation(ctx context.Context, conversation *entity.Conversation) error
	GetConversationByID(ctx context.Context, id *string) (*entity.Conversation, error)
	GetAllConversations(ctx context.Context, userID *string, page, pageSize *int) ([]entity.Conversation, error)
	DeleteConversation(ctx context.Context, id *string) error
}

type conversationImpl struct {
	DB *postgresql.Database
}

func NewRepo(db *postgresql.Database) Repo {
	return &conversationImpl{DB: db}
}
