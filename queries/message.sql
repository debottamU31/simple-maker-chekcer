-- name: CreateMessage :exec
INSERT INTO messages (id, content, recipient, status, created_by)
VALUES (?, ?, ?, 'PENDING', ?);

-- name: GetMessageByID :one
SELECT id, content, recipient, status, created_by, approved_by, rejected_by, created_at, updated_at
FROM messages
WHERE id = ?;

-- name: UpdateMessage :exec
UPDATE messages
SET content = ?,
    recipient = ?,
    status = ?,
    approved_by = ?,
    rejected_by = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: FindMessagesByStatus :many
SELECT id, content, recipient, status, created_by, approved_by, rejected_by, created_at, updated_at
FROM messages
WHERE status = ?;