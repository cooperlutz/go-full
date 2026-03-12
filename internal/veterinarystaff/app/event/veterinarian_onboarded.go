//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type VeterinarianOnboarded struct {
	//
	//VeterinarianId string,
	//
	//FirstName string,
	//
	//LastName string,
	//
	//LicenseNumber string,
	//
	// TODO
}

type VeterinarianOnboardedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewVeterinarianOnboardedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) VeterinarianOnboardedHandler {
	return VeterinarianOnboardedHandler{
		publisher: publisher,
	}
}

func (h VeterinarianOnboardedHandler) Handle(ctx context.Context, event VeterinarianOnboarded) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.event.veterinarian_onboarded.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("veterinarystaff.veterinarian_onboarded", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
