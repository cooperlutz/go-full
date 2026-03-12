//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type OwnerRegistered struct {
	//
	//OwnerId string,
	//
	//FirstName string,
	//
	//LastName string,
	//
	//Email string,
	//
	// TODO
}

type OwnerRegisteredHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewOwnerRegisteredHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) OwnerRegisteredHandler {
	return OwnerRegisteredHandler{
		publisher: publisher,
	}
}

func (h OwnerRegisteredHandler) Handle(ctx context.Context, event OwnerRegistered) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.owner_registered.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.owner_registered", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
