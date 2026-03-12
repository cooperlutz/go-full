//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentConfirmed struct {
	//
	//AppointmentId string,
	//
	//OwnerId string,
	//
	// TODO
}

type AppointmentConfirmedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentConfirmedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentConfirmedHandler {
	return AppointmentConfirmedHandler{
		publisher: publisher,
	}
}

func (h AppointmentConfirmedHandler) Handle(ctx context.Context, event AppointmentConfirmed) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_confirmed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_confirmed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
