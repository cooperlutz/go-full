//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type VaccinationDueReminderScheduled struct {
	//
	//PetId string,
	//
	//VaccineName string,
	//
	//ExpiryDate string,
	//
	// TODO
}

type VaccinationDueReminderScheduledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewVaccinationDueReminderScheduledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) VaccinationDueReminderScheduledHandler {
	return VaccinationDueReminderScheduledHandler{
		publisher: publisher,
	}
}

func (h VaccinationDueReminderScheduledHandler) Handle(ctx context.Context, event VaccinationDueReminderScheduled) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.event.vaccination_due_reminder_scheduled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("notificationsandcommunications.vaccination_due_reminder_scheduled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
