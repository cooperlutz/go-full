package app

import (
	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/app/command"
	"github.com/cooperlutz/go-full/internal/examination/app/event"
	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	StartExam      command.StartExamHandler
	AnswerQuestion command.AnswerQuestionHandler
	SubmitExam     command.SubmitExamHandler
}

type Queries struct {
	FindAllExams query.AllExamsHandler
	FindQuestion query.FindQuestionHandler
	FindExam     query.FindExamHandler
}

type Events struct {
	ExamStarted   event.ExamStartedHandler
	ExamSubmitted event.ExamSubmittedHandler
}

// NewApplication initializes the Examination application with its dependencies.
func NewApplication(
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
	examLibraryUseCase usecase.IExamLibraryUseCase,
) (Application, error) {
	examinationRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			StartExam: command.NewStartExamHandler(
				examinationRepository,
				outbound.NewExamLibraryAdapter(
					examLibraryUseCase,
				),
			),
			AnswerQuestion: command.NewAnswerQuestionHandler(
				examinationRepository,
				outbound.NewExamLibraryAdapter(
					examLibraryUseCase,
				),
			),
			SubmitExam: command.NewSubmitExamHandler(
				examinationRepository,
				outbound.NewExamLibraryAdapter(
					examLibraryUseCase,
				),
			),
		},
		Queries: Queries{
			FindAllExams: query.NewFindAllExamsHandler(
				examinationRepository,
			),
			FindQuestion: query.NewFindQuestionHandler(
				examinationRepository,
			),
			FindExam: query.NewFindExamHandler(
				examinationRepository,
			),
		},
		Events: Events{
			ExamStarted: event.NewExamStartedHandler(
				pubSub,
			),
			ExamSubmitted: event.NewExamSubmittedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
