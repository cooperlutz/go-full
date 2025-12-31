package pgxutil_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
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

func TestUUIDToPgtypeUUID(t *testing.T) {
	id := uuid.New()
	pgUUID := pgxutil.UUIDToPgtypeUUID(id)

	assert.True(t, pgUUID.Valid)
	assert.Equal(t, [16]byte(id[:]), pgUUID.Bytes)
}

func TestStrToPgtypeText(t *testing.T) {
	str := "Hello, World!"
	pgText := pgxutil.StrToPgtypeText(&str)

	assert.True(t, pgText.Valid)
	assert.Equal(t, str, pgText.String)

	var nilStr *string
	pgTextNil := pgxutil.StrToPgtypeText(nilStr)
	assert.False(t, pgTextNil.Valid)
}
