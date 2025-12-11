package base

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// UpdatedAt represents the last updated timestamp of an entity.
type UpdatedAt time.Time

// getTime returns the time.Time value of UpdatedAt.
func (u UpdatedAt) getTime() time.Time {
	return time.Time(u)
}

// NewUpdatedAt creates a new UpdatedAt with the current time.
func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(utilitee.RightNow())
}

// UpdatedAtFromTime creates a UpdatedAt from a given time.Time.
func UpdatedAtFromTime(t time.Time) UpdatedAt {
	return UpdatedAt(t)
}
