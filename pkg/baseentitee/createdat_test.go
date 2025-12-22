package baseentitee_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func TestNewCreatedAt(t *testing.T) {
	createdAt := baseentitee.NewCreatedAt()
	assert.WithinDuration(t, utilitee.RightNow(), (time.Time)(createdAt), time.Second)
}

func TestCreatedAtFromTime(t *testing.T) {
	specificTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	createdAt := baseentitee.CreatedAtFromTime(specificTime)
	assert.Equal(t, specificTime, (time.Time)(createdAt))
}
