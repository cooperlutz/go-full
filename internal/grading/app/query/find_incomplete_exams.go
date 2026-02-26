package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindIncompleteExamsReadModel interface {
	FindIncompleteExams(ctx context.Context) ([]Exam, error)
}

type FindIncompleteExamsHandler struct {
	readModel FindIncompleteExamsReadModel
}

func NewFindIncompleteExamsHandler(
	readModel FindIncompleteExamsReadModel,
) FindIncompleteExamsHandler {
	return FindIncompleteExamsHandler{readModel: readModel}
}

func (h FindIncompleteExamsHandler) Handle(ctx context.Context) ([]Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.query.find_incomplete_exams.handle")
	defer span.End()

	exams, err := h.readModel.FindIncompleteExams(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
