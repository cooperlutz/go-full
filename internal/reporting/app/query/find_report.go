package query

import (
	"context"

	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindReport struct {
	Name string
}

type FindReportHandler struct {
	readModel FindReportReadModel
}

func NewFindReportHandler(
	readModel FindReportReadModel,
) FindReportHandler {
	return FindReportHandler{readModel: readModel}
}

type FindReportReadModel interface {
	FindReport(ctx context.Context, name string) (ReportDefinition, error)
	FindMetric(ctx context.Context, name reporting.MetricName) (ReportMetric, error)
}

type ReportDefinition struct {
	Name        string
	Description string
	Components  []ReportComponentDefinition
}

type ReportComponentDefinition struct {
	Type string
	Key  string
}

type Report struct {
	Name        string
	Description string
	Components  []ReportComponent
}

type ReportComponent struct {
	Type  string
	Key   string
	Value any
}

type ReportMetric struct {
	Name  string
	Value float64
}

func (h FindReportHandler) Handle(ctx context.Context, qry FindReport) (Report, error) {
	ctx, span := telemetree.AddSpan(ctx, "reporting.app.query.findreport.handle")
	defer span.End()

	r, err := h.readModel.FindReport(ctx, qry.Name)
	if err != nil {
		return Report{}, err
	}

	var components []ReportComponent

	for _, component := range r.Components {
		switch component.Type {
		case "metric":
			metricName, err := reporting.MetricNameFromString(component.Key)
			if err != nil {
				telemetree.RecordError(ctx, err)

				return Report{}, err
			}

			m, err := h.readModel.FindMetric(ctx, metricName)
			if err != nil {
				telemetree.RecordError(ctx, err)

				return Report{}, err
			}

			components = append(components, ReportComponent{
				Type:  "metric",
				Key:   component.Key,
				Value: m.Value,
			})
		}
	}

	return Report{
		Name:        r.Name,
		Description: r.Description,
		Components:  components,
	}, nil
}
