//nolint:dupl // the event handlers look similar but we don't want to abstract them
package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type GradingCompleted struct {
	ExamId string `json:"examId"`
}

type GradingCompletedHandler struct {
	reportingRepo reporting.Repository
}

func NewGradingCompletedHandler(
	reportingRepo reporting.Repository,
) GradingCompletedHandler {
	return GradingCompletedHandler{
		reportingRepo: reportingRepo,
	}
}

func (h GradingCompletedHandler) Handle() message.NoPublishHandlerFunc {
	return eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
		ctx, span := telemetree.AddSpan(msg.Context(), "reporting.app.event.gradingcompleted.handle")
		defer span.End()

		var event GradingCompleted

		err := json.Unmarshal(msg.Payload, &event)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.reportingRepo.UpdateMetric(ctx, reporting.MetricNumberOfExamsBeingGraded, func(m *reporting.Metric) (*reporting.Metric, error) {
			m.DecrementValueByOne()

			return m, nil
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.reportingRepo.UpdateMetric(ctx, reporting.MetricNumberOfExamsGradingCompleted, func(m *reporting.Metric) (*reporting.Metric, error) {
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
