package examination

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type Exam struct {
	*baseentitee.EntityMetadata
	studentId   uuid.UUID
	startedAt   *time.Time
	completedAt *time.Time
	completed   bool
	questions   []*Question
}

func NewExam(studentId uuid.UUID, questions []*Question) *Exam {
	return &Exam{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		studentId:      studentId,
		questions:      questions,
	}
}

func (e Exam) GetQuestions() []*Question {
	return e.questions
}

func (e Exam) numberOfQuestions() int32 {
	length := len(e.questions)

	return utilitee.SafeIntToInt32(&length)
}

func (e Exam) GetQuestionByIndex(index int32) *Question {
	if index < 1 || int(index) > len(e.questions) {
		return nil
	}

	return e.questions[index-1]
}

func (e Exam) GetFirstQuestion() *Question {
	return e.GetQuestionByIndex(1)
}

func (e Exam) GetCompletedAtTime() *time.Time {
	return e.completedAt
}

func (e Exam) IsCompleted() bool {
	return e.completed
}

func (e Exam) GetStudentIdUUID() uuid.UUID {
	return e.studentId
}

func (e Exam) GetStartedAtTime() *time.Time {
	return e.startedAt
}

func (e Exam) GetStudentIdString() string {
	return e.studentId.String()
}

type Question struct {
	*baseentitee.EntityMetadata
	index           int32
	answered        bool
	questionText    string
	questionType    QuestionType
	providedAnswer  *string
	responseOptions *[]string
}

// GetIndex returns the index of the question in the exam.
func (q Question) GetIndex() int32 {
	return q.index
}

// IsAnswered returns whether the question has been answered.
func (q Question) IsAnswered() bool {
	return q.answered
}

// GetQuestionText returns the text of the question.
func (q Question) GetQuestionText() string {
	return q.questionText
}

// GetQuestionType returns the type of the question.
func (q Question) GetQuestionType() QuestionType {
	return q.questionType
}

// GetProvidedAnswer returns the provided answer for the question.
func (q Question) GetProvidedAnswer() *string {
	return q.providedAnswer
}

// GetResponseOptions returns the response options for the question.
func (q Question) GetResponseOptions() *[]string {
	return q.responseOptions
}

// NewQuestion creates a new Question entity.
func NewQuestion(
	index int32,
	questionText string,
	questionType QuestionType,
	options *[]string,
) *Question {
	return &Question{
		EntityMetadata:  baseentitee.NewEntityMetadata(),
		index:           index,
		questionText:    questionText,
		questionType:    questionType,
		responseOptions: options,
	}
}

type ErrInvalidQuestionType struct{}

func (e ErrInvalidQuestionType) Error() string {
	return "invalid question type"
}

// QuestionType represents the type of a question in an exam.
// It is defined as an enumeration.
type QuestionType int

const (
	QuestionMultipleChoice QuestionType = iota
	QuestionEssay
	QuestionShortAnswer
)

var questionName = map[QuestionType]string{ //nolint:gochecknoglobals // global is ok here
	QuestionMultipleChoice: "multiple-choice",
	QuestionEssay:          "essay",
	QuestionShortAnswer:    "short-answer",
}

func (qt QuestionType) String() string {
	return questionName[qt]
}

func (qt QuestionType) Int() int {
	return int(qt)
}

func QuestionTypeFromString(s string) (QuestionType, error) {
	for qt, name := range questionName {
		if name == s {
			return qt, nil
		}
	}

	return QuestionMultipleChoice, ErrInvalidQuestionType{}
}
