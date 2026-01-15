package command

import (
	"context"
	"errors"
	"time"

	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
)

type MakeHoursAvailable struct {
	Hours []time.Time
}

type MakeHoursAvailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursAvailableHandler(
	hourRepo hour.Repository,
) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return MakeHoursAvailableHandler{hourRepo: hourRepo}
}

func (c MakeHoursAvailableHandler) Handle(ctx context.Context, cmd MakeHoursAvailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeAvailable(); err != nil {
				return nil, err
			}

			return h, nil
		}); err != nil {
			return errors.New(err.Error()) //nolint:goerr113 // wrapping is not needed here
		}
	}

	return nil
}
