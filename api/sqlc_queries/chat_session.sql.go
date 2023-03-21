// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: chat_session.sql

package sqlc_queries

import (
	"context"
	"time"
)

const createChatSession = `-- name: CreateChatSession :one
INSERT INTO chat_session (user_id, topic, max_length, uuid)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type CreateChatSessionParams struct {
	UserID    int32
	Topic     string
	MaxLength int32
	Uuid      string
}

func (q *Queries) CreateChatSession(ctx context.Context, arg CreateChatSessionParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, createChatSession,
		arg.UserID,
		arg.Topic,
		arg.MaxLength,
		arg.Uuid,
	)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const createChatSessionByUUID = `-- name: CreateChatSessionByUUID :one
INSERT INTO chat_session (user_id, uuid, topic, created_at, active,  max_length)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type CreateChatSessionByUUIDParams struct {
	UserID    int32
	Uuid      string
	Topic     string
	CreatedAt time.Time
	Active    bool
	MaxLength int32
}

func (q *Queries) CreateChatSessionByUUID(ctx context.Context, arg CreateChatSessionByUUIDParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, createChatSessionByUUID,
		arg.UserID,
		arg.Uuid,
		arg.Topic,
		arg.CreatedAt,
		arg.Active,
		arg.MaxLength,
	)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const createOrUpdateChatSessionByUUID = `-- name: CreateOrUpdateChatSessionByUUID :one
INSERT INTO chat_session(uuid, user_id, topic, max_length, temperature, max_tokens, top_p)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (uuid) 
DO UPDATE SET
max_length = EXCLUDED.max_length, 
max_tokens = EXCLUDED.max_tokens,
temperature = EXCLUDED.temperature, 
top_p = EXCLUDED.top_p,
topic = CASE WHEN chat_session.topic IS NULL THEN EXCLUDED.topic ELSE chat_session.topic END,
updated_at = now()
returning id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type CreateOrUpdateChatSessionByUUIDParams struct {
	Uuid        string
	UserID      int32
	Topic       string
	MaxLength   int32
	Temperature float64
	MaxTokens   int32
	TopP        float64
}

func (q *Queries) CreateOrUpdateChatSessionByUUID(ctx context.Context, arg CreateOrUpdateChatSessionByUUIDParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, createOrUpdateChatSessionByUUID,
		arg.Uuid,
		arg.UserID,
		arg.Topic,
		arg.MaxLength,
		arg.Temperature,
		arg.MaxTokens,
		arg.TopP,
	)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const deleteChatSession = `-- name: DeleteChatSession :exec
DELETE FROM chat_session 
WHERE id = $1
`

func (q *Queries) DeleteChatSession(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteChatSession, id)
	return err
}

const deleteChatSessionByUUID = `-- name: DeleteChatSessionByUUID :exec
update chat_session set active = false
WHERE uuid = $1
returning id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

func (q *Queries) DeleteChatSessionByUUID(ctx context.Context, uuid string) error {
	_, err := q.db.ExecContext(ctx, deleteChatSessionByUUID, uuid)
	return err
}

const getAllChatSessions = `-- name: GetAllChatSessions :many
SELECT id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens FROM chat_session 
where active = true
ORDER BY id
`

func (q *Queries) GetAllChatSessions(ctx context.Context) ([]ChatSession, error) {
	rows, err := q.db.QueryContext(ctx, getAllChatSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatSession
	for rows.Next() {
		var i ChatSession
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Uuid,
			&i.Topic,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Active,
			&i.MaxLength,
			&i.Temperature,
			&i.TopP,
			&i.MaxTokens,
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

const getChatSessionByID = `-- name: GetChatSessionByID :one
SELECT id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens FROM chat_session WHERE id = $1
`

func (q *Queries) GetChatSessionByID(ctx context.Context, id int32) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, getChatSessionByID, id)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const getChatSessionByUUID = `-- name: GetChatSessionByUUID :one
SELECT id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens FROM chat_session 
WHERE active = true and uuid = $1
order by updated_at
`

func (q *Queries) GetChatSessionByUUID(ctx context.Context, uuid string) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, getChatSessionByUUID, uuid)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const getChatSessionsByUserID = `-- name: GetChatSessionsByUserID :many
SELECT cs.id, cs.user_id, cs.uuid, cs.topic, cs.created_at, cs.updated_at, cs.active, cs.max_length, cs.temperature, cs.top_p, cs.max_tokens
FROM chat_session cs
WHERE cs.user_id = $1 and cs.active = true
ORDER BY cs.id
`

func (q *Queries) GetChatSessionsByUserID(ctx context.Context, userID int32) ([]ChatSession, error) {
	rows, err := q.db.QueryContext(ctx, getChatSessionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatSession
	for rows.Next() {
		var i ChatSession
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Uuid,
			&i.Topic,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Active,
			&i.MaxLength,
			&i.Temperature,
			&i.TopP,
			&i.MaxTokens,
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

const hasChatSessionPermission = `-- name: HasChatSessionPermission :one
SELECT COUNT(*) > 0 as has_permission
FROM chat_session cs
INNER JOIN auth_user au ON cs.user_id = au.id
WHERE cs.id = $1 AND (cs.user_id = $2 OR au.is_superuser)
`

type HasChatSessionPermissionParams struct {
	ID     int32
	UserID int32
}

func (q *Queries) HasChatSessionPermission(ctx context.Context, arg HasChatSessionPermissionParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, hasChatSessionPermission, arg.ID, arg.UserID)
	var has_permission bool
	err := row.Scan(&has_permission)
	return has_permission, err
}

const updateChatSession = `-- name: UpdateChatSession :one
UPDATE chat_session SET user_id = $2, topic = $3, updated_at = now(), active = $4
WHERE id = $1
RETURNING id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type UpdateChatSessionParams struct {
	ID     int32
	UserID int32
	Topic  string
	Active bool
}

func (q *Queries) UpdateChatSession(ctx context.Context, arg UpdateChatSessionParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, updateChatSession,
		arg.ID,
		arg.UserID,
		arg.Topic,
		arg.Active,
	)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const updateChatSessionByUUID = `-- name: UpdateChatSessionByUUID :one
UPDATE chat_session SET user_id = $2, topic = $3, updated_at = now()
WHERE uuid = $1
RETURNING id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type UpdateChatSessionByUUIDParams struct {
	Uuid   string
	UserID int32
	Topic  string
}

func (q *Queries) UpdateChatSessionByUUID(ctx context.Context, arg UpdateChatSessionByUUIDParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, updateChatSessionByUUID, arg.Uuid, arg.UserID, arg.Topic)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const updateChatSessionTopicByUUID = `-- name: UpdateChatSessionTopicByUUID :one
INSERT INTO chat_session(uuid, user_id, topic)
VALUES ($1, $2, $3)
ON CONFLICT (uuid) 
DO UPDATE SET
topic = EXCLUDED.topic, 
updated_at = now()
returning id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type UpdateChatSessionTopicByUUIDParams struct {
	Uuid   string
	UserID int32
	Topic  string
}

func (q *Queries) UpdateChatSessionTopicByUUID(ctx context.Context, arg UpdateChatSessionTopicByUUIDParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, updateChatSessionTopicByUUID, arg.Uuid, arg.UserID, arg.Topic)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}

const updateSessionMaxLength = `-- name: UpdateSessionMaxLength :one
UPDATE chat_session
SET max_length = $2,
    updated_at = now()
WHERE uuid = $1
RETURNING id, user_id, uuid, topic, created_at, updated_at, active, max_length, temperature, top_p, max_tokens
`

type UpdateSessionMaxLengthParams struct {
	Uuid      string
	MaxLength int32
}

func (q *Queries) UpdateSessionMaxLength(ctx context.Context, arg UpdateSessionMaxLengthParams) (ChatSession, error) {
	row := q.db.QueryRowContext(ctx, updateSessionMaxLength, arg.Uuid, arg.MaxLength)
	var i ChatSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.Topic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.MaxLength,
		&i.Temperature,
		&i.TopP,
		&i.MaxTokens,
	)
	return i, err
}
