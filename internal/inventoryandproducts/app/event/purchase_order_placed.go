//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PurchaseOrderPlaced struct {
	//
	//PurchaseOrderId string,
	//
	//SupplierId string,
	//
	//TotalCost float32,
	//
	// TODO
}

type PurchaseOrderPlacedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPurchaseOrderPlacedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PurchaseOrderPlacedHandler {
	return PurchaseOrderPlacedHandler{
		publisher: publisher,
	}
}

func (h PurchaseOrderPlacedHandler) Handle(ctx context.Context, event PurchaseOrderPlaced) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.purchase_order_placed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.purchase_order_placed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
