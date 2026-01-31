//nolint:dupl // basic publish logic will be duplicative, but we want to be explicit for each event
package event

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamSubmitted struct {
	ExamID    string
	StudentID string
}

type ExamSubmittedHandler struct {
	publisher outbound.SqlPublisherAdapter
}

func NewExamSubmittedHandler(
	publisher outbound.SqlPublisherAdapter,
) ExamSubmittedHandler {
	return ExamSubmittedHandler{
		publisher: publisher,
	}
}

func (h ExamSubmittedHandler) Handle(ctx context.Context, event ExamSubmitted) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.event.examsubmitted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.Publish("examination.examsubmitted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
