package reporting

import (
	"context"
)

type Repository interface {
	// Reports

	AddReport(ctx context.Context, report *Report) error

	GetReport(ctx context.Context, name string) (*Report, error)

	UpdateReport(
		ctx context.Context,
		name string,
		updateFn func(h *Report) (*Report, error),
	) error

	// Metrics

	AddMetric(ctx context.Context, metric *Metric) error

	GetMetric(ctx context.Context, name MetricName) (*Metric, error)

	UpdateMetric(
		ctx context.Context,
		name MetricName,
		updateFn func(h *Metric) (*Metric, error),
	) error
}

func MapToMetric(
	name string,
	value float64,
) (*Metric, error) {
	metricName, err := MetricNameFromString(name)
	if err != nil {
		return nil, err
	}

	return &Metric{
		name:  metricName,
		value: value,
	}, nil
}

func MapToReport(
	name string,
	description string,
	components ...*Component,
) (*Report, error) {
	return &Report{
		name:        name,
		description: description,
		components:  components,
	}, nil
}
