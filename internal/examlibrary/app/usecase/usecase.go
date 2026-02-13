package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/repository"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

// IExamLibraryUseCase is the interface that describes the examlibrary usecases.
type IExamLibraryUseCase interface {
	AddExamToLibrary(ctx context.Context, cmd command.AddExamToLibrary) (command.AddExamToLibraryResult, error)
	FindOneExamByID(ctx context.Context, qry query.FindOneExamByID) (query.FindOneExamByIDResponse, error)
	FindAllExamsWithoutQuestions(ctx context.Context, qry query.FindAllExamsWithoutQuestions) (query.FindAllExamsWithoutQuestionsResponse, error)
	FindAllExamsWithQuestions(ctx context.Context, qry query.FindAllExamsWithQuestions) (query.FindAllExamsWithQuestionsResponse, error)
}

type examLibraryUseCase struct {
	Persist repository.IExamLibraryRepository
	Events  eeventdriven.IPubSubEventProcessor
}

// NewExamLibraryUseCase creates a new instance of the ExamLibraryUseCase.
func NewExamLibraryUseCase(repo repository.IExamLibraryRepository, events eeventdriven.IPubSubEventProcessor) *examLibraryUseCase {
	return &examLibraryUseCase{
		Persist: repo,
		Events:  events,
	}
}

// emitEvents emits the given domain events using the event processor.
func (uc *examLibraryUseCase) emitEvents(events []any) error {
	for _, ev := range events {
		err := uc.Events.EmitEvent("examlibrary", ev)
		if err != nil {
			return err
		}
	}

	return nil
}
