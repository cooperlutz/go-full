package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type SubmitExam struct {
	ExamID string
}

type SubmitExamHandler struct {
	examinationRepo    examination.Repository
	examLibraryAdapter outbound.ExamLibraryAdapter
}

func NewSubmitExamHandler(
	examinationRepo examination.Repository,
	examLibraryAdapter outbound.ExamLibraryAdapter,
) SubmitExamHandler {
	return SubmitExamHandler{examinationRepo: examinationRepo, examLibraryAdapter: examLibraryAdapter}
}

func (h SubmitExamHandler) Handle(ctx context.Context, cmd SubmitExam) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.command.submit_exam.handle")
	defer span.End()

	examIdUuid, err := uuid.Parse(cmd.ExamID)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return h.examinationRepo.UpdateExam(ctx, examIdUuid, func(e *examination.Exam) (*examination.Exam, error) {
		err = e.Submit()
		if err != nil {
			telemetree.RecordError(ctx, err)

			return nil, err
		}

		return e, nil
	})
}
