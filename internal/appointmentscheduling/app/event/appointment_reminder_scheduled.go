//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentReminderScheduled struct {
	//
	//AppointmentId string,
	//
	//OwnerId string,
	//
	//ScheduledDate string,
	//
	//ScheduledTime string,
	//
	// TODO
}

type AppointmentReminderScheduledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentReminderScheduledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentReminderScheduledHandler {
	return AppointmentReminderScheduledHandler{
		publisher: publisher,
	}
}

func (h AppointmentReminderScheduledHandler) Handle(ctx context.Context, event AppointmentReminderScheduled) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_reminder_scheduled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_reminder_scheduled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
