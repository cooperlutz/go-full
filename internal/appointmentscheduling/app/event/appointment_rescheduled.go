//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentRescheduled struct {
	//
	//AppointmentId string,
	//
	//NewScheduledDate string,
	//
	//NewScheduledTime string,
	//
	// TODO
}

type AppointmentRescheduledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentRescheduledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentRescheduledHandler {
	return AppointmentRescheduledHandler{
		publisher: publisher,
	}
}

func (h AppointmentRescheduledHandler) Handle(ctx context.Context, event AppointmentRescheduled) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_rescheduled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_rescheduled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
