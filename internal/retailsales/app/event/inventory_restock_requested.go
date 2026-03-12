//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InventoryRestockRequested struct {
	//
	//OrderId string,
	//
	//LineItems string,
	//
	// TODO
}

type InventoryRestockRequestedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInventoryRestockRequestedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InventoryRestockRequestedHandler {
	return InventoryRestockRequestedHandler{
		publisher: publisher,
	}
}

func (h InventoryRestockRequestedHandler) Handle(ctx context.Context, event InventoryRestockRequested) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.inventory_restock_requested.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.inventory_restock_requested", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
