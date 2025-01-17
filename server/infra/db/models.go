// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CallJoinableBy string

const (
	CallJoinableByAll       CallJoinableBy = "all"
	CallJoinableByFollowers CallJoinableBy = "followers"
	CallJoinableByFriends   CallJoinableBy = "friends"
	CallJoinableByNobody    CallJoinableBy = "nobody"
)

func (e *CallJoinableBy) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CallJoinableBy(s)
	case string:
		*e = CallJoinableBy(s)
	default:
		return fmt.Errorf("unsupported scan type for CallJoinableBy: %T", src)
	}
	return nil
}

type NullCallJoinableBy struct {
	CallJoinableBy CallJoinableBy `json:"call_joinable_by"`
	Valid          bool           `json:"valid"` // Valid is true if CallJoinableBy is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCallJoinableBy) Scan(value interface{}) error {
	if value == nil {
		ns.CallJoinableBy, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CallJoinableBy.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCallJoinableBy) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CallJoinableBy), nil
}

type CallParticipantRole string

const (
	CallParticipantRoleHost        CallParticipantRole = "host"
	CallParticipantRoleCoHost      CallParticipantRole = "co-host"
	CallParticipantRoleParticipant CallParticipantRole = "participant"
)

func (e *CallParticipantRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CallParticipantRole(s)
	case string:
		*e = CallParticipantRole(s)
	default:
		return fmt.Errorf("unsupported scan type for CallParticipantRole: %T", src)
	}
	return nil
}

type NullCallParticipantRole struct {
	CallParticipantRole CallParticipantRole `json:"call_participant_role"`
	Valid               bool                `json:"valid"` // Valid is true if CallParticipantRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCallParticipantRole) Scan(value interface{}) error {
	if value == nil {
		ns.CallParticipantRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CallParticipantRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCallParticipantRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CallParticipantRole), nil
}

type CallType string

const (
	CallTypeVoice CallType = "voice"
	CallTypeVideo CallType = "video"
)

func (e *CallType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CallType(s)
	case string:
		*e = CallType(s)
	default:
		return fmt.Errorf("unsupported scan type for CallType: %T", src)
	}
	return nil
}

type NullCallType struct {
	CallType CallType `json:"call_type"`
	Valid    bool     `json:"valid"` // Valid is true if CallType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCallType) Scan(value interface{}) error {
	if value == nil {
		ns.CallType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CallType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCallType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CallType), nil
}

type Call struct {
	ID         uuid.UUID          `json:"id"`
	Title      *string            `json:"title"`
	Type       CallType           `json:"type"`
	JoinableBy CallJoinableBy     `json:"joinable_by"`
	HostID     uuid.UUID          `json:"host_id"`
	StartedAt  pgtype.Timestamptz `json:"started_at"`
	EndedAt    pgtype.Timestamptz `json:"ended_at"`
}

type CallParticipant struct {
	CallID        uuid.UUID           `json:"call_id"`
	ParticipantID uuid.UUID           `json:"participant_id"`
	Role          CallParticipantRole `json:"role"`
	JoinedAt      pgtype.Timestamptz  `json:"joined_at"`
	LeftAt        pgtype.Timestamptz  `json:"left_at"`
}

type Post struct {
	ID        uuid.UUID          `json:"id"`
	AuthorID  uuid.UUID          `json:"author_id"`
	ReplyToID *uuid.UUID         `json:"reply_to_id"`
	RepostID  *uuid.UUID         `json:"repost_id"`
	Text      *string            `json:"text"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type PostFavorite struct {
	UserID    uuid.UUID          `json:"user_id"`
	PostID    uuid.UUID          `json:"post_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID             uuid.UUID          `json:"id"`
	CustomID       string             `json:"custom_id"`
	Nickname       string             `json:"nickname"`
	Biography      *string            `json:"biography"`
	AvatarImageUrl *string            `json:"avatar_image_url"`
	BannerImageUrl *string            `json:"banner_image_url"`
	IsPrivate      *bool              `json:"is_private"`
	Birthdate      pgtype.Timestamptz `json:"birthdate"`
	LineID         *string            `json:"line_id"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}

type UserBlock struct {
	BlockerID uuid.UUID          `json:"blocker_id"`
	BlockedID uuid.UUID          `json:"blocked_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type UserFollow struct {
	FollowerID uuid.UUID          `json:"follower_id"`
	FollowedID uuid.UUID          `json:"followed_id"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
}
