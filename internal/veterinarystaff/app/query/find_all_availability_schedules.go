//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllAvailabilitySchedulesReadModel interface {
	FindAllAvailabilitySchedules(ctx context.Context) ([]AvailabilitySchedule, error)
}

type FindAllAvailabilitySchedulesHandler struct {
	readModel FindAllAvailabilitySchedulesReadModel
}

func NewFindAllAvailabilitySchedulesHandler(
	readModel FindAllAvailabilitySchedulesReadModel,
) FindAllAvailabilitySchedulesHandler {
	return FindAllAvailabilitySchedulesHandler{readModel: readModel}
}

func (h FindAllAvailabilitySchedulesHandler) Handle(ctx context.Context) ([]AvailabilitySchedule, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_all_availability_schedules.handle")
	defer span.End()

	exams, err := h.readModel.FindAllAvailabilitySchedules(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
