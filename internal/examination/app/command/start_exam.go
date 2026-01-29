package command

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type StartExam struct {
	StudentId     string
	ExamLibraryID string
}

type StartExamHandler struct {
	examinationRepo    examination.Repository
	examLibraryAdapter outbound.ExamLibraryAdapter
}

func NewStartExamHandler(
	examinationRepo examination.Repository,
	examLibraryAdapter outbound.ExamLibraryAdapter,
) StartExamHandler {
	return StartExamHandler{examinationRepo: examinationRepo, examLibraryAdapter: examLibraryAdapter}
}

func (h StartExamHandler) Handle(ctx context.Context, cmd StartExam) (Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.app.command.startexam.handle")
	defer span.End()

	questions, err := h.examLibraryAdapter.RetrieveExamQuestionsFromLibrary(ctx, cmd.ExamLibraryID)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	examIdUuid, err := uuid.Parse(cmd.StudentId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	exam := examination.NewExam(examIdUuid, uuid.MustParse(cmd.ExamLibraryID), questions)

	err = exam.StartExam()
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	err = h.examinationRepo.AddExam(ctx, exam)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	var questionsForExam []Question
	for _, q := range questions {
		questionsForExam = append(questionsForExam, Question{
			QuestionID:      q.GetIdString(),
			QuestionIndex:   q.GetIndex(),
			QuestionText:    q.GetQuestionText(),
			QuestionType:    q.GetQuestionType().String(),
			ResponseOptions: *q.GetResponseOptions(),
		})
	}

	return Exam{
		ExamId:            exam.GetIdString(),
		StudentId:         exam.GetStudentIdString(),
		LibraryExamId:     exam.GetLibraryExamIdUUID().String(),
		AnsweredQuestions: exam.AnsweredQuestionsCount(),
		TotalQuestions:    exam.NumberOfQuestions(),
		Questions:         questionsForExam,
	}, nil
}
