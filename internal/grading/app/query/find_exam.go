package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindExamReadModel interface {
	FindExam(ctx context.Context, examId string) (Exam, error)
}

type FindExamHandler struct {
	readModel FindExamReadModel
}

func NewFindExamHandler(
	readModel FindExamReadModel,
) FindExamHandler {
	return FindExamHandler{readModel: readModel}
}

func (h FindExamHandler) Handle(ctx context.Context, examId string) (Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.query.find_exam.handle")
	defer span.End()

	exam, err := h.readModel.FindExam(ctx, examId)
	if err != nil {
		return Exam{}, err
	}

	return exam, nil
}
