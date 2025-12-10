package base

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type UpdatedAt time.Time

func (u UpdatedAt) getTime() time.Time {
	return time.Time(u)
}

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(utilitee.RightNow())
}

func UpdatedAtFromTime(t time.Time) UpdatedAt {
	return UpdatedAt(t)
}
