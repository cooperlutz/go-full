package query

import (
	"context"

	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindMetric struct {
	Name string
}

type FindMetricHandler struct {
	readModel FindMetricReadModel
}

func NewFindMetricHandler(
	readModel FindMetricReadModel,
) FindMetricHandler {
	return FindMetricHandler{readModel: readModel}
}

type FindMetricReadModel interface {
	FindMetric(ctx context.Context, name reporting.MetricName) (ReportMetric, error)
}

func (h FindMetricHandler) Handle(ctx context.Context, qry FindMetric) (ReportMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "reporting.app.query.findmetric.handle")
	defer span.End()

	metricName, err := reporting.MetricNameFromString(qry.Name)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return ReportMetric{}, err
	}

	r, err := h.readModel.FindMetric(ctx, metricName)
	if err != nil {
		return ReportMetric{}, err
	}

	return r, nil
}
