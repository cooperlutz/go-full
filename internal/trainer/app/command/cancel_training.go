package command

import (
	"context"
	"errors"
	"time"

	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
)

type CancelTraining struct {
	Hour time.Time
}

type CancelTrainingHandler struct {
	hourRepo hour.Repository
}

func NewCancelTrainingHandler(
	hourRepo hour.Repository,
) CancelTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return CancelTrainingHandler{hourRepo: hourRepo}
}

func (h CancelTrainingHandler) Handle(ctx context.Context, cmd CancelTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(h *hour.Hour) (*hour.Hour, error) {
		if err := h.CancelTraining(); err != nil {
			return nil, err
		}

		return h, nil
	}); err != nil {
		return errors.New(err.Error()) //nolint:goerr113 // wrapping is not needed here
	}

	return nil
}
