package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func (uc *examLibraryUseCase) FindAllExamsWithQuestions(ctx context.Context, qry query.FindAllExamsWithQuestions) (query.FindAllExamsWithQuestionsResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.usecase.findallexamswithquestions")
	defer span.End()

	entity, err := uc.Persist.FindAllExams(ctx)
	if err != nil {
		return query.FindAllExamsWithQuestionsResponse{}, err
	}

	result := mapper.FromDomainExamsToAppFindAllExamsWithQuestionsResponse(entity)

	return result, nil
}
