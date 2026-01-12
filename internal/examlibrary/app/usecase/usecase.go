package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
)

// IExamLibraryUseCase is the interface that describes the examlibrary usecases.
type IExamLibraryUseCase interface {
	AddExamToLibrary(ctx context.Context, cmd command.AddExamToLibrary) (command.AddExamToLibraryResult, error)
	FindOneExamByID(ctx context.Context, qry query.FindOneExamByID) (query.FindOneExamByIDResponse, error)
	FindAllExamsWithoutQuestions(ctx context.Context, qry query.FindAllExamsWithoutQuestions) (query.FindAllExamsWithoutQuestionsResponse, error)
	FindAllExamsWithQuestions(ctx context.Context, qry query.FindAllExamsWithQuestions) (query.FindAllExamsWithQuestionsResponse, error)
}
