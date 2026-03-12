//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InventoryRestocked struct {
	//
	//ProductId string,
	//
	//QuantityAdded int32,
	//
	//NewQuantityOnHand int,
	//
	// TODO
}

type InventoryRestockedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInventoryRestockedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InventoryRestockedHandler {
	return InventoryRestockedHandler{
		publisher: publisher,
	}
}

func (h InventoryRestockedHandler) Handle(ctx context.Context, event InventoryRestocked) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.inventory_restocked.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.inventory_restocked", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
