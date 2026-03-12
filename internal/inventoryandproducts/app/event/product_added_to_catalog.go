//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ProductAddedToCatalog struct {
	//
	//ProductId string,
	//
	//Name string,
	//
	//Category string,
	//
	//UnitPrice float32,
	//
	// TODO
}

type ProductAddedToCatalogHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewProductAddedToCatalogHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ProductAddedToCatalogHandler {
	return ProductAddedToCatalogHandler{
		publisher: publisher,
	}
}

func (h ProductAddedToCatalogHandler) Handle(ctx context.Context, event ProductAddedToCatalog) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.product_added_to_catalog.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.product_added_to_catalog", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
