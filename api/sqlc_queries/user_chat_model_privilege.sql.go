// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user_chat_model_privilege.sql

package sqlc_queries

import (
	"context"
)

const createUserChatModelPrivilege = `-- name: CreateUserChatModelPrivilege :one
INSERT INTO user_chat_model_privilege (user_id, chat_model_id, rate_limit, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by
`

type CreateUserChatModelPrivilegeParams struct {
	UserID      int32
	ChatModelID int32
	RateLimit   int32
	CreatedBy   int32
	UpdatedBy   int32
}

func (q *Queries) CreateUserChatModelPrivilege(ctx context.Context, arg CreateUserChatModelPrivilegeParams) (UserChatModelPrivilege, error) {
	row := q.db.QueryRowContext(ctx, createUserChatModelPrivilege,
		arg.UserID,
		arg.ChatModelID,
		arg.RateLimit,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i UserChatModelPrivilege
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatModelID,
		&i.RateLimit,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteUserChatModelPrivilege = `-- name: DeleteUserChatModelPrivilege :exec
DELETE FROM user_chat_model_privilege WHERE id = $1
`

func (q *Queries) DeleteUserChatModelPrivilege(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserChatModelPrivilege, id)
	return err
}

const listUserChatModelPrivileges = `-- name: ListUserChatModelPrivileges :many
SELECT id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by FROM user_chat_model_privilege ORDER BY id
`

func (q *Queries) ListUserChatModelPrivileges(ctx context.Context) ([]UserChatModelPrivilege, error) {
	rows, err := q.db.QueryContext(ctx, listUserChatModelPrivileges)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserChatModelPrivilege
	for rows.Next() {
		var i UserChatModelPrivilege
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ChatModelID,
			&i.RateLimit,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
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

const listUserChatModelPrivilegesByUserID = `-- name: ListUserChatModelPrivilegesByUserID :many
SELECT id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by FROM user_chat_model_privilege 
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) ListUserChatModelPrivilegesByUserID(ctx context.Context, userID int32) ([]UserChatModelPrivilege, error) {
	rows, err := q.db.QueryContext(ctx, listUserChatModelPrivilegesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserChatModelPrivilege
	for rows.Next() {
		var i UserChatModelPrivilege
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ChatModelID,
			&i.RateLimit,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
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

const rateLimiteByUserAndSessionUUID = `-- name: RateLimiteByUserAndSessionUUID :one
SELECT ucmp.rate_limit, cm.name AS chat_model_name
FROM user_chat_model_privilege ucmp
JOIN chat_session cs ON cs.user_id = ucmp.user_id
JOIN chat_model cm ON (cm.id = ucmp.chat_model_id AND cs.model = cm.name and cm.enable_per_mode_ratelimit = true)
WHERE cs.uuid = $1
  AND ucmp.user_id = $2
`

type RateLimiteByUserAndSessionUUIDParams struct {
	Uuid   string
	UserID int32
}

type RateLimiteByUserAndSessionUUIDRow struct {
	RateLimit     int32
	ChatModelName string
}

func (q *Queries) RateLimiteByUserAndSessionUUID(ctx context.Context, arg RateLimiteByUserAndSessionUUIDParams) (RateLimiteByUserAndSessionUUIDRow, error) {
	row := q.db.QueryRowContext(ctx, rateLimiteByUserAndSessionUUID, arg.Uuid, arg.UserID)
	var i RateLimiteByUserAndSessionUUIDRow
	err := row.Scan(&i.RateLimit, &i.ChatModelName)
	return i, err
}

const updateUserChatModelPrivilege = `-- name: UpdateUserChatModelPrivilege :one
UPDATE user_chat_model_privilege SET rate_limit = $2, updated_at = now(), updated_by = $3
WHERE id = $1
RETURNING id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by
`

type UpdateUserChatModelPrivilegeParams struct {
	ID        int32
	RateLimit int32
	UpdatedBy int32
}

func (q *Queries) UpdateUserChatModelPrivilege(ctx context.Context, arg UpdateUserChatModelPrivilegeParams) (UserChatModelPrivilege, error) {
	row := q.db.QueryRowContext(ctx, updateUserChatModelPrivilege, arg.ID, arg.RateLimit, arg.UpdatedBy)
	var i UserChatModelPrivilege
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatModelID,
		&i.RateLimit,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const userChatModelPrivilegeByID = `-- name: UserChatModelPrivilegeByID :one
SELECT id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by FROM user_chat_model_privilege WHERE id = $1
`

func (q *Queries) UserChatModelPrivilegeByID(ctx context.Context, id int32) (UserChatModelPrivilege, error) {
	row := q.db.QueryRowContext(ctx, userChatModelPrivilegeByID, id)
	var i UserChatModelPrivilege
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatModelID,
		&i.RateLimit,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const userChatModelPrivilegeByUserAndModelID = `-- name: UserChatModelPrivilegeByUserAndModelID :one
SELECT id, user_id, chat_model_id, rate_limit, created_at, updated_at, created_by, updated_by FROM user_chat_model_privilege WHERE user_id = $1 AND chat_model_id = $2
`

type UserChatModelPrivilegeByUserAndModelIDParams struct {
	UserID      int32
	ChatModelID int32
}

func (q *Queries) UserChatModelPrivilegeByUserAndModelID(ctx context.Context, arg UserChatModelPrivilegeByUserAndModelIDParams) (UserChatModelPrivilege, error) {
	row := q.db.QueryRowContext(ctx, userChatModelPrivilegeByUserAndModelID, arg.UserID, arg.ChatModelID)
	var i UserChatModelPrivilege
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ChatModelID,
		&i.RateLimit,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
