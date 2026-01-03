package pgxutil

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
