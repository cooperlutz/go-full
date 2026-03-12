//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentCompleted struct {
	//
	//AppointmentId string,
	//
	//PetId string,
	//
	//VeterinarianId string,
	//
	//OwnerId string,
	//
	// TODO
}

type AppointmentCompletedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentCompletedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentCompletedHandler {
	return AppointmentCompletedHandler{
		publisher: publisher,
	}
}

func (h AppointmentCompletedHandler) Handle(ctx context.Context, event AppointmentCompleted) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_completed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_completed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
