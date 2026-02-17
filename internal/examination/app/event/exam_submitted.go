package event

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamSubmitted struct {
	ExamId            string
	LibraryExamId     string
	StudentId         string
	Completed         bool
	AnsweredQuestions int32
	TotalQuestions    int32
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

	return nil
}
