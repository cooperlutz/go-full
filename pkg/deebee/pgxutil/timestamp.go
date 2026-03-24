package pgxutil

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// TimestampzToTimePtr converts a pgx Timestamptz type to a time.Time pointer.
func TimestampzToTimePtr(ts pgtype.Timestamptz) *time.Time {
	if !ts.Valid {
		return nil
	}

	return &ts.Time
}

// TimeToTimestampz converts a time.Time pointer to a pgx Timestamptz type.
func TimeToTimestampz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false}
	}

	return pgtype.Timestamptz{Time: *t, InfinityModifier: pgtype.Finite, Valid: true}
}

// TimePtrToTimestampz converts a time.Time pointer to a pgx Timestamptz type, handling nil pointers appropriately.
func TimeToPgtypeTimestampz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, InfinityModifier: pgtype.Finite, Valid: true}
}

// TimePtrToPgtypeTimestampz converts a time.Time pointer to a pgx Timestamptz type, handling nil pointers appropriately.
func TimePtrToPgtypeTimestampz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false}
	}

	return pgtype.Timestamptz{Time: *t, InfinityModifier: pgtype.Finite, Valid: true}
}

// TimePtrSliceToPgtypeTimestampzSlice converts a slice of time.Time pointers to a slice of pgx Timestamptz types, handling nil pointers appropriately.
func TimePtrSliceToPgtypeTimestampzSlice(times []*time.Time) []pgtype.Timestamptz {
	pgTimes := make([]pgtype.Timestamptz, len(times))
	for i, t := range times {
		pgTimes[i] = TimePtrToPgtypeTimestampz(t)
	}

	return pgTimes
}

// TimeSliceToPgtypeTimestampzSlice converts a slice of time.Time values to a slice of pgx Timestamptz types.
func TimeSliceToPgtypeTimestampzSlice(times []time.Time) []pgtype.Timestamptz {
	pgTimes := make([]pgtype.Timestamptz, len(times))
	for i, t := range times {
		pgTimes[i] = TimeToPgtypeTimestampz(t)
	}

	return pgTimes
}

// PgtypeTimestampzToTime converts a pgx Timestamptz type to a time.Time value, returning zero time if the pgx Timestamptz is null.
func PgtypeTimestampzToTime(ts pgtype.Timestamptz) time.Time {
	if !ts.Valid {
		return time.Time{}
	}

	return ts.Time
}

// PgtypeTimestampzToTimePtr converts a pgx Timestamptz type to a time.Time pointer, returning nil if the pgx Timestamptz is null.
func PgtypeTimestampzSliceToTimeSlice(ts []pgtype.Timestamptz) []time.Time {
	times := make([]time.Time, len(ts))
	for i, t := range ts {
		times[i] = PgtypeTimestampzToTime(t)
	}

	return times
}

// PgtypeTimestampzSliceToTimeSlicePtr converts a slice of pgx Timestamptz types to a slice of time.Time pointers, returning nil for any null pgx Timestamptz values.
func PgtypeTimestampzSliceToTimeSlicePtr(ts []pgtype.Timestamptz) []*time.Time {
	times := make([]*time.Time, len(ts))
	for i, t := range ts {
		if t.Valid {
			times[i] = &t.Time
		} else {
			times[i] = nil
		}
	}

	return times
}
