//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ExamStarted struct {
	ExamID    string
	StudentID string
}

type ExamStartedHandler struct {
	reportingRepo reporting.Repository
}

func NewExamStartedHandler(
	repo reporting.Repository,
) ExamStartedHandler {
	return ExamStartedHandler{
		reportingRepo: repo,
	}
}

func (h ExamStartedHandler) Handle() message.NoPublishHandlerFunc {
	return eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
		ctx, span := telemetree.AddSpan(msg.Context(), "reporting.app.event.examstarted.handle")
		defer span.End()

		var event ExamStarted

		err := json.Unmarshal(msg.Payload, &event)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.reportingRepo.UpdateMetric(ctx, reporting.MetricNumberOfExamsInProgress, func(m *reporting.Metric) (*reporting.Metric, error) {
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
