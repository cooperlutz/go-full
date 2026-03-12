//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AppointmentCancelled struct {
	//
	//AppointmentId string,
	//
	//Reason string,
	//
	//CancelledBy string,
	//
	// TODO
}

type AppointmentCancelledHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewAppointmentCancelledHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) AppointmentCancelledHandler {
	return AppointmentCancelledHandler{
		publisher: publisher,
	}
}

func (h AppointmentCancelledHandler) Handle(ctx context.Context, event AppointmentCancelled) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.appointment_cancelled.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.appointment_canceled", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
