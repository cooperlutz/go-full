package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindExamQuestionReadModel interface {
	FindExamQuestion(ctx context.Context, examId string, questionIndex int32) (Question, error)
}

type FindExamQuestionHandler struct {
	readModel FindExamQuestionReadModel
}

func NewFindExamQuestionHandler(
	readModel FindExamQuestionReadModel,
) FindExamQuestionHandler {
	return FindExamQuestionHandler{readModel: readModel}
}

func (h FindExamQuestionHandler) Handle(ctx context.Context, examId string, questionIndex int32) (Question, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.query.findexamquestion.handle")
	defer span.End()

	question, err := h.readModel.FindExamQuestion(ctx, examId, questionIndex)
	if err != nil {
		return Question{}, err
	}

	return question, nil
}
