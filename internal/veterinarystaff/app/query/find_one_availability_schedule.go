//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneAvailabilitySchedule struct {
	AvailabilityScheduleID string
}

type FindOneAvailabilityScheduleReadModel interface {
	FindOneAvailabilitySchedule(ctx context.Context, availabilityscheduleId uuid.UUID) (AvailabilitySchedule, error)
}

type FindOneAvailabilityScheduleHandler struct {
	readModel FindOneAvailabilityScheduleReadModel
}

func NewFindOneAvailabilityScheduleHandler(
	readModel FindOneAvailabilityScheduleReadModel,
) FindOneAvailabilityScheduleHandler {
	return FindOneAvailabilityScheduleHandler{readModel: readModel}
}

func (h FindOneAvailabilityScheduleHandler) Handle(ctx context.Context, qry FindOneAvailabilitySchedule) (AvailabilitySchedule, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_one_availability_schedule.handle")
	defer span.End()

	availabilityschedule, err := h.readModel.FindOneAvailabilitySchedule(ctx, uuid.MustParse(qry.AvailabilityScheduleID))
	if err != nil {
		return AvailabilitySchedule{}, err
	}

	return availabilityschedule, nil
}
