//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ProductPriceUpdated struct {
	//
	//ProductId string,
	//
	//NewUnitPrice float32,
	//
	// TODO
}

type ProductPriceUpdatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewProductPriceUpdatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ProductPriceUpdatedHandler {
	return ProductPriceUpdatedHandler{
		publisher: publisher,
	}
}

func (h ProductPriceUpdatedHandler) Handle(ctx context.Context, event ProductPriceUpdated) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.product_price_updated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.product_price_updated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
