//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PaymentProcessed struct {
	//
	//PaymentId string,
	//
	//InvoiceId string,
	//
	//OwnerId string,
	//
	//AmountPaid float32,
	//
	// TODO
}

type PaymentProcessedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewPaymentProcessedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) PaymentProcessedHandler {
	return PaymentProcessedHandler{
		publisher: publisher,
	}
}

func (h PaymentProcessedHandler) Handle(ctx context.Context, event PaymentProcessed) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.event.payment_processed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("billingandpayments.payment_processed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
