//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InsuranceProviderRegistered struct {
	//
	//ProviderId string,
	//
	//Name string,
	//
	// TODO
}

type InsuranceProviderRegisteredHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInsuranceProviderRegisteredHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InsuranceProviderRegisteredHandler {
	return InsuranceProviderRegisteredHandler{
		publisher: publisher,
	}
}

func (h InsuranceProviderRegisteredHandler) Handle(ctx context.Context, event InsuranceProviderRegistered) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.event.insurance_provider_registered.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("insuranceandclaims.insurance_provider_registered", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
