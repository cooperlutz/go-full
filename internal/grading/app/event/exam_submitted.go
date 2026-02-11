//nolint:funlen // I just don't want to make this function smaller right now
package event

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"

	exam_library_query "github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// {"ExamId":"af6830f8-b146-42f0-bbcd-473d8068b234","LibraryExamId":"11111111-1111-1111-1111-111111111111","StudentId":"e95cb30c-2012-41f5-8de8-02ef608f56e5","Completed":true,"AnsweredQuestions":3,"TotalQuestions":3,"Questions":[{"ExamId":"af6830f8-b146-42f0-bbcd-473d8068b234","Answered":true,"QuestionID":"5ece76af-2359-4ade-9130-3e404473c536","QuestionIndex":1,"QuestionText":"What is the capital of France?","QuestionType":"multiple-choice","ResponseOptions":["Berlin","Madrid","Paris","Rome"],"ProvidedAnswer":"Berlin"},{"ExamId":"af6830f8-b146-42f0-bbcd-473d8068b234","Answered":true,"QuestionID":"2fffee65-94d1-481c-80b0-5a6390dcd2e7","QuestionIndex":2,"QuestionText":"What is Go?","QuestionType":"short-answer","ResponseOptions":null,"ProvidedAnswer":"idk"},{"ExamId":"af6830f8-b146-42f0-bbcd-473d8068b234","Answered":true,"QuestionID":"536969b1-11fe-4bc1-a13f-824b34af0398","QuestionIndex":3,"QuestionText":"Explain the concept of concurrency.","QuestionType":"essay","ResponseOptions":null,"ProvidedAnswer":"huh"}]}.
type ExamSubmitted struct {
	ExamId            string                  `json:"ExamId"`
	LibraryExamId     string                  `json:"LibraryExamId"`
	StudentId         string                  `json:"StudentId"`
	Completed         bool                    `json:"Completed"`
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

// ExamSubmittedHandler is a handler that does nothing and acknowledges the message.
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

// Handle returns a message handler function that acknowledges the message without processing.
func (h ExamSubmittedHandler) Handle() message.NoPublishHandlerFunc {
	ctx := context.Background()

	return func(msg *message.Message) error {
		var event ExamSubmitted

		err := json.Unmarshal(msg.Payload, &event)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		examFromLibrary, err := h.examLibrary.FindOneExamByID(
			ctx, exam_library_query.FindOneExamByID{
				ExamID: event.LibraryExamId,
			},
		)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

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

		exam := grading.NewExam(
			uuid.MustParse(event.StudentId),
			uuid.MustParse(event.LibraryExamId),
			uuid.MustParse(event.ExamId),
			questions,
		)

		err = exam.GradeMultipleChoiceQuestions()
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		err = h.gradingRepo.AddExam(ctx, exam)
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}

		msg.Ack()

		return nil
	}
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
