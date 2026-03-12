//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type SalesOrderFulfilled struct {
	//
	//OrderId string,
	//
	//OwnerId string,
	//
	// TODO
}

type SalesOrderFulfilledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewSalesOrderFulfilledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) SalesOrderFulfilledHandler {
	return SalesOrderFulfilledHandler{
		publisher: publisher,
	}
}

func (h SalesOrderFulfilledHandler) Handle(ctx context.Context, event SalesOrderFulfilled) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.sales_order_fulfilled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.sales_order_fulfilled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
