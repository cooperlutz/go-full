//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type SalesOrderPlaced struct {
	//
	//OrderId string,
	//
	//OwnerId string,
	//
	//TotalAmount float32,
	//
	//Channel string,
	//
	// TODO
}

type SalesOrderPlacedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewSalesOrderPlacedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) SalesOrderPlacedHandler {
	return SalesOrderPlacedHandler{
		publisher: publisher,
	}
}

func (h SalesOrderPlacedHandler) Handle(ctx context.Context, event SalesOrderPlaced) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.sales_order_placed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.sales_order_placed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
