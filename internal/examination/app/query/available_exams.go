package query

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
)

type Exam struct {
	Name       string
	GradeLevel int32
}

type AvailableExamsReadModel interface {
	FindAll(ctx context.Context) ([]examination.Exam, error)
}

type AvailableExamsHandler struct {
	readModel AvailableExamsReadModel
}

func NewAvailableExamsHandler(
	readModel AvailableExamsReadModel,
) AvailableExamsHandler {
	return AvailableExamsHandler{readModel: readModel}
}

func (h AvailableExamsHandler) Handle(ctx context.Context) (d []Exam, err error) {
	exams, err := h.readModel.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, e := range exams {
		d = append(d, Exam{
			Name:       e.Name,
			GradeLevel: e.GradeLevel,
		})
	}

	return d, nil
}
