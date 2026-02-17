package reporting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetric(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsCompleted, 50)
	assert.Equal(t, MetricNumberOfExamsCompleted, metric.name)
	assert.Equal(t, 50.0, metric.value)
}

func TestMetric_GetName(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsBeingGraded, 5)
	assert.Equal(t, "number_of_exams_being_graded", metric.GetName())
}

func TestMetric_GetValue(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsGradingCompleted, 45)
	assert.Equal(t, 45.0, metric.GetValue())
}

func TestMetric_SetValue(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsInProgress, 10)
	metric.SetValue(20)
	assert.Equal(t, 20.0, metric.GetValue())
}

func TestMetric_IncrementValue(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsInProgress, 10)
	metric.IncrementValue(5)
	assert.Equal(t, 15.0, metric.GetValue())
}

func TestMetric_IncrementValueByOne(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsInProgress, 10)
	metric.IncrementValueByOne()
	assert.Equal(t, 11.0, metric.GetValue())
}

func TestMetric_DecrementValue(t *testing.T) {
	metric := NewMetric(MetricNumberOfExamsInProgress, 10)
	metric.DecrementValue(3)
	assert.Equal(t, 7.0, metric.GetValue())
}

func TestMetricNameFromString_ValidName(t *testing.T) {
	metricName, err := MetricNameFromString("number_of_exams_completed")
	assert.NoError(t, err)
	assert.Equal(t, MetricNumberOfExamsCompleted, metricName)
}

func TestMetricNameFromString_EmptyName(t *testing.T) {
	_, err := MetricNameFromString("")
	assert.Error(t, err)
	assert.IsType(t, ErrInvalidMetricName{}, err)
	assert.Equal(t, "invalid metric name", err.Error())
}

func TestMetricNameFromString_InvalidName(t *testing.T) {
	_, err := MetricNameFromString("invalid_metric_name")
	assert.Error(t, err)
	assert.IsType(t, ErrInvalidMetricName{}, err)
	assert.Equal(t, "invalid metric name", err.Error())
}
