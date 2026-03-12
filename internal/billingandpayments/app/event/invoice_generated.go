//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InvoiceGenerated struct {
	//
	//InvoiceId string,
	//
	//OwnerId string,
	//
	//TotalAmount float32,
	//
	//DueDate string,
	//
	// TODO
}

type InvoiceGeneratedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInvoiceGeneratedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InvoiceGeneratedHandler {
	return InvoiceGeneratedHandler{
		publisher: publisher,
	}
}

func (h InvoiceGeneratedHandler) Handle(ctx context.Context, event InvoiceGenerated) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.event.invoice_generated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("billingandpayments.invoice_generated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
