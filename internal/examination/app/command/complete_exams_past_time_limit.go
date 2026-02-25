package command

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type CompleteExamsPastTimeLimit struct{}

type AllExamsReadModel interface {
	FindAllInProgress(ctx context.Context) ([]query.Exam, error)
}

type CompleteExamsPastTimeLimitHandler struct {
	examinationRepo   examination.Repository
	allExamsReadModel AllExamsReadModel
}

func NewCompleteExamsPastTimeLimitHandler(
	examinationRepo examination.Repository,
	allExamsReadModel AllExamsReadModel,
) CompleteExamsPastTimeLimitHandler {
	return CompleteExamsPastTimeLimitHandler{
		examinationRepo:   examinationRepo,
		allExamsReadModel: allExamsReadModel,
	}
}

func (h CompleteExamsPastTimeLimitHandler) Handle(ctx context.Context) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.command.complete_exams_past_time_limit.handle")
	defer span.End()

	exams, err := h.allExamsReadModel.FindAllInProgress(ctx)
	if err != nil {
		return err
	}

	for _, exam := range exams {
		err = h.examinationRepo.UpdateExam(ctx, uuid.MustParse(exam.ExamId), func(e *examination.Exam) (*examination.Exam, error) {
			err = e.CheckTimeLimit()
			if err != nil && !errors.Is(err, examination.ErrTimeLimitExceeded{}) {
				telemetree.RecordError(ctx, err)

				return nil, err
			}

			return e, nil
		})
	}

	return nil
}
