//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type NotificationCancelled struct {
	//
	//NotificationId string,
	//
	// TODO
}

type NotificationCancelledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewNotificationCancelledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) NotificationCancelledHandler {
	return NotificationCancelledHandler{
		publisher: publisher,
	}
}

func (h NotificationCancelledHandler) Handle(ctx context.Context, event NotificationCancelled) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.event.notification_cancelled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("notificationsandcommunications.notification_canceled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
