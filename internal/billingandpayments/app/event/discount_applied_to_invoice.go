//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type DiscountAppliedToInvoice struct {
	//
	//InvoiceId string,
	//
	//DiscountAmount float32,
	//
	// TODO
}

type DiscountAppliedToInvoiceHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewDiscountAppliedToInvoiceHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) DiscountAppliedToInvoiceHandler {
	return DiscountAppliedToInvoiceHandler{
		publisher: publisher,
	}
}

func (h DiscountAppliedToInvoiceHandler) Handle(ctx context.Context, event DiscountAppliedToInvoice) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.event.discount_applied_to_invoice.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("billingandpayments.discount_applied_to_invoice", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
