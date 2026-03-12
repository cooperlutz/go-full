//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type NotificationSent struct {
	//
	//NotificationId string,
	//
	//RecipientId string,
	//
	//Channel string,
	//
	//SentAt string,
	//
	// TODO
}

type NotificationSentHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewNotificationSentHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) NotificationSentHandler {
	return NotificationSentHandler{
		publisher: publisher,
	}
}

func (h NotificationSentHandler) Handle(ctx context.Context, event NotificationSent) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.event.notification_sent.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("notificationsandcommunications.notification_sent", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
