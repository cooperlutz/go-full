package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// FindAllExamsWithQuestions finds all exams with their questions.
func (uc *examLibraryUseCase) FindAllExamsWithQuestions(ctx context.Context, qry query.FindAllExamsWithQuestions) (query.FindAllExamsWithQuestionsResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.usecase.find_all_exams_with_questions")
	defer span.End()

	entity, err := uc.Persist.FindAllExams(ctx)
	if err != nil {
		return query.FindAllExamsWithQuestionsResponse{}, err
	}

	result := mapper.FromDomainExamsToAppFindAllExamsWithQuestionsResponse(entity)

	return result, nil
}
