// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Message struct {
	ID         string         `json:"id"`
	Content    string         `json:"content"`
	Recipient  string         `json:"recipient"`
	Status     string         `json:"status"`
	CreatedBy  string         `json:"created_by"`
	ApprovedBy sql.NullString `json:"approved_by"`
	RejectedBy sql.NullString `json:"rejected_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
