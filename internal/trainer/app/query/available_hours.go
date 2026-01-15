package query

import (
	"context"
	"errors"
	"time"
)

type AvailableHours struct {
	From time.Time
	To   time.Time
}

// type AvailableHoursHandler decorator.QueryHandler[AvailableHours, []Date]

type AvailableHoursReadModel interface {
	AvailableHours(ctx context.Context, from, to time.Time) ([]Date, error)
}

type AvailableHoursHandler struct {
	readModel AvailableHoursReadModel
}

func NewAvailableHoursHandler(
	readModel AvailableHoursReadModel,
) AvailableHoursHandler {
	return AvailableHoursHandler{readModel: readModel}
}

func (h AvailableHoursHandler) Handle(ctx context.Context, query AvailableHours) (d []Date, err error) {
	if query.From.After(query.To) {
		return nil, errors.New("date-from-after-date-to: Date from after date to") //nolint:goerr113 // generic error
	}

	return h.readModel.AvailableHours(ctx, query.From, query.To)
}
