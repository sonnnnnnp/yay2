// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: leave_call.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const leaveCall = `-- name: LeaveCall :exec
UPDATE
    call_participants
SET
    left_at = NOW()
WHERE
    call_participants.call_id = $1::uuid
    AND call_participants.participant_id = $2::uuid
`

type LeaveCallParams struct {
	CallID        uuid.UUID `json:"call_id"`
	ParticipantID uuid.UUID `json:"participant_id"`
}

func (q *Queries) LeaveCall(ctx context.Context, arg LeaveCallParams) error {
	_, err := q.db.Exec(ctx, leaveCall, arg.CallID, arg.ParticipantID)
	return err
}
