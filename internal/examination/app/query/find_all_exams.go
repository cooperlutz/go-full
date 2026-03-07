package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AllExamsReadModel interface {
	FindAll(ctx context.Context) ([]Exam, error)
}

type AllExamsHandler struct {
	readModel AllExamsReadModel
}

func NewFindAllExamsHandler(
	readModel AllExamsReadModel,
) AllExamsHandler {
	return AllExamsHandler{readModel: readModel}
}

func (h AllExamsHandler) Handle(ctx context.Context) ([]Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.query.find_all_exams.handle")
	defer span.End()

	exams, err := h.readModel.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
