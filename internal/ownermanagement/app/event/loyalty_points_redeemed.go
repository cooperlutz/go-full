//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type LoyaltyPointsRedeemed struct {
	//
	//OwnerId string,
	//
	//Points int32,
	//
	//NewBalance int32,
	//
	// TODO
}

type LoyaltyPointsRedeemedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewLoyaltyPointsRedeemedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) LoyaltyPointsRedeemedHandler {
	return LoyaltyPointsRedeemedHandler{
		publisher: publisher,
	}
}

func (h LoyaltyPointsRedeemedHandler) Handle(ctx context.Context, event LoyaltyPointsRedeemed) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.loyalty_points_redeemed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.loyalty_points_redeemed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
