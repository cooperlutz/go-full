//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type OwnerProfileUpdated struct {
	//
	//OwnerId string,
	//
	//UpdatedFields string,
	//
	// TODO
}

type OwnerProfileUpdatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewOwnerProfileUpdatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) OwnerProfileUpdatedHandler {
	return OwnerProfileUpdatedHandler{
		publisher: publisher,
	}
}

func (h OwnerProfileUpdatedHandler) Handle(ctx context.Context, event OwnerProfileUpdated) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.owner_profile_updated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.owner_profile_updated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
