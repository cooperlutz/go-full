//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type NotificationTemplateCreated struct {
	//
	//TemplateId string,
	//
	//Name string,
	//
	//NotificationType string,
	//
	// TODO
}

type NotificationTemplateCreatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewNotificationTemplateCreatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) NotificationTemplateCreatedHandler {
	return NotificationTemplateCreatedHandler{
		publisher: publisher,
	}
}

func (h NotificationTemplateCreatedHandler) Handle(ctx context.Context, event NotificationTemplateCreated) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.app.event.notification_template_created.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("notificationsandcommunications.notification_template_created", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
