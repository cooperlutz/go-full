//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type TelemedicineSessionEnded struct {
	//
	//SessionId string,
	//
	//AppointmentId string,
	//
	//EndedAt string,
	//
	// TODO
}

type TelemedicineSessionEndedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewTelemedicineSessionEndedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) TelemedicineSessionEndedHandler {
	return TelemedicineSessionEndedHandler{
		publisher: publisher,
	}
}

func (h TelemedicineSessionEndedHandler) Handle(ctx context.Context, event TelemedicineSessionEnded) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.telemedicine_session_ended.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.telemedicine_session_ended", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
