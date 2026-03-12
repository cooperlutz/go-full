//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PetDetailsUpdated struct {
	//
	//PetId string,
	//
	//UpdatedFields string,
	//
	// TODO
}

type PetDetailsUpdatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPetDetailsUpdatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PetDetailsUpdatedHandler {
	return PetDetailsUpdatedHandler{
		publisher: publisher,
	}
}

func (h PetDetailsUpdatedHandler) Handle(ctx context.Context, event PetDetailsUpdated) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.event.pet_details_updated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("patientmanagement.pet_details_updated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
