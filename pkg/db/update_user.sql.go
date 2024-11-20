// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: update_user.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const updateUser = `-- name: UpdateUser :exec
UPDATE
    users
SET
    name = COALESCE($1, name)::text,
    nickname = COALESCE($2, nickname)::text,
    biography = COALESCE($3, biography)::text,
    avatar_image_url = COALESCE($4, avatar_image_url)::text,
    banner_image_url = COALESCE($5, banner_image_url)::text,
    birthdate = COALESCE($6, birthdate)::timestamptz,
    line_id = COALESCE($7, line_id)::text,
    updated_at = COALESCE($8, updated_at)::timestamptz
WHERE
    users.id = $9::uuid
`

type UpdateUserParams struct {
	Name           *string    `json:"name"`
	Nickname       *string    `json:"nickname"`
	Biography      *string    `json:"biography"`
	AvatarImageUrl *string    `json:"avatar_image_url"`
	BannerImageUrl *string    `json:"banner_image_url"`
	Birthdate      *time.Time `json:"birthdate"`
	LineID         *string    `json:"line_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
	UserID         uuid.UUID  `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Name,
		arg.Nickname,
		arg.Biography,
		arg.AvatarImageUrl,
		arg.BannerImageUrl,
		arg.Birthdate,
		arg.LineID,
		arg.UpdatedAt,
		arg.UserID,
	)
	return err
}
