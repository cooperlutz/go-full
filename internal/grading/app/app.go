package app

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/grading/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/grading/app/command"
	"github.com/cooperlutz/go-full/internal/grading/app/event"
	"github.com/cooperlutz/go-full/internal/grading/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	GradeQuestion command.GradeQuestionHandler
}

type Queries struct {
	FindExam         query.FindExamHandler
	FindExamQuestion query.FindExamQuestionHandler
	IncompleteExams  query.IncompleteExamsHandler
}

type Events struct {
	ExamSubmitted    event.ExamSubmittedHandler
	GradingStarted   event.GradingStartedHandler
	GradingCompleted event.GradingCompletedHandler
}

// NewApplication initializes the Grading application with its dependencies.
func NewApplication(
	pgconn deebee.IDatabase,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
	examLibraryUseCase usecase.IExamLibraryUseCase,
) (Application, error) {
	gradingRepo := outbound.NewPostgresAdapter(pgconn)

	app := Application{
		Commands: Commands{
			GradeQuestion: command.NewGradeQuestionHandler(gradingRepo),
		},
		Queries: Queries{
			FindExam:         query.NewFindExamHandler(gradingRepo),
			FindExamQuestion: query.NewFindExamQuestionHandler(gradingRepo),
			IncompleteExams:  query.NewIncompleteExamsHandler(gradingRepo),
		},
		Events: Events{
			ExamSubmitted:    event.NewExamSubmittedHandler(gradingRepo, examLibraryUseCase),
			GradingStarted:   event.NewGradingStartedHandler(pubSub),
			GradingCompleted: event.NewGradingCompletedHandler(pubSub),
		},
	}

	return app, nil
}
