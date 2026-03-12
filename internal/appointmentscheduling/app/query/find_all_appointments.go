//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllAppointmentsReadModel interface {
	FindAllAppointments(ctx context.Context) ([]Appointment, error)
}

type FindAllAppointmentsHandler struct {
	readModel FindAllAppointmentsReadModel
}

func NewFindAllAppointmentsHandler(
	readModel FindAllAppointmentsReadModel,
) FindAllAppointmentsHandler {
	return FindAllAppointmentsHandler{readModel: readModel}
}

func (h FindAllAppointmentsHandler) Handle(ctx context.Context) ([]Appointment, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.query.find_all_appointments.handle")
	defer span.End()

	exams, err := h.readModel.FindAllAppointments(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
