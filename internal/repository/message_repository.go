package repository

import (
	"context"
	"database/sql"
	"maker-checker/internal/db"	
	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(ctx context.Context, content, recipient, createdBy string) (string, error)
	GetByID(ctx context.Context, id string) (*db.Message, error)
	Update(ctx context.Context, msg *db.Message) error
	FindByStatus(ctx context.Context, status string) ([]db.Message, error)
}

type messageRepository struct {
	q *db.Queries
}

func NewMessageRepository(q *db.Queries) MessageRepository {
	return &messageRepository{q: q}
}

func (r *messageRepository) Create(ctx context.Context, content, recipient, createdBy string) (string, error) {
	id := uuid.New().String()
	err := r.q.CreateMessage(ctx, db.CreateMessageParams{
		ID:        id,
		Content:   content,
		Recipient: recipient,
		CreatedBy: createdBy,
	})
	return id, err
}

func (r *messageRepository) GetByID(ctx context.Context, id string) (*db.Message, error) {
	msg, err := r.q.GetMessageByID(ctx, id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (r *messageRepository) Update(ctx context.Context, msg *db.Message) error {
	return r.q.UpdateMessage(ctx, db.UpdateMessageParams{
		Content:    msg.Content,
		Recipient:  msg.Recipient,
		Status:     msg.Status,
		ApprovedBy: msg.ApprovedBy,
		RejectedBy: msg.RejectedBy,
		ID:         msg.ID,
	})
}

func (r *messageRepository) FindByStatus(ctx context.Context, status string) ([]db.Message, error) {
	return r.q.FindMessagesByStatus(ctx, status)
}
