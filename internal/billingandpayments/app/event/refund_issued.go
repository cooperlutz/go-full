//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type RefundIssued struct {
	//
	//RefundId string,
	//
	//PaymentId string,
	//
	//OwnerId string,
	//
	//RefundAmount float32,
	//
	// TODO
}

type RefundIssuedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewRefundIssuedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) RefundIssuedHandler {
	return RefundIssuedHandler{
		publisher: publisher,
	}
}

func (h RefundIssuedHandler) Handle(ctx context.Context, event RefundIssued) error {
	ctx, span := telemetree.AddSpan(ctx, "billingandpayments.app.event.refund_issued.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("billingandpayments.refund_issued", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
