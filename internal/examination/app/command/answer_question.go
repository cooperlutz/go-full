package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type AnswerQuestion struct {
	ExamID        string
	QuestionIndex int32
	Answer        string
}

type AnswerQuestionHandler struct {
	examinationRepo    examination.Repository
	examLibraryAdapter outbound.ExamLibraryAdapter
}

func NewAnswerQuestionHandler(
	examinationRepo examination.Repository,
	examLibraryAdapter outbound.ExamLibraryAdapter,
) AnswerQuestionHandler {
	return AnswerQuestionHandler{examinationRepo: examinationRepo, examLibraryAdapter: examLibraryAdapter}
}

func (h AnswerQuestionHandler) Handle(ctx context.Context, cmd AnswerQuestion) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.command.answerquestion.handle")
	defer span.End()

	examIdUuid, err := uuid.Parse(cmd.ExamID)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	exam, err := h.examinationRepo.GetExam(ctx, examIdUuid)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = h.examinationRepo.UpdateExam(ctx, exam, func(e *examination.Exam) (*examination.Exam, error) {
		err = e.AnswerQuestion(cmd.QuestionIndex, cmd.Answer)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return nil, err
		}

		return e, nil
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
