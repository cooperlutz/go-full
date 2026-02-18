package reporting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	newReport := NewReport(
		"fifth grade report",
		"this report provides insights into fifth grade performance on the midterm exam",
		WithMetric(Metric{name: MetricNumberOfExamsInProgress, value: 12}),
		WithMetric(Metric{name: MetricNumberOfExamsCompleted, value: 45}),
	)
	assert.Equal(t, "fifth grade report", newReport.GetName())
}
