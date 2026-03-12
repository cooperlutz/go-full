//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type OwnerDeactivated struct {
	//
	//OwnerId string,
	//
	//Reason string,
	//
	// TODO
}

type OwnerDeactivatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewOwnerDeactivatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) OwnerDeactivatedHandler {
	return OwnerDeactivatedHandler{
		publisher: publisher,
	}
}

func (h OwnerDeactivatedHandler) Handle(ctx context.Context, event OwnerDeactivated) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.owner_deactivated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.owner_deactivated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
