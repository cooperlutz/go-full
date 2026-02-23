//nolint:funlen // I just don't want to make this function smaller right now
package event

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"

	exam_library_query "github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type ExamSubmitted struct {
	ExamId            string                  `json:"ExamId"`
	LibraryExamId     string                  `json:"LibraryExamId"`
	StudentId         string                  `json:"StudentId"`
	State             string                  `json:"State"`
	AnsweredQuestions int32                   `json:"AnsweredQuestions"`
	TotalQuestions    int32                   `json:"TotalQuestions"`
	Questions         []ExamSubmittedQuestion `json:"Questions"`
}

type ExamSubmittedQuestion struct {
	ExamId          string    `json:"ExamId"`
	Answered        bool      `json:"Answered"`
	QuestionID      string    `json:"QuestionID"`
	QuestionIndex   int32     `json:"QuestionIndex"`
	QuestionText    string    `json:"QuestionText"`
	QuestionType    string    `json:"QuestionType"`
	ResponseOptions *[]string `json:"ResponseOptions"`
	ProvidedAnswer  *string   `json:"ProvidedAnswer"`
}

// ExamSubmittedHandler is a handler that processes exam submitted events.
type ExamSubmittedHandler struct {
	gradingRepo grading.Repository
	examLibrary usecase.IExamLibraryUseCase
}

// NewExamSubmittedHandler creates a new ExamSubmittedHandler instance.
func NewExamSubmittedHandler(gradingRepo grading.Repository, examLibraryUseCase usecase.IExamLibraryUseCase) ExamSubmittedHandler {
	return ExamSubmittedHandler{
		gradingRepo: gradingRepo,
		examLibrary: examLibraryUseCase,
	}
}

// Handle returns a message handler function that processes exam submitted events.
//
//nolint:cyclop,gocyclo,gocognit // I don't want to split this function up right now
func (h ExamSubmittedHandler) Handle(
	gradingStartedHandler GradingStartedHandler,
	gradingCompletedHandler GradingCompletedHandler,
) message.NoPublishHandlerFunc {
	return eeventdriven.TraceConsumerHandler(func(msg *message.Message) error {
		ctx, span := telemetree.AddSpan(msg.Context(), "grading.app.event.exam_submitted.handle")
		defer span.End()

		// unmarshal the event
		var event ExamSubmitted

		err := json.Unmarshal(msg.Payload, &event)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		// get the exam from the exam library to get the correct answers and possible points for each question
		examFromLibrary, err := h.examLibrary.FindOneExamByID(
			ctx, exam_library_query.FindOneExamByID{
				ExamID: event.LibraryExamId,
			},
		)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		// create grading questions from the submitted exam questions and the exam from the library
		var questions []*grading.Question

		for _, q := range event.Questions {
			libQuestion := parseExamLibraryQuestion(q.QuestionIndex, examFromLibrary)

			qType, err := grading.QuestionTypeFromString(q.QuestionType)
			if err != nil {
				telemetree.RecordError(ctx, err)

				return err
			}

			questions = append(questions, grading.NewQuestion(
				qType,
				q.QuestionIndex,
				*q.ProvidedAnswer,
				libQuestion.CorrectAnswer,
				utilitee.SafeIntToInt32(&libQuestion.PossiblePoints),
			))
		}

		// create an exam aggregate from the submitted exam event and the grading questions
		exam := grading.NewExam(
			uuid.MustParse(event.StudentId),
			uuid.MustParse(event.LibraryExamId),
			uuid.MustParse(event.ExamId),
			questions,
		)

		// store the exam aggregate in the grading repository
		err = h.gradingRepo.AddExam(ctx, exam)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		// update the exam aggregate in the grading repository to grade multiple choice questions and emit grading started and grading completed events if grading is complete
		err = h.gradingRepo.UpdateExam(ctx, exam.GetIdUUID(), func(e *grading.Exam) (*grading.Exam, error) {
			// grade multiple choice questions automatically when the exam is submitted, and check if grading is complete
			gradingIsComplete, err := e.GradeMultipleChoiceQuestions()
			if err != nil {
				telemetree.RecordError(ctx, err)

				return nil, err
			}

			// emit grading started event
			err = gradingStartedHandler.Handle(ctx, GradingStarted{
				ExamId: event.ExamId,
			})
			if err != nil {
				telemetree.RecordError(ctx, err)

				return nil, err
			}

			// if grading is complete after grading multiple choice questions, emit grading completed event
			if gradingIsComplete {
				err = gradingCompletedHandler.Handle(ctx, GradingCompleted{
					ExamId: event.ExamId,
				})
				if err != nil {
					telemetree.RecordError(ctx, err)

					return nil, err
				}
			}

			return e, nil
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		return nil
	})
}

type QuestionFromExamLib struct {
	CorrectAnswer  *string
	PossiblePoints int
}

func parseExamLibraryQuestion(questionIndex int32, examFromLibrary exam_library_query.FindOneExamByIDResponse) QuestionFromExamLib {
	if examFromLibrary.Questions != nil {
		for _, q := range *examFromLibrary.Questions {
			if utilitee.SafeIntToInt32(&q.Index) == questionIndex {
				return QuestionFromExamLib{
					CorrectAnswer:  q.CorrectAnswer,
					PossiblePoints: q.PossiblePoints,
				}
			}
		}
	}

	return QuestionFromExamLib{}
}
