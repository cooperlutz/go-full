package base

import "time"

type DeletedAt time.Time

func (d DeletedAt) getTime() *time.Time {
	return (*time.Time)(&d)
}

func (d DeletedAt) String() string {
	if d.getTime().IsZero() {
		return ""
	}

	return d.getTime().Format(time.RFC3339)
}

func NewDeletedAt() *DeletedAt {
	return nil
}

func DeletedAtFromTime(t *time.Time) *DeletedAt {
	if t == nil {
		return nil
	}

	d := DeletedAt(*t)

	return &d
}
