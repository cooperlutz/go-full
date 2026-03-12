//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type OwnerEnrolledInLoyaltyProgram struct {
	//
	//OwnerId string,
	//
	//LoyaltyAccountId string,
	//
	//EnrolledDate string,
	//
	// TODO
}

type OwnerEnrolledInLoyaltyProgramHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewOwnerEnrolledInLoyaltyProgramHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) OwnerEnrolledInLoyaltyProgramHandler {
	return OwnerEnrolledInLoyaltyProgramHandler{
		publisher: publisher,
	}
}

func (h OwnerEnrolledInLoyaltyProgramHandler) Handle(ctx context.Context, event OwnerEnrolledInLoyaltyProgram) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.owner_enrolled_in_loyalty_program.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.owner_enrolled_in_loyalty_program", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
