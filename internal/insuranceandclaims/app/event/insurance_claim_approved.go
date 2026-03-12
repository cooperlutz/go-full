//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InsuranceClaimApproved struct {
	//
	//ClaimId string,
	//
	//OwnerId string,
	//
	//ApprovedAmount float32,
	//
	//InvoiceId string,
	//
	// TODO
}

type InsuranceClaimApprovedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInsuranceClaimApprovedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InsuranceClaimApprovedHandler {
	return InsuranceClaimApprovedHandler{
		publisher: publisher,
	}
}

func (h InsuranceClaimApprovedHandler) Handle(ctx context.Context, event InsuranceClaimApproved) error {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.event.insurance_claim_approved.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("insuranceandclaims.insurance_claim_approved", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
