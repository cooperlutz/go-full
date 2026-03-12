//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type LoyaltyPointsAwarded struct {
	//
	//OwnerId string,
	//
	//Points int32,
	//
	//NewBalance int32,
	//
	// TODO
}

type LoyaltyPointsAwardedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewLoyaltyPointsAwardedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) LoyaltyPointsAwardedHandler {
	return LoyaltyPointsAwardedHandler{
		publisher: publisher,
	}
}

func (h LoyaltyPointsAwardedHandler) Handle(ctx context.Context, event LoyaltyPointsAwarded) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.loyalty_points_awarded.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.loyalty_points_awarded", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
