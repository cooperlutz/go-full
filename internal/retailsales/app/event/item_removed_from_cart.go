//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ItemRemovedFromCart struct {
	//
	//CartId string,
	//
	//ProductId string,
	//
	// TODO
}

type ItemRemovedFromCartHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewItemRemovedFromCartHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ItemRemovedFromCartHandler {
	return ItemRemovedFromCartHandler{
		publisher: publisher,
	}
}

func (h ItemRemovedFromCartHandler) Handle(ctx context.Context, event ItemRemovedFromCart) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.item_removed_from_cart.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.item_removed_from_cart", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
