//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type VeterinarianProfileUpdated struct {
	//
	//VeterinarianId string,
	//
	//UpdatedFields string,
	//
	// TODO
}

type VeterinarianProfileUpdatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewVeterinarianProfileUpdatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) VeterinarianProfileUpdatedHandler {
	return VeterinarianProfileUpdatedHandler{
		publisher: publisher,
	}
}

func (h VeterinarianProfileUpdatedHandler) Handle(ctx context.Context, event VeterinarianProfileUpdated) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.event.veterinarian_profile_updated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("veterinarystaff.veterinarian_profile_updated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
