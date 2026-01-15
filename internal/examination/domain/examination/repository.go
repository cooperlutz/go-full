package examination

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Exam, error)
	// UpdateHour(
	// 	ctx context.Context,
	// 	hourTime time.Time,
	// 	updateFn func(h *Hour) (*Hour, error),
	// ) error
}
