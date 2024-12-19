// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateMessage(ctx context.Context, arg CreateMessageParams) error
	FindMessagesByStatus(ctx context.Context, status string) ([]Message, error)
	GetMessageByID(ctx context.Context, id string) (Message, error)
	UpdateMessage(ctx context.Context, arg UpdateMessageParams) error
}

var _ Querier = (*Queries)(nil)
