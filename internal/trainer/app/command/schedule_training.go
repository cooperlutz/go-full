package command

import (
	"context"
	"errors"
	"time"

	"github.com/cooperlutz/go-full/internal/trainer/domain/hour"
)

type ScheduleTraining struct {
	Hour time.Time
}

// type ScheduleTrainingHandler decorator.CommandHandler[ScheduleTraining]

type ScheduleTrainingHandler struct {
	hourRepo hour.Repository
}

func NewScheduleTrainingHandler(
	hourRepo hour.Repository,
) ScheduleTrainingHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return ScheduleTrainingHandler{hourRepo: hourRepo}
}

func (h ScheduleTrainingHandler) Handle(ctx context.Context, cmd ScheduleTraining) error {
	if err := h.hourRepo.UpdateHour(ctx, cmd.Hour, func(h *hour.Hour) (*hour.Hour, error) {
		if err := h.ScheduleTraining(); err != nil {
			return nil, err
		}

		return h, nil
	}); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
