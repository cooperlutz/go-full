//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type LowStockAlertTriggered struct {
	//
	//ProductId string,
	//
	//CurrentQuantity int32,
	//
	//ReorderThreshold int32,
	//
	// TODO
}

type LowStockAlertTriggeredHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewLowStockAlertTriggeredHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) LowStockAlertTriggeredHandler {
	return LowStockAlertTriggeredHandler{
		publisher: publisher,
	}
}

func (h LowStockAlertTriggeredHandler) Handle(ctx context.Context, event LowStockAlertTriggered) error {
	ctx, span := telemetree.AddSpan(ctx, "inventoryandproducts.app.event.low_stock_alert_triggered.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("inventoryandproducts.low_stock_alert_triggered", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
