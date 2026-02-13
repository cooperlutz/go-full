package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type GradeQuestion struct {
	ExamId        string
	QuestionIndex int32
	Points        int32
	Feedback      string
}

type GradeQuestionHandler struct {
	gradingRepo grading.Repository
}

func NewGradeQuestionHandler(
	repo grading.Repository,
) GradeQuestionHandler {
	return GradeQuestionHandler{
		gradingRepo: repo,
	}
}

func (h GradeQuestionHandler) Handle(ctx context.Context, cmd GradeQuestion) error {
	ctx, span := telemetree.AddSpan(ctx, "grading.app.command.gradequestion.handle")
	defer span.End()

	return h.gradingRepo.UpdateExam(ctx, uuid.MustParse(cmd.ExamId), func(e *grading.Exam) (*grading.Exam, error) {
		err := e.GradeQuestion(cmd.QuestionIndex, grading.GradeQuestionOption{
			Feedback: cmd.Feedback,
			Points:   cmd.Points,
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return nil, err
		}

		return e, nil
	})
}
