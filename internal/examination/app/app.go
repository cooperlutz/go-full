package app

import (
	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/app/command"
	"github.com/cooperlutz/go-full/internal/examination/app/event"
	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/deebee"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	StartExam command.StartExamHandler
}

type Queries struct {
	AvailableExams query.AvailableExamsHandler
}

type Events struct {
	ExamStarted event.ExamStartedHandler
	NoOp        event.NoOpEventHandler
}

// NewApplication initializes the Examination application with its dependencies.
func NewApplication(
	pgconn deebee.IDatabase,
	examLibraryUseCase usecase.IExamLibraryUseCase,
) (Application, error) {
	publisher, err := outbound.NewSqlPublisherAdapter(pgconn)
	if err != nil {
		return Application{}, err
	}

	examinationRepository := outbound.NewPostgresAdapter(pgconn)
	app := Application{
		Commands: Commands{
			StartExam: command.NewStartExamHandler(
				examinationRepository,
				outbound.NewExamLibraryAdapter(
					examLibraryUseCase,
				),
			),
		},
		Queries: Queries{
			AvailableExams: query.NewAvailableExamsHandler(
				examinationRepository,
			),
		},
		Events: Events{
			ExamStarted: event.NewExamStartedHandler(
				publisher,
			),
			NoOp: event.NewNoOpEventHandler(),
		},
	}

	return app, nil
}
