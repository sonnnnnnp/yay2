// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// AvatarURL applies equality check predicate on the "avatar_url" field. It's identical to AvatarURLEQ.
func AvatarURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// CoverURL applies equality check predicate on the "cover_url" field. It's identical to CoverURLEQ.
func CoverURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCoverURL, v))
}

// Biography applies equality check predicate on the "biography" field. It's identical to BiographyEQ.
func Biography(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// Birthdate applies equality check predicate on the "birthdate" field. It's identical to BirthdateEQ.
func Birthdate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthdate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// AvatarURLEQ applies the EQ predicate on the "avatar_url" field.
func AvatarURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// AvatarURLNEQ applies the NEQ predicate on the "avatar_url" field.
func AvatarURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarURL, v))
}

// AvatarURLIn applies the In predicate on the "avatar_url" field.
func AvatarURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarURL, vs...))
}

// AvatarURLNotIn applies the NotIn predicate on the "avatar_url" field.
func AvatarURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarURL, vs...))
}

// AvatarURLGT applies the GT predicate on the "avatar_url" field.
func AvatarURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarURL, v))
}

// AvatarURLGTE applies the GTE predicate on the "avatar_url" field.
func AvatarURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarURL, v))
}

// AvatarURLLT applies the LT predicate on the "avatar_url" field.
func AvatarURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarURL, v))
}

// AvatarURLLTE applies the LTE predicate on the "avatar_url" field.
func AvatarURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarURL, v))
}

// AvatarURLContains applies the Contains predicate on the "avatar_url" field.
func AvatarURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatarURL, v))
}

// AvatarURLHasPrefix applies the HasPrefix predicate on the "avatar_url" field.
func AvatarURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatarURL, v))
}

// AvatarURLHasSuffix applies the HasSuffix predicate on the "avatar_url" field.
func AvatarURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatarURL, v))
}

// AvatarURLIsNil applies the IsNil predicate on the "avatar_url" field.
func AvatarURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatarURL))
}

// AvatarURLNotNil applies the NotNil predicate on the "avatar_url" field.
func AvatarURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatarURL))
}

// AvatarURLEqualFold applies the EqualFold predicate on the "avatar_url" field.
func AvatarURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatarURL, v))
}

// AvatarURLContainsFold applies the ContainsFold predicate on the "avatar_url" field.
func AvatarURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatarURL, v))
}

// CoverURLEQ applies the EQ predicate on the "cover_url" field.
func CoverURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCoverURL, v))
}

// CoverURLNEQ applies the NEQ predicate on the "cover_url" field.
func CoverURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCoverURL, v))
}

// CoverURLIn applies the In predicate on the "cover_url" field.
func CoverURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCoverURL, vs...))
}

// CoverURLNotIn applies the NotIn predicate on the "cover_url" field.
func CoverURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCoverURL, vs...))
}

// CoverURLGT applies the GT predicate on the "cover_url" field.
func CoverURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCoverURL, v))
}

// CoverURLGTE applies the GTE predicate on the "cover_url" field.
func CoverURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCoverURL, v))
}

// CoverURLLT applies the LT predicate on the "cover_url" field.
func CoverURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCoverURL, v))
}

// CoverURLLTE applies the LTE predicate on the "cover_url" field.
func CoverURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCoverURL, v))
}

// CoverURLContains applies the Contains predicate on the "cover_url" field.
func CoverURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCoverURL, v))
}

// CoverURLHasPrefix applies the HasPrefix predicate on the "cover_url" field.
func CoverURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCoverURL, v))
}

// CoverURLHasSuffix applies the HasSuffix predicate on the "cover_url" field.
func CoverURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCoverURL, v))
}

// CoverURLIsNil applies the IsNil predicate on the "cover_url" field.
func CoverURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldCoverURL))
}

// CoverURLNotNil applies the NotNil predicate on the "cover_url" field.
func CoverURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldCoverURL))
}

// CoverURLEqualFold applies the EqualFold predicate on the "cover_url" field.
func CoverURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCoverURL, v))
}

// CoverURLContainsFold applies the ContainsFold predicate on the "cover_url" field.
func CoverURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCoverURL, v))
}

// BiographyEQ applies the EQ predicate on the "biography" field.
func BiographyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// BiographyNEQ applies the NEQ predicate on the "biography" field.
func BiographyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBiography, v))
}

// BiographyIn applies the In predicate on the "biography" field.
func BiographyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBiography, vs...))
}

// BiographyNotIn applies the NotIn predicate on the "biography" field.
func BiographyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBiography, vs...))
}

// BiographyGT applies the GT predicate on the "biography" field.
func BiographyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBiography, v))
}

// BiographyGTE applies the GTE predicate on the "biography" field.
func BiographyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBiography, v))
}

// BiographyLT applies the LT predicate on the "biography" field.
func BiographyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBiography, v))
}

// BiographyLTE applies the LTE predicate on the "biography" field.
func BiographyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBiography, v))
}

// BiographyContains applies the Contains predicate on the "biography" field.
func BiographyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBiography, v))
}

// BiographyHasPrefix applies the HasPrefix predicate on the "biography" field.
func BiographyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBiography, v))
}

// BiographyHasSuffix applies the HasSuffix predicate on the "biography" field.
func BiographyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBiography, v))
}

// BiographyIsNil applies the IsNil predicate on the "biography" field.
func BiographyIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBiography))
}

// BiographyNotNil applies the NotNil predicate on the "biography" field.
func BiographyNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBiography))
}

// BiographyEqualFold applies the EqualFold predicate on the "biography" field.
func BiographyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBiography, v))
}

// BiographyContainsFold applies the ContainsFold predicate on the "biography" field.
func BiographyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBiography, v))
}

// BirthdateEQ applies the EQ predicate on the "birthdate" field.
func BirthdateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthdate, v))
}

// BirthdateNEQ applies the NEQ predicate on the "birthdate" field.
func BirthdateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBirthdate, v))
}

// BirthdateIn applies the In predicate on the "birthdate" field.
func BirthdateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldBirthdate, vs...))
}

// BirthdateNotIn applies the NotIn predicate on the "birthdate" field.
func BirthdateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBirthdate, vs...))
}

// BirthdateGT applies the GT predicate on the "birthdate" field.
func BirthdateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldBirthdate, v))
}

// BirthdateGTE applies the GTE predicate on the "birthdate" field.
func BirthdateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBirthdate, v))
}

// BirthdateLT applies the LT predicate on the "birthdate" field.
func BirthdateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldBirthdate, v))
}

// BirthdateLTE applies the LTE predicate on the "birthdate" field.
func BirthdateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBirthdate, v))
}

// BirthdateIsNil applies the IsNil predicate on the "birthdate" field.
func BirthdateIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBirthdate))
}

// BirthdateNotNil applies the NotNil predicate on the "birthdate" field.
func BirthdateNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBirthdate))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
