package event

import (
	"context"
	"time"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamSubmitted struct {
	ExamId            string
	LibraryExamId     string
	StudentId         string
	ExamState         string
	AnsweredQuestions int32
	TotalQuestions    int32
	TimeLimitSeconds  int64
	TimeOfTimeLimit   time.Time
	StartedAt         time.Time
	CompletedAt       time.Time
	Questions         []ExamSubmittedQuestion
}

type ExamSubmittedQuestion struct {
	ExamId          string
	Answered        bool
	QuestionID      string
	QuestionIndex   int32
	QuestionText    string
	QuestionType    string
	ResponseOptions *[]string
	ProvidedAnswer  *string
}

type ExamSubmittedHandler struct {
	publisher eeventdriven.IPubSubEventProcessor
}

func NewExamSubmittedHandler(
	publisher eeventdriven.IPubSubEventProcessor,
) ExamSubmittedHandler {
	return ExamSubmittedHandler{
		publisher: publisher,
	}
}

func (h ExamSubmittedHandler) Handle(ctx context.Context, event ExamSubmitted) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.event.exam_submitted.handle")
	defer span.End()

	msg, err := eeventdriven.EventPayloadToMessage(event)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("examination.exam_submitted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("reporting.exam_submitted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.publisher.EmitEventMessage("grading.exam_submitted", msg)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
