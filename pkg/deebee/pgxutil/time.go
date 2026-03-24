package pgxutil

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// PgtypeTimeNullToTimePtr converts a sql.NullTime type to a time.Time pointer, returning nil if the sql.NullTime is null.
func PgtypeTimeNullToTimePtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}

	return &t.Time
}

// TimeToPgtypeTimeNull converts a time.Time pointer to a sql.NullTime type, treating nil pointers as null values.
func TimeToPgtypeTimeNull(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Time: time.Time{}, Valid: false}
	}

	return sql.NullTime{Time: *t, Valid: true}
}

// TimeToPgtypeDate converts a time.Time pointer to a pgx Date type.
func TimeToPgtypeDate(t *time.Time) pgtype.Date {
	if t == nil {
		return pgtype.Date{Time: time.Time{}, Valid: false}
	}

	return pgtype.Date{Time: *t, Valid: true}
}

// TimeToPgtypeTime converts a time.Time pointer to a pgx Time type, representing the time of day without the date component. If the pointer is nil, it returns a null pgx Time.
func TimeToPgtypeTime(t *time.Time) pgtype.Time {
	if t == nil {
		return pgtype.Time{Microseconds: 0, Valid: false}
	}
	// return the microseconds since midnight, which is how pgx represents time without date
	return pgtype.Time{Microseconds: t.Sub(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())).Microseconds(), Valid: true}
}
