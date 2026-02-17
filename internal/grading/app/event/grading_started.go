//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type GradingStarted struct {
	ExamId string `json:"examId"`
}

type GradingStartedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewGradingStartedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) GradingStartedHandler {
	return GradingStartedHandler{
		publisher: publisher,
	}
}

func (h GradingStartedHandler) Handle(ctx context.Context, event GradingStarted) error {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.event.grading_started.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("grading.grading_started", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
