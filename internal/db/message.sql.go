// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: message.sql

package db

import (
	"context"
	"database/sql"
)

const createMessage = `-- name: CreateMessage :exec
INSERT INTO messages (id, content, recipient, status, created_by)
VALUES (?, ?, ?, 'PENDING', ?)
`

type CreateMessageParams struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Recipient string `json:"recipient"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) error {
	_, err := q.db.ExecContext(ctx, createMessage,
		arg.ID,
		arg.Content,
		arg.Recipient,
		arg.CreatedBy,
	)
	return err
}

const findMessagesByStatus = `-- name: FindMessagesByStatus :many
SELECT id, content, recipient, status, created_by, approved_by, rejected_by, created_at, updated_at
FROM messages
WHERE status = ?
`

func (q *Queries) FindMessagesByStatus(ctx context.Context, status string) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, findMessagesByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Recipient,
			&i.Status,
			&i.CreatedBy,
			&i.ApprovedBy,
			&i.RejectedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessageByID = `-- name: GetMessageByID :one
SELECT id, content, recipient, status, created_by, approved_by, rejected_by, created_at, updated_at
FROM messages
WHERE id = ?
`

func (q *Queries) GetMessageByID(ctx context.Context, id string) (Message, error) {
	row := q.db.QueryRowContext(ctx, getMessageByID, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Recipient,
		&i.Status,
		&i.CreatedBy,
		&i.ApprovedBy,
		&i.RejectedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateMessage = `-- name: UpdateMessage :exec
UPDATE messages
SET content = ?,
    recipient = ?,
    status = ?,
    approved_by = ?,
    rejected_by = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateMessageParams struct {
	Content    string         `json:"content"`
	Recipient  string         `json:"recipient"`
	Status     string         `json:"status"`
	ApprovedBy sql.NullString `json:"approved_by"`
	RejectedBy sql.NullString `json:"rejected_by"`
	ID         string         `json:"id"`
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) error {
	_, err := q.db.ExecContext(ctx, updateMessage,
		arg.Content,
		arg.Recipient,
		arg.Status,
		arg.ApprovedBy,
		arg.RejectedBy,
		arg.ID,
	)
	return err
}
