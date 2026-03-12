//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type StaffMemberDeactivated struct {
	//
	//StaffId string,
	//
	//Reason string,
	//
	// TODO
}

type StaffMemberDeactivatedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewStaffMemberDeactivatedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) StaffMemberDeactivatedHandler {
	return StaffMemberDeactivatedHandler{
		publisher: publisher,
	}
}

func (h StaffMemberDeactivatedHandler) Handle(ctx context.Context, event StaffMemberDeactivated) error {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.event.staff_member_deactivated.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("veterinarystaff.staff_member_deactivated", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
