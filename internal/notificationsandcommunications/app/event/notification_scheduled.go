//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type NotificationScheduled struct {
	//
	//NotificationId string,
	//
	//RecipientId string,
	//
	//ScheduledSendTime string,
	//
	// TODO
}

type NotificationScheduledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewNotificationScheduledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) NotificationScheduledHandler {
	return NotificationScheduledHandler{
		publisher: publisher,
	}
}

func (h NotificationScheduledHandler) Handle(ctx context.Context, event NotificationScheduled) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.event.notification_scheduled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("notificationsandcommunications.notification_scheduled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
