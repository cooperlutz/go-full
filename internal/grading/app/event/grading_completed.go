//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type GradingCompleted struct {
	ExamId string `json:"examId"`
}

type GradingCompletedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewGradingCompletedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) GradingCompletedHandler {
	return GradingCompletedHandler{
		publisher: publisher,
	}
}

func (h GradingCompletedHandler) Handle(ctx context.Context, event GradingCompleted) error {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.event.gradingcompleted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("grading.grading_completed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("reporting.grading_completed", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
