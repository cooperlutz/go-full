package baseentitee

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// CreatedAt represents the creation timestamp of an entity.
type CreatedAt time.Time

// NewCreatedAt creates a new CreatedAt with the current time.
func NewCreatedAt() CreatedAt {
	return CreatedAt(utilitee.RightNow())
}

// getTime returns the time.Time value of CreatedAt.
func (c CreatedAt) getTime() time.Time {
	return time.Time(c)
}

// CreatedAtFromTime creates a CreatedAt from a given time.Time.
func CreatedAtFromTime(t time.Time) CreatedAt {
	return CreatedAt(t)
}
