package pgxutil_test

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestTimeToTimestampzAndBack(t *testing.T) {
	now := time.Now()
	ts := pgxutil.TimeToTimestampz(&now)
	assert.True(t, ts.Valid)
	assert.Equal(t, now, ts.Time)

	timePtr := pgxutil.TimestampzToTimePtr(ts)
	assert.NotNil(t, timePtr)
	assert.Equal(t, now, *timePtr)

	var nilTime *time.Time
	tsNil := pgxutil.TimeToTimestampz(nilTime)
	assert.False(t, tsNil.Valid)

	timePtrNil := pgxutil.TimestampzToTimePtr(tsNil)
	assert.Nil(t, timePtrNil)
}

func TestTimeToPgtypeTimestampz(t *testing.T) {
	now := time.Now()
	ts := pgxutil.TimeToPgtypeTimestampz(now)
	assert.True(t, ts.Valid)
	assert.Equal(t, now, ts.Time)
	assert.Equal(t, pgtype.Finite, ts.InfinityModifier)

	zeroTime := time.Time{}
	tsZero := pgxutil.TimeToPgtypeTimestampz(zeroTime)
	assert.True(t, tsZero.Valid)
	assert.Equal(t, zeroTime, tsZero.Time)
	assert.Equal(t, pgtype.Finite, tsZero.InfinityModifier)
}

func TestTimePtrToPgtypeTimestampz(t *testing.T) {
	now := time.Now()
	ts := pgxutil.TimePtrToPgtypeTimestampz(&now)
	assert.True(t, ts.Valid)
	assert.Equal(t, now, ts.Time)
	assert.Equal(t, pgtype.Finite, ts.InfinityModifier)

	var nilTime *time.Time
	tsNil := pgxutil.TimePtrToPgtypeTimestampz(nilTime)
	assert.False(t, tsNil.Valid)
	assert.Equal(t, time.Time{}, tsNil.Time)
	assert.Equal(t, pgtype.Finite, tsNil.InfinityModifier)

	zeroTime := time.Time{}
	tsZero := pgxutil.TimePtrToPgtypeTimestampz(&zeroTime)
	assert.True(t, tsZero.Valid)
	assert.Equal(t, zeroTime, tsZero.Time)
	assert.Equal(t, pgtype.Finite, tsZero.InfinityModifier)
}

func TestTimePtrSliceToPgtypeTimestampzSlice(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	times := []*time.Time{&now, &yesterday, nil, &tomorrow}
	result := pgxutil.TimePtrSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 4)
	assert.True(t, result[0].Valid)
	assert.Equal(t, now, result[0].Time)
	assert.Equal(t, pgtype.Finite, result[0].InfinityModifier)

	assert.True(t, result[1].Valid)
	assert.Equal(t, yesterday, result[1].Time)
	assert.Equal(t, pgtype.Finite, result[1].InfinityModifier)

	assert.False(t, result[2].Valid)
	assert.Equal(t, time.Time{}, result[2].Time)
	assert.Equal(t, pgtype.Finite, result[2].InfinityModifier)

	assert.True(t, result[3].Valid)
	assert.Equal(t, tomorrow, result[3].Time)
	assert.Equal(t, pgtype.Finite, result[3].InfinityModifier)
}

func TestTimePtrSliceToPgtypeTimestampzSliceEmpty(t *testing.T) {
	var times []*time.Time
	result := pgxutil.TimePtrSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 0)
}

func TestTimePtrSliceToPgtypeTimestampzSliceAllNil(t *testing.T) {
	times := []*time.Time{nil, nil, nil}
	result := pgxutil.TimePtrSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 3)
	for _, ts := range result {
		assert.False(t, ts.Valid)
		assert.Equal(t, time.Time{}, ts.Time)
	}
}

func TestTimeSliceToPgtypeTimestampzSlice(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	times := []time.Time{now, yesterday, tomorrow}
	result := pgxutil.TimeSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 3)
	assert.True(t, result[0].Valid)
	assert.Equal(t, now, result[0].Time)
	assert.Equal(t, pgtype.Finite, result[0].InfinityModifier)

	assert.True(t, result[1].Valid)
	assert.Equal(t, yesterday, result[1].Time)
	assert.Equal(t, pgtype.Finite, result[1].InfinityModifier)

	assert.True(t, result[2].Valid)
	assert.Equal(t, tomorrow, result[2].Time)
	assert.Equal(t, pgtype.Finite, result[2].InfinityModifier)
}

func TestTimeSliceToPgtypeTimestampzSliceEmpty(t *testing.T) {
	var times []time.Time
	result := pgxutil.TimeSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 0)
}

func TestTimeSliceToPgtypeTimestampzSliceWithZeroTime(t *testing.T) {
	now := time.Now()
	zeroTime := time.Time{}

	times := []time.Time{now, zeroTime}
	result := pgxutil.TimeSliceToPgtypeTimestampzSlice(times)

	assert.Len(t, result, 2)
	assert.True(t, result[0].Valid)
	assert.Equal(t, now, result[0].Time)

	assert.True(t, result[1].Valid)
	assert.Equal(t, zeroTime, result[1].Time)
}

func TestPgtypeTimestampzToTime(t *testing.T) {
	now := time.Now()
	ts := pgtype.Timestamptz{Time: now, InfinityModifier: pgtype.Finite, Valid: true}
	result := pgxutil.PgtypeTimestampzToTime(ts)

	assert.Equal(t, now, result)
}

func TestPgtypeTimestampzToTimeInvalid(t *testing.T) {
	ts := pgtype.Timestamptz{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false}
	result := pgxutil.PgtypeTimestampzToTime(ts)

	assert.Equal(t, time.Time{}, result)
}

func TestPgtypeTimestampzToTimeZeroTime(t *testing.T) {
	zeroTime := time.Time{}
	ts := pgtype.Timestamptz{Time: zeroTime, InfinityModifier: pgtype.Finite, Valid: true}
	result := pgxutil.PgtypeTimestampzToTime(ts)

	assert.Equal(t, zeroTime, result)
}

func TestPgtypeTimestampzSliceToTimeSlice(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	ts := []pgtype.Timestamptz{
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: yesterday, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: tomorrow, InfinityModifier: pgtype.Finite, Valid: true},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlice(ts)

	assert.Len(t, result, 3)
	assert.Equal(t, now, result[0])
	assert.Equal(t, yesterday, result[1])
	assert.Equal(t, tomorrow, result[2])
}

func TestPgtypeTimestampzSliceToTimeSliceEmpty(t *testing.T) {
	var ts []pgtype.Timestamptz
	result := pgxutil.PgtypeTimestampzSliceToTimeSlice(ts)

	assert.Len(t, result, 0)
}

func TestPgtypeTimestampzSliceToTimeSliceWithInvalidValues(t *testing.T) {
	now := time.Now()

	ts := []pgtype.Timestamptz{
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlice(ts)

	assert.Len(t, result, 3)
	assert.Equal(t, now, result[0])
	assert.Equal(t, time.Time{}, result[1])
	assert.Equal(t, now, result[2])
}

func TestPgtypeTimestampzSliceToTimeSliceAllInvalid(t *testing.T) {
	ts := []pgtype.Timestamptz{
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlice(ts)

	assert.Len(t, result, 2)
	assert.Equal(t, time.Time{}, result[0])
	assert.Equal(t, time.Time{}, result[1])
}

func TestPgtypeTimestampzSliceToTimeSlicePtrEmpty(t *testing.T) {
	var ts []pgtype.Timestamptz
	result := pgxutil.PgtypeTimestampzSliceToTimeSlicePtr(ts)

	assert.Len(t, result, 0)
}

func TestPgtypeTimestampzSliceToTimeSlicePtrWithValidValues(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	ts := []pgtype.Timestamptz{
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: yesterday, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: tomorrow, InfinityModifier: pgtype.Finite, Valid: true},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlicePtr(ts)

	assert.Len(t, result, 3)
	assert.NotNil(t, result[0])
	assert.Equal(t, now, *result[0])
	assert.NotNil(t, result[1])
	assert.Equal(t, yesterday, *result[1])
	assert.NotNil(t, result[2])
	assert.Equal(t, tomorrow, *result[2])
}

func TestPgtypeTimestampzSliceToTimeSlicePtrWithNilValues(t *testing.T) {
	now := time.Now()

	ts := []pgtype.Timestamptz{
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
		{Time: now, InfinityModifier: pgtype.Finite, Valid: true},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlicePtr(ts)

	assert.Len(t, result, 3)
	assert.NotNil(t, result[0])
	assert.Equal(t, now, *result[0])
	assert.Nil(t, result[1])
	assert.NotNil(t, result[2])
	assert.Equal(t, now, *result[2])
}

func TestPgtypeTimestampzSliceToTimeSlicePtrAllInvalid(t *testing.T) {
	ts := []pgtype.Timestamptz{
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
		{Time: time.Time{}, InfinityModifier: pgtype.Finite, Valid: false},
	}
	result := pgxutil.PgtypeTimestampzSliceToTimeSlicePtr(ts)

	assert.Len(t, result, 2)
	assert.Nil(t, result[0])
	assert.Nil(t, result[1])
}
