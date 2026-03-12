//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InventoryDeducted struct {
	//
	//ProductId string,
	//
	//QuantityDeducted int32,
	//
	//NewQuantityOnHand int32,
	//
	// TODO
}

type InventoryDeductedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInventoryDeductedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InventoryDeductedHandler {
	return InventoryDeductedHandler{
		publisher: publisher,
	}
}

func (h InventoryDeductedHandler) Handle(ctx context.Context, event InventoryDeducted) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.inventory_deducted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.inventory_deducted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
