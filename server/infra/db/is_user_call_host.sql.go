// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: is_user_call_host.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const isUserCallHost = `-- name: IsUserCallHost :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            calls
        WHERE
            calls.host_id = $1::uuid
            AND calls.id = $2::uuid
    ) AS is_host
`

type IsUserCallHostParams struct {
	UserID uuid.UUID `json:"user_id"`
	CallID uuid.UUID `json:"call_id"`
}

func (q *Queries) IsUserCallHost(ctx context.Context, arg IsUserCallHostParams) (bool, error) {
	row := q.db.QueryRow(ctx, isUserCallHost, arg.UserID, arg.CallID)
	var is_host bool
	err := row.Scan(&is_host)
	return is_host, err
}