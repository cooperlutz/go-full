package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// FindAllExamsWithoutQuestions finds all exams without their questions.
func (uc *examLibraryUseCase) FindAllExamsWithoutQuestions(ctx context.Context, qry query.FindAllExamsWithoutQuestions) (query.FindAllExamsWithoutQuestionsResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.usecase.find_all_exams_without_questions")
	defer span.End()

	entity, err := uc.Persist.FindAllExams(ctx)
	if err != nil {
		return query.FindAllExamsWithoutQuestionsResponse{}, err
	}

	result := mapper.FromDomainExamsToAppFindAllExamsWithoutQuestionsResponse(entity)

	return result, nil
}
