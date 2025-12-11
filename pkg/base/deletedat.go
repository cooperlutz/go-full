package base

import "time"

// DeletedAt represents the deletion timestamp of an entity.
type DeletedAt time.Time

// getTime returns the time.Time value of DeletedAt.
func (d DeletedAt) getTime() *time.Time {
	return (*time.Time)(&d)
}

// String returns the string representation of DeletedAt in RFC3339 format.
func (d DeletedAt) String() string {
	if d.getTime().IsZero() {
		return ""
	}

	return d.getTime().Format(time.RFC3339)
}

// NewDeletedAt creates a new DeletedAt set to nil (not deleted).
func NewDeletedAt() *DeletedAt {
	return nil
}

// DeletedAtFromTime creates a DeletedAt from a given *time.Time.
func DeletedAtFromTime(t *time.Time) *DeletedAt {
	if t == nil {
		return nil
	}

	d := DeletedAt(*t)

	return &d
}
