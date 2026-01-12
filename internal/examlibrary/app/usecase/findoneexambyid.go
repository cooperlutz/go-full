package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func (uc *examLibraryUseCase) FindOneExamByID(ctx context.Context, qry query.FindOneExamByID) (query.FindOneExamByIDResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.usecase.findoneexambyid")
	defer span.End()

	entity, err := uc.Persist.FindExamByID(ctx, uuid.MustParse(qry.ExamID))
	if err != nil {
		return query.FindOneExamByIDResponse{}, err
	}

	result := mapper.FromDomainExamToAppFindOneExamByIDResponse(entity)

	return result, nil
}
