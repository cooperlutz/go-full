package query

import (
	"context"
	"time"

	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
)

type HourAvailability struct {
	Hour time.Time
}

// type HourAvailabilityHandler decorator.QueryHandler[HourAvailability, bool]

type HourAvailabilityHandler struct {
	hourRepo hour.Repository
}

func NewHourAvailabilityHandler(
	hourRepo hour.Repository,
) HourAvailabilityHandler {
	return HourAvailabilityHandler{hourRepo: hourRepo}
}

func (h HourAvailabilityHandler) Handle(ctx context.Context, query HourAvailability) (bool, error) {
	hour, err := h.hourRepo.GetHour(ctx, query.Hour)
	if err != nil {
		return false, err
	}

	return hour.IsAvailable(), nil
}
