package reporting

type Report struct {
	name        string
	description string
	components  []*Component
}

func NewReport(name, description string, components ...ReportComponent) *Report {
	report := &Report{
		name:        name,
		description: description,
	}

	for _, option := range components {
		component := new(Component)
		option(component)
		report.components = append(report.components, component)
	}

	return report
}

func (r *Report) GetName() string {
	return r.name
}

func (r *Report) GetDescription() string {
	return r.description
}

func (r *Report) GetComponents() []*Component {
	return r.components
}

type ReportComponent func(*Component)

type Component struct {
	name          string
	componentType string
	metrics       *Metric
}

func WithMetric(metric Metric) ReportComponent {
	return func(c *Component) {
		c.name = metric.name.String()
		c.componentType = "metric"
		c.metrics = &metric
	}
}
