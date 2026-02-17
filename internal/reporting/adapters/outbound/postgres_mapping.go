package outbound

import (
	"github.com/cooperlutz/go-full/internal/reporting/app/query"
	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
)

// toDomain maps the ReportingMetric to the domain entity.
func (m ReportingMetric) toDomain() (*reporting.Metric, error) {
	return reporting.MapToMetric(
		m.Name,
		m.Value,
	)
}

// mapEntityMetricToDB maps a domain Metric entity to the ReportingMetric database model.
func mapEntityMetricToDb(metric *reporting.Metric) ReportingMetric {
	return ReportingMetric{
		Name:  metric.GetName(),
		Value: metric.GetValue(),
	}
}

// mapEntityMetricToQuery maps a domain Metric entity to the query ReportMetric used in the application layer.
func mapEntityMetricToQuery(metric *reporting.Metric) query.ReportMetric {
	return query.ReportMetric{
		Name:  metric.GetName(),
		Value: metric.GetValue(),
	}
}
