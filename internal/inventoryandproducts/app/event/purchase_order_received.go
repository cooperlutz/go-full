//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PurchaseOrderReceived struct {
	//
	//PurchaseOrderId string,
	//
	//ReceivedDate string,
	//
	// TODO
}

type PurchaseOrderReceivedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPurchaseOrderReceivedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PurchaseOrderReceivedHandler {
	return PurchaseOrderReceivedHandler{
		publisher: publisher,
	}
}

func (h PurchaseOrderReceivedHandler) Handle(ctx context.Context, event PurchaseOrderReceived) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.purchase_order_received.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.purchase_order_received", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
