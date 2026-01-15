package command

import (
	"context"
	"errors"
	"time"

	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
)

type MakeHoursUnavailable struct {
	Hours []time.Time
}

// type MakeHoursUnavailableHandler decorator.CommandHandler[MakeHoursUnavailable]

type MakeHoursUnavailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursUnavailableHandler(
	hourRepo hour.Repository,
) MakeHoursUnavailableHandler {
	if hourRepo == nil {
		panic("hourRepo is nil")
	}

	return MakeHoursUnavailableHandler{hourRepo: hourRepo}
}

func (c MakeHoursUnavailableHandler) Handle(ctx context.Context, cmd MakeHoursUnavailable) error {
	for _, hourToUpdate := range cmd.Hours {
		if err := c.hourRepo.UpdateHour(ctx, hourToUpdate, func(h *hour.Hour) (*hour.Hour, error) {
			if err := h.MakeNotAvailable(); err != nil {
				return nil, err
			}

			return h, nil
		}); err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}
