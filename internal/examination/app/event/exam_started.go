//nolint:dupl // basic publish logic will be duplicative, but we want to be explicit for each event
package event

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamStarted struct {
	ExamID    string
	StudentID string
}

type ExamStartedHandler struct {
	publisher outbound.SqlPublisherAdapter
}

func NewExamStartedHandler(
	publisher outbound.SqlPublisherAdapter,
) ExamStartedHandler {
	return ExamStartedHandler{
		publisher: publisher,
	}
}

func (h ExamStartedHandler) Handle(ctx context.Context, event ExamStarted) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.event.examstarted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.Publish("examination.examstarted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
