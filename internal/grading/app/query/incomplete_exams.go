package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type IncompleteExamsReadModel interface {
	FindIncompleteExams(ctx context.Context) ([]Exam, error)
}

type IncompleteExamsHandler struct {
	readModel IncompleteExamsReadModel
}

func NewIncompleteExamsHandler(
	readModel IncompleteExamsReadModel,
) IncompleteExamsHandler {
	return IncompleteExamsHandler{readModel: readModel}
}

func (h IncompleteExamsHandler) Handle(ctx context.Context) ([]Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.query.incompleteexams.handle")
	defer span.End()

	exams, err := h.readModel.FindIncompleteExams(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
