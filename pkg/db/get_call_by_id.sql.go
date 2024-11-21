// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get_call_by_id.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getCallByID = `-- name: GetCallByID :one
SELECT
    calls.id, calls.title, calls.type, calls.joinable_by, calls.host_id, calls.started_at, calls.ended_at,
    json_agg(json_build_object(
        'role', call_participants.role,
        'user', to_jsonb(users) || jsonb_build_object(
            'block_status', jsonb_build_object(
                'blocking', EXISTS (
                    SELECT 1
                    FROM user_blocks
                    WHERE blocker_id = $1::uuid
                    AND blocked_id = users.id
                ),
                'blocked_by', EXISTS (
                    SELECT 1
                    FROM user_blocks
                    WHERE blocker_id = users.id
                    AND blocked_id = $1::uuid
                )
            )
        )
    )) AS participants
FROM
    calls
    LEFT JOIN
        call_participants
        ON calls.id = call_participants.call_id
    LEFT JOIN
        users
        ON call_participants.participant_id = users.id
WHERE
    calls.id = $2::uuid
GROUP BY
    calls.id
`

type GetCallByIDParams struct {
	SelfID uuid.UUID `json:"self_id"`
	CallID uuid.UUID `json:"call_id"`
}

type GetCallByIDRow struct {
	Call         Call   `json:"call"`
	Participants []byte `json:"participants"`
}

func (q *Queries) GetCallByID(ctx context.Context, arg GetCallByIDParams) (GetCallByIDRow, error) {
	row := q.db.QueryRow(ctx, getCallByID, arg.SelfID, arg.CallID)
	var i GetCallByIDRow
	err := row.Scan(
		&i.Call.ID,
		&i.Call.Title,
		&i.Call.Type,
		&i.Call.JoinableBy,
		&i.Call.HostID,
		&i.Call.StartedAt,
		&i.Call.EndedAt,
		&i.Participants,
	)
	return i, err
}
