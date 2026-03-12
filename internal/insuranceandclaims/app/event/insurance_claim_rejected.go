//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InsuranceClaimRejected struct {
	//
	//ClaimId string,
	//
	//OwnerId string,
	//
	//Reason string,
	//
	// TODO
}

type InsuranceClaimRejectedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInsuranceClaimRejectedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InsuranceClaimRejectedHandler {
	return InsuranceClaimRejectedHandler{
		publisher: publisher,
	}
}

func (h InsuranceClaimRejectedHandler) Handle(ctx context.Context, event InsuranceClaimRejected) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.event.insurance_claim_rejected.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("insuranceandclaims.insurance_claim_rejected", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
