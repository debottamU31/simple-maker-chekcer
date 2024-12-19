package service

import (
	"context"
	"database/sql"
	"errors"

	"maker-checker/internal/db"
	"maker-checker/internal/mailer"
	"maker-checker/internal/repository"
)

// Define message status constants
const (
	StatusPending  = "PENDING"
	StatusApproved = "APPROVED"
	StatusRejected = "REJECTED"
	StatusSent     = "SENT"
)

type MessageService interface {
	CreateMessage(ctx context.Context, content, recipient, createdBy string) (*db.Message, error)
	ApproveMessage(ctx context.Context, id, checker string) error
	RejectMessage(ctx context.Context, id, checker string) error
	ListPending(ctx context.Context) ([]db.Message, error)
}

type messageService struct {
	repo   repository.MessageRepository
	mailer mailer.Mailer
}

func NewMessageService(repo repository.MessageRepository, mailer mailer.Mailer) MessageService {
	return &messageService{repo: repo, mailer: mailer}
}

func (s *messageService) CreateMessage(ctx context.Context, content, recipient, createdBy string) (*db.Message, error) {
	id, err := s.repo.Create(ctx, content, recipient, createdBy)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *messageService) ApproveMessage(ctx context.Context, id, checker string) error {
	msg, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if msg == nil {
		return errors.New("message not found")
	}
	if msg.Status != StatusPending {
		return errors.New("message not in pending state")
	}

	msg.Status = StatusApproved
	msg.ApprovedBy = sql.NullString{String: checker, Valid: true}

	if err := s.repo.Update(ctx, msg); err != nil {
		return err
	}

	// Simulate sending
	if err := s.mailer.Send(msg.Recipient, msg.Content); err != nil {
		return err
	}

	msg.Status = StatusSent
	return s.repo.Update(ctx, msg)
}

func (s *messageService) RejectMessage(ctx context.Context, id, checker string) error {
	msg, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if msg == nil {
		return errors.New("message not found")
	}
	if msg.Status != StatusPending {
		return errors.New("message not in pending state")
	}

	msg.Status = StatusRejected
	msg.RejectedBy = sql.NullString{String: checker, Valid: true}
	return s.repo.Update(ctx, msg)
}

func (s *messageService) ListPending(ctx context.Context) ([]db.Message, error) {
	return s.repo.FindByStatus(ctx, StatusPending)
}
