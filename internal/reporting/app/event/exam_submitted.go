//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
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
	reportingRepo reporting.Repository
}

func NewExamSubmittedHandler(
	repo reporting.Repository,
) ExamSubmittedHandler {
	return ExamSubmittedHandler{
		reportingRepo: repo,
	}
}

func (h ExamSubmittedHandler) Handle() message.NoPublishHandlerFunc {
	return eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
		ctx, span := telemetree.AddSpan(msg.Context(), "reporting.app.event.examsubmitted.handle")
		defer span.End()

		var event ExamSubmitted

		err := json.Unmarshal(msg.Payload, &event)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.reportingRepo.UpdateMetric(ctx, reporting.MetricNumberOfExamsInProgress, func(m *reporting.Metric) (*reporting.Metric, error) {
			m.DecrementValueByOne()

			return m, nil
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.reportingRepo.UpdateMetric(ctx, reporting.MetricNumberOfExamsCompleted, func(m *reporting.Metric) (*reporting.Metric, error) {
			m.IncrementValueByOne()

			return m, nil
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		return nil
	})
}
