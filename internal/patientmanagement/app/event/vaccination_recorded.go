//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type VaccinationRecorded struct {
	//
	//VaccinationId string,
	//
	//PetId string,
	//
	//VaccineName string,
	//
	//ExpiryDate *string,
	//
	// TODO
}

type VaccinationRecordedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewVaccinationRecordedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) VaccinationRecordedHandler {
	return VaccinationRecordedHandler{
		publisher: publisher,
	}
}

func (h VaccinationRecordedHandler) Handle(ctx context.Context, event VaccinationRecorded) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.event.vaccination_recorded.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("patientmanagement.vaccination_recorded", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
