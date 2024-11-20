// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get_block_status.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getBlockStatus = `-- name: GetBlockStatus :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            user_blocks
        WHERE
            blocker_id = $1::uuid
            AND blocked_id = $2::uuid
    ) AS blocking,
    EXISTS (
        SELECT
            1
        FROM
            user_blocks
        WHERE
            blocker_id = $2::uuid
            AND blocked_id = $1::uuid
    ) AS blocked_by
`

type GetBlockStatusParams struct {
	SelfID   uuid.UUID `json:"self_id"`
	TargetID uuid.UUID `json:"target_id"`
}

type GetBlockStatusRow struct {
	Blocking  bool `json:"blocking"`
	BlockedBy bool `json:"blocked_by"`
}

func (q *Queries) GetBlockStatus(ctx context.Context, arg GetBlockStatusParams) (GetBlockStatusRow, error) {
	row := q.db.QueryRow(ctx, getBlockStatus, arg.SelfID, arg.TargetID)
	var i GetBlockStatusRow
	err := row.Scan(&i.Blocking, &i.BlockedBy)
	return i, err
}
