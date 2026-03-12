//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PetRegistered struct {
	//
	//PetId string,
	//
	//OwnerId string,
	//
	//Name string,
	//
	//Species string,
	//
	// TODO
}

type PetRegisteredHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPetRegisteredHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PetRegisteredHandler {
	return PetRegisteredHandler{
		publisher: publisher,
	}
}

func (h PetRegisteredHandler) Handle(ctx context.Context, event PetRegistered) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.event.pet_registered.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("patientmanagement.pet_registered", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
