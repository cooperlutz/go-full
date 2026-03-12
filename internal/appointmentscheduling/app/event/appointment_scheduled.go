//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentScheduled struct {
	//
	//AppointmentId string,
	//
	//PetId string,
	//
	//OwnerId string,
	//
	//VeterinarianId string,
	//
	//ScheduledDate string,
	//
	//ScheduledTime string,
	//
	//IsTelemedicine bool,
	//
	// TODO
}

type AppointmentScheduledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentScheduledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentScheduledHandler {
	return AppointmentScheduledHandler{
		publisher: publisher,
	}
}

func (h AppointmentScheduledHandler) Handle(ctx context.Context, event AppointmentScheduled) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_scheduled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_scheduled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
