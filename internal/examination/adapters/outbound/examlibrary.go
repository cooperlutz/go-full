package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/utilitee"
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
) ([]*examination.Question, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.examlibrary.retrieveexamquestionsfromlibrary")
	defer span.End()

	response, err := a.uc.FindOneExamByID(ctx, query.FindOneExamByID{ExamID: examID})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	var questions []*examination.Question

	if response.Questions != nil {
		for _, q := range *response.Questions {
			questionType, err := examination.QuestionTypeFromString(q.QuestionType)
			if err != nil {
				return nil, err
			}

			question := examination.NewQuestion(
				utilitee.SafeIntToInt32(&q.Index),
				q.QuestionText,
				questionType,
				q.ResponseOptions,
			)
			questions = append(questions, question)
		}
	}

	return questions, nil
}
