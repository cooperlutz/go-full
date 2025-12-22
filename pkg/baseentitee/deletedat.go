package baseentitee

import "time"

// DeletedAt represents the deletion timestamp of an entity.
type DeletedAt time.Time

// getTime returns the time.Time value of DeletedAt.
func (d DeletedAt) getTime() *time.Time {
	return (*time.Time)(&d)
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
