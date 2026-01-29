package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AvailableExamsReadModel interface {
	FindAll(ctx context.Context) ([]Exam, error)
}

type AvailableExamsHandler struct {
	readModel AvailableExamsReadModel
}

func NewAvailableExamsHandler(
	readModel AvailableExamsReadModel,
) AvailableExamsHandler {
	return AvailableExamsHandler{readModel: readModel}
}

func (h AvailableExamsHandler) Handle(ctx context.Context) ([]Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.query.availableexams.handle")
	defer span.End()

	exams, err := h.readModel.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
