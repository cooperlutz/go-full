//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type InvoiceVoided struct {
	//
	//InvoiceId string,
	//
	//Reason string,
	//
	// TODO
}

type InvoiceVoidedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewInvoiceVoidedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) InvoiceVoidedHandler {
	return InvoiceVoidedHandler{
		publisher: publisher,
	}
}

func (h InvoiceVoidedHandler) Handle(ctx context.Context, event InvoiceVoided) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.event.invoice_voided.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("billingandpayments.invoice_voided", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
