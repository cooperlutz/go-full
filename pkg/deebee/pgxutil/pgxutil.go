package pgxutil

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// TimeToTimestampz converts a time.Time pointer to a pgx Timestamptz type.
func TimeToTimestampz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false}
	}

	return pgtype.Timestamptz{Time: *t, InfinityModifier: pgtype.Finite, Valid: true}
}

// TimestampzToTimePtr converts a pgx Timestamptz type to a time.Time pointer.
func TimestampzToTimePtr(ts pgtype.Timestamptz) *time.Time {
	if !ts.Valid {
		return nil
	}

	return &ts.Time
}

// UUIDToPgtypeUUID converts a uuid.UUID to a pgx UUID type.
func UUIDToPgtypeUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

// StrToPgtypeText converts a string pointer to a pgx Text type.
func StrToPgtypeText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{String: "", Valid: false}
	}

	return pgtype.Text{String: *s, Valid: true}
}

// IntToPgtypeInt4 converts an int pointer to a pgx Int4 type.
func IntToPgtypeInt4(i *int) pgtype.Int4 {
	// Check for overflow/underflow
	if i == nil || utilitee.SafeIntToInt32(i) == 0 {
		return pgtype.Int4{Int32: 0, Valid: false}
	}

	return pgtype.Int4{Int32: utilitee.SafeIntToInt32(i), Valid: true}
}

// Int32ToPgtypeInt4 converts an int32 pointer to a pgx Int4 type.
func Int32ToPgtypeInt4(i *int32) pgtype.Int4 {
	if i == nil {
		return pgtype.Int4{Int32: 0, Valid: false}
	}

	return pgtype.Int4{Int32: *i, Valid: true}
}

// BoolToPgtypeBool converts a bool pointer to a pgx Bool type.
func BoolToPgtypeBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Bool: false, Valid: false}
	}

	return pgtype.Bool{Bool: *b, Valid: true}
}
