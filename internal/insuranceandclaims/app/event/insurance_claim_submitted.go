//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InsuranceClaimSubmitted struct {
	//
	//ClaimId string,
	//
	//OwnerId string,
	//
	//PetId string,
	//
	//ProviderId string,
	//
	//ClaimAmount float32,
	//
	// TODO
}

type InsuranceClaimSubmittedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInsuranceClaimSubmittedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InsuranceClaimSubmittedHandler {
	return InsuranceClaimSubmittedHandler{
		publisher: publisher,
	}
}

func (h InsuranceClaimSubmittedHandler) Handle(ctx context.Context, event InsuranceClaimSubmitted) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.event.insurance_claim_submitted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("insuranceandclaims.insurance_claim_submitted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
