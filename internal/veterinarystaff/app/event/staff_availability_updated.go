//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type StaffAvailabilityUpdated struct {
	//
	//StaffId string,
	//
	//DayOfWeek string,
	//
	//IsAvailable bool,
	//
	// TODO
}

type StaffAvailabilityUpdatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewStaffAvailabilityUpdatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) StaffAvailabilityUpdatedHandler {
	return StaffAvailabilityUpdatedHandler{
		publisher: publisher,
	}
}

func (h StaffAvailabilityUpdatedHandler) Handle(ctx context.Context, event StaffAvailabilityUpdated) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.event.staff_availability_updated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("veterinarystaff.staff_availability_updated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
