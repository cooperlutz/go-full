package reporting

// MetricName represents the name of a metric.
// It is defined as an enumeration.
type MetricName int

const (
	// MetricNumberOfExamsInProgress represents the number of exams that are currently in progress. this metric is event-driven and should be updated whenever an exam is started or completed.
	MetricNumberOfExamsInProgress MetricName = iota
	// MetricNumberOfExamsCompleted represents the number of exams that have been completed. this metric is event-driven and should be updated whenever an exam is completed.
	MetricNumberOfExamsCompleted
	// MetricNumberOfExamsBeingGraded represents the number of exams that are currently being graded. this metric is event-driven and should be updated whenever an exam starts or finishes grading.
	MetricNumberOfExamsBeingGraded
	// MetricNumberOfExamsGradingCompleted represents the number of exams that have completed grading. this metric is event-driven and should be updated whenever an exam finishes grading.
	MetricNumberOfExamsGradingCompleted
)

var metricName = map[MetricName]string{ //nolint:gochecknoglobals // global is ok here
	MetricNumberOfExamsInProgress:       "number_of_exams_in_progress",
	MetricNumberOfExamsCompleted:        "number_of_exams_completed",
	MetricNumberOfExamsBeingGraded:      "number_of_exams_being_graded",
	MetricNumberOfExamsGradingCompleted: "number_of_exams_grading_completed",
}

func (mn MetricName) String() string {
	return metricName[mn]
}

type ErrInvalidMetricName struct{}

func (e ErrInvalidMetricName) Error() string {
	return "invalid metric name"
}

func MetricNameFromString(s string) (MetricName, error) {
	for mn, name := range metricName {
		if name == s {
			return mn, nil
		}
	}

	return 0, ErrInvalidMetricName{}
}

type Metric struct {
	name  MetricName
	value float64
}

func NewMetric(name MetricName, value float64) *Metric {
	return &Metric{
		name:  name,
		value: value,
	}
}

func (m *Metric) GetName() string {
	return m.name.String()
}

func (m *Metric) GetValue() float64 {
	return m.value
}

func (m *Metric) SetValue(value float64) {
	m.value = value
}

func (m *Metric) IncrementValue(delta float64) {
	m.value += delta
}

func (m *Metric) IncrementValueByOne() {
	m.value += 1
}

func (m *Metric) DecrementValue(delta float64) {
	m.value -= delta
}

func (m *Metric) DecrementValueByOne() {
	m.value -= 1
}
