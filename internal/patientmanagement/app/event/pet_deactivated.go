//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PetDeactivated struct {
	//
	//PetId string,
	//
	//Reason string,
	//
	// TODO
}

type PetDeactivatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPetDeactivatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PetDeactivatedHandler {
	return PetDeactivatedHandler{
		publisher: publisher,
	}
}

func (h PetDeactivatedHandler) Handle(ctx context.Context, event PetDeactivated) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.event.pet_deactivated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("patientmanagement.pet_deactivated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
