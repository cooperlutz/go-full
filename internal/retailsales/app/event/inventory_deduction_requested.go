//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InventoryDeductionRequested struct {
	//
	//OrderId string,
	//
	//LineItems string,
	//
	// TODO
}

type InventoryDeductionRequestedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInventoryDeductionRequestedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InventoryDeductionRequestedHandler {
	return InventoryDeductionRequestedHandler{
		publisher: publisher,
	}
}

func (h InventoryDeductionRequestedHandler) Handle(ctx context.Context, event InventoryDeductionRequested) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.inventory_deduction_requested.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.inventory_deduction_requested", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
