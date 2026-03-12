//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ServicePaymentCompleted struct {
	//
	//OwnerId string,
	//
	//AmountPaid float32,
	//
	//ServiceType string,
	//
	// TODO
}

type ServicePaymentCompletedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewServicePaymentCompletedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ServicePaymentCompletedHandler {
	return ServicePaymentCompletedHandler{
		publisher: publisher,
	}
}

func (h ServicePaymentCompletedHandler) Handle(ctx context.Context, event ServicePaymentCompleted) error {
	ctx, span := telemetree.AddSpan(ctx, "ownermanagement.app.event.service_payment_completed.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("ownermanagement.service_payment_completed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
