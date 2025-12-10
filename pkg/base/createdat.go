package base

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type CreatedAt time.Time

func (c CreatedAt) getTime() time.Time {
	return time.Time(c)
}

func NewCreatedAt() CreatedAt {
	return CreatedAt(utilitee.RightNow())
}

func CreatedAtFromTime(t time.Time) CreatedAt {
	return CreatedAt(t)
}
