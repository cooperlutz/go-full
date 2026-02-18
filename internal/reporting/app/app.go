package app

import (
	"github.com/cooperlutz/go-full/internal/reporting/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/reporting/app/event"
	"github.com/cooperlutz/go-full/internal/reporting/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Queries Queries
	Events  Events
}

type Queries struct {
	FindReport query.FindReportHandler
	FindMetric query.FindMetricHandler
}

type Events struct {
	ExamStarted      event.ExamStartedHandler
	ExamSubmitted    event.ExamSubmittedHandler
	GradingStarted   event.GradingStartedHandler
	GradingCompleted event.GradingCompletedHandler
}

// NewApplication initializes the Examination application with its dependencies.
func NewApplication(
	pgConn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	reportingRepository := outbound.NewPostgresAdapter(
		pgConn,
	)

	app := Application{
		Queries: Queries{
			FindReport: query.NewFindReportHandler(
				reportingRepository,
			),
			FindMetric: query.NewFindMetricHandler(
				reportingRepository,
			),
		},
		Events: Events{
			ExamStarted: event.NewExamStartedHandler(
				reportingRepository,
			),
			ExamSubmitted: event.NewExamSubmittedHandler(
				reportingRepository,
			),
			GradingStarted: event.NewGradingStartedHandler(
				reportingRepository,
			),
			GradingCompleted: event.NewGradingCompletedHandler(
				reportingRepository,
			),
		},
	}

	return app, nil
}
