package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// ExamLibraryAdapter provides an adapter to interact with the Exam Library service.
// It wraps methods exposed via the IExamLibraryUseCase interface and maps the inputs/outputs
// to the domain entities used within the Examination module.
type ExamLibraryAdapter struct {
	uc usecase.IExamLibraryUseCase
}

// NewExamLibraryAdapter creates a new instance of ExamLibraryAdapter.
func NewExamLibraryAdapter(uc usecase.IExamLibraryUseCase) ExamLibraryAdapter {
	return ExamLibraryAdapter{uc: uc}
}

// RetrieveExamQuestionsFromLibrary fetches exam questions from the Exam Library service
// based on the provided examID. It maps the retrieved questions to the domain entity format.
func (a ExamLibraryAdapter) RetrieveExamQuestionsFromLibrary(
	ctx context.Context,
	examID string,
) (ExamLibraryExam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.examlibrary.retrieveexamquestionsfromlibrary")
	defer span.End()

	response, err := a.uc.FindOneExamByID(ctx, query.FindOneExamByID{ExamID: examID})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return ExamLibraryExam{}, err
	}

	var questions []ExamLibraryExamQuestion

	if response.Questions != nil {
		for _, q := range *response.Questions {
			question := ExamLibraryExamQuestion{
				Index:           q.Index,
				QuestionText:    q.QuestionText,
				QuestionType:    q.QuestionType,
				PossiblePoints:  q.PossiblePoints,
				CorrectAnswer:   q.CorrectAnswer,
				ResponseOptions: q.ResponseOptions,
			}
			questions = append(questions, question)
		}
	}

	exam := ExamLibraryExam{
		ExamID:     response.ExamID,
		Name:       response.Name,
		GradeLevel: response.GradeLevel,
		TimeLimit:  response.TimeLimit,
		Questions:  &questions,
	}

	return exam, nil
}

// ExamLibraryExam represents the response for the FindOneExamByID query.
type ExamLibraryExam struct {
	ExamID     string
	Name       string
	GradeLevel int
	TimeLimit  int64
	Questions  *[]ExamLibraryExamQuestion
}

// ExamLibraryExamQuestion represents a question in an exam.
type ExamLibraryExamQuestion struct {
	Index           int
	QuestionText    string
	QuestionType    string
	PossiblePoints  int
	CorrectAnswer   *string
	ResponseOptions *[]string
}
