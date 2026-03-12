//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneAppointment struct {
	AppointmentID string
}

type FindOneAppointmentReadModel interface {
	FindOneAppointment(ctx context.Context, appointmentId uuid.UUID) (Appointment, error)
}

type FindOneAppointmentHandler struct {
	readModel FindOneAppointmentReadModel
}

func NewFindOneAppointmentHandler(
	readModel FindOneAppointmentReadModel,
) FindOneAppointmentHandler {
	return FindOneAppointmentHandler{readModel: readModel}
}

func (h FindOneAppointmentHandler) Handle(ctx context.Context, qry FindOneAppointment) (Appointment, error) {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.query.find_one_appointment.handle")
	defer span.End()

	appointment, err := h.readModel.FindOneAppointment(ctx, uuid.MustParse(qry.AppointmentID))
	if err != nil {
		return Appointment{}, err
	}

	return appointment, nil
}
