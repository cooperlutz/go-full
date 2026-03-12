//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type MedicalRecordAdded struct {
	//
	//MedicalRecordId string,
	//
	//PetId string,
	//
	//VisitDate string,
	//
	//FollowUpRequired bool,
	//
	// TODO
}

type MedicalRecordAddedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewMedicalRecordAddedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) MedicalRecordAddedHandler {
	return MedicalRecordAddedHandler{
		publisher: publisher,
	}
}

func (h MedicalRecordAddedHandler) Handle(ctx context.Context, event MedicalRecordAdded) error {
	ctx, span := telemetree.AddSpan(ctx, "patientmanagement.app.event.medical_record_added.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("patientmanagement.medical_record_added", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
