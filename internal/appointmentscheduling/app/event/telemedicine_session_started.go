//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type TelemedicineSessionStarted struct {
	//
	//SessionId string,
	//
	//AppointmentId string,
	//
	//SessionUrl string,
	//
	// TODO
}

type TelemedicineSessionStartedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewTelemedicineSessionStartedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) TelemedicineSessionStartedHandler {
	return TelemedicineSessionStartedHandler{
		publisher: publisher,
	}
}

func (h TelemedicineSessionStartedHandler) Handle(ctx context.Context, event TelemedicineSessionStarted) error {
	ctx, span := telemetree.AddSpan(ctx, "appointmentscheduling.app.event.telemedicine_session_started.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("appointmentscheduling.telemedicine_session_started", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
