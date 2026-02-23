package examination

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Exam struct {
	*baseentitee.EntityMetadata
	studentId     uuid.UUID
	libraryExamId uuid.UUID

	startedAt   *time.Time
	completedAt *time.Time

	state ExamState

	questions []*Question
}

func NewExam(studentId, libraryExamId uuid.UUID, questions []*Question) *Exam {
	return &Exam{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		studentId:      studentId,
		libraryExamId:  libraryExamId,
		questions:      questions,
		state:          StateNotStarted,
	}
}

func (e Exam) GetQuestions() []*Question {
	return e.questions
}

func (e Exam) GetQuestionByIndex(index int32) *Question {
	if index < 1 || int(index) > len(e.questions) {
		return nil
	}

	return e.questions[index-1]
}

func (e Exam) GetLibraryExamIdUUID() uuid.UUID {
	return e.libraryExamId
}

func (e Exam) GetFirstQuestion() *Question {
	return e.GetQuestionByIndex(1)
}

func (e Exam) GetCompletedAtTime() *time.Time {
	return e.completedAt
}

func (e Exam) IsCompleted() bool {
	return e.state == StateCompleted
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

type ExamState int

const (
	StateNotStarted ExamState = iota
	StateInProgress
	StateCompleted
)

var stateName = map[ExamState]string{
	StateNotStarted: "not-started",
	StateInProgress: "in-progress",
	StateCompleted:  "completed",
}

func (es ExamState) String() string {
	return stateName[es]
}

type ErrInvalidExamState struct{}

func (e ErrInvalidExamState) Error() string {
	return "invalid exam state"
}

func ExamStateFromString(s string) (ExamState, error) {
	for es, name := range stateName {
		if name == s {
			return es, nil
		}
	}

	return StateNotStarted, ErrInvalidExamState{}
}

func (e Exam) GetState() ExamState {
	return e.state
}

type Question struct {
	*baseentitee.EntityMetadata
	examId       uuid.UUID
	questionType QuestionType
	index        int32
	answered     bool

	questionText    string
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
