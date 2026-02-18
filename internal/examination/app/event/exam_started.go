package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamStarted struct {
	ExamID    string
	StudentID string
}

type ExamStartedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewExamStartedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ExamStartedHandler {
	return ExamStartedHandler{
		publisher: publisher,
	}
}

func (h ExamStartedHandler) Handle(ctx context.Context, event ExamStarted) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.event.exam_started.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("examination.exam_started", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("reporting.exam_started", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
