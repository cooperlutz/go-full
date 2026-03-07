package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindQuestion struct {
	ExamID        string
	QuestionIndex int32
}

type FindQuestionHandler struct {
	readModel FindQuestionReadModel
}

func NewFindQuestionHandler(
	readModel FindQuestionReadModel,
) FindQuestionHandler {
	return FindQuestionHandler{readModel: readModel}
}

type FindQuestionReadModel interface {
	FindQuestion(ctx context.Context, examID uuid.UUID, questionIndex int32) (Question, error)
}

func (h FindQuestionHandler) Handle(ctx context.Context, query FindQuestion) (Question, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.query.find_question.handle")
	defer span.End()

	examId, err := uuid.Parse(query.ExamID)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Question{}, err
	}

	return h.readModel.FindQuestion(ctx, examId, query.QuestionIndex)
}
