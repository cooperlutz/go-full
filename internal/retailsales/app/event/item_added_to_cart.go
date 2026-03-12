//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ItemAddedToCart struct {
	//
	//CartId string,
	//
	//ProductId string,
	//
	//Quantity int32,
	//
	// TODO
}

type ItemAddedToCartHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewItemAddedToCartHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ItemAddedToCartHandler {
	return ItemAddedToCartHandler{
		publisher: publisher,
	}
}

func (h ItemAddedToCartHandler) Handle(ctx context.Context, event ItemAddedToCart) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.item_added_to_cart.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.item_added_to_cart", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
