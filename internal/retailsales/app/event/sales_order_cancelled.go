//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type SalesOrderCancelled struct {
	//
	//OrderId string,
	//
	//Reason string,
	//
	// TODO
}

type SalesOrderCancelledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewSalesOrderCancelledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) SalesOrderCancelledHandler {
	return SalesOrderCancelledHandler{
		publisher: publisher,
	}
}

func (h SalesOrderCancelledHandler) Handle(ctx context.Context, event SalesOrderCancelled) error {
	ctx, span := telemetree.AddSpan(ctx, "retailsales.app.event.sales_order_cancelled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("retailsales.sales_order_canceled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
