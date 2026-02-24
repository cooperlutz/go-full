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

func (h StartExamHandler) Handle(ctx context.Context, cmd StartExam) (Exam, error) { //nolint:funlen // it's fine
	ctx, span := telemetree.AddSpan(ctx, "examination.app.command.start_exam.handle")
	defer span.End()

	// retrieve the exam questions from the Exam Library service using the provided exam library ID
	examLibraryExam, err := h.examLibraryAdapter.RetrieveExamQuestionsFromLibrary(ctx, cmd.ExamLibraryID)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	// map the retrieved exam questions to the domain entity format
	var questions []*examination.Question

	if examLibraryExam.Questions != nil {
		for _, q := range *examLibraryExam.Questions {
			questionType, err := examination.QuestionTypeFromString(q.QuestionType)
			if err != nil {
				telemetree.RecordError(ctx, err)

				return Exam{}, err
			}

			question := examination.NewQuestion(
				int32(q.Index),
				q.QuestionText,
				questionType,
				q.ResponseOptions,
			)
			questions = append(questions, question)
		}
	}

	studentIdUuid, err := uuid.Parse(cmd.StudentId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return Exam{}, err
	}

	exam := examination.NewExam(
		studentIdUuid,
		uuid.MustParse(cmd.ExamLibraryID),
		examLibraryExam.TimeLimit,
		questions,
	)

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
			ExamId:          exam.GetIdString(),
			Answered:        q.IsAnswered(),
			QuestionID:      q.GetIdString(),
			QuestionIndex:   q.GetIndex(),
			QuestionText:    q.GetQuestionText(),
			QuestionType:    q.GetQuestionType().String(),
			ResponseOptions: q.GetResponseOptions(),
			ProvidedAnswer:  q.GetProvidedAnswer(),
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
