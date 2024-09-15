// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldCoverURL holds the string denoting the cover_url field in the database.
	FieldCoverURL = "cover_url"
	// FieldBiography holds the string denoting the biography field in the database.
	FieldBiography = "biography"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldDisplayName,
	FieldAvatarURL,
	FieldCoverURL,
	FieldBiography,
	FieldBirthdate,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByAvatarURL orders the results by the avatar_url field.
func ByAvatarURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatarURL, opts...).ToFunc()
}

// ByCoverURL orders the results by the cover_url field.
func ByCoverURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverURL, opts...).ToFunc()
}

// ByBiography orders the results by the biography field.
func ByBiography(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBiography, opts...).ToFunc()
}

// ByBirthdate orders the results by the birthdate field.
func ByBirthdate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthdate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
