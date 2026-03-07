package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/app/event"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type SubmitExam struct {
	ExamID string
}

type SubmitExamHandler struct {
	examinationRepo           examination.Repository
	examLibraryAdapter        outbound.ExamLibraryAdapter
	examSubmittedEventHandler event.ExamSubmittedHandler
}

func NewSubmitExamHandler(
	examinationRepo examination.Repository,
	examLibraryAdapter outbound.ExamLibraryAdapter,
	examSubmittedEventHandler event.ExamSubmittedHandler,
) SubmitExamHandler {
	return SubmitExamHandler{
		examinationRepo:           examinationRepo,
		examLibraryAdapter:        examLibraryAdapter,
		examSubmittedEventHandler: examSubmittedEventHandler,
	}
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

		err = h.examSubmittedEventHandler.Handle(ctx, event.ExamSubmitted{
			ExamId:            e.GetIdString(),
			LibraryExamId:     e.GetLibraryExamIdUUID().String(),
			StudentId:         e.GetStudentIdString(),
			ExamState:         e.GetState().String(),
			AnsweredQuestions: e.AnsweredQuestionsCount(),
			TotalQuestions:    e.NumberOfQuestions(),
			TimeLimitSeconds:  e.GetTimeLimitSeconds(),
			TimeOfTimeLimit:   *e.GetTimeOfTimeLimit(),
			StartedAt:         *e.GetStartedAtTime(),
			CompletedAt:       *e.GetCompletedAtTime(),
			Questions: func() []event.ExamSubmittedQuestion {
				var questions []event.ExamSubmittedQuestion
				for _, q := range e.GetQuestions() {
					questions = append(questions, event.ExamSubmittedQuestion{
						ExamId:          q.GetExamId().String(),
						Answered:        q.IsAnswered(),
						QuestionID:      q.GetIdString(),
						QuestionIndex:   q.GetIndex(),
						QuestionText:    q.GetQuestionText(),
						QuestionType:    q.GetQuestionType().String(),
						ResponseOptions: q.GetResponseOptions(),
						ProvidedAnswer:  q.GetProvidedAnswer(),
					})
				}

				return questions
			}(),
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return nil, err
		}

		return e, nil
	})
}
