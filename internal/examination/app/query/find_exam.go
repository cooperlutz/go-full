package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindExam struct {
	ExamID string
}

type FindExamHandler struct {
	readModel FindExamReadModel
}

func NewFindExamHandler(
	readModel FindExamReadModel,
) FindExamHandler {
	return FindExamHandler{readModel: readModel}
}

type FindExamReadModel interface {
	FindExam(ctx context.Context, examID uuid.UUID) (Exam, error)
}

func (h FindExamHandler) Handle(ctx context.Context, qry FindExam) (Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.query.find_exam.handle")
	defer span.End()

	examId := uuid.MustParse(qry.ExamID)

	return h.readModel.FindExam(ctx, examId)
}
