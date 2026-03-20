package examination

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
	"github.com/google/uuid"
)

type Exam struct {
	*baseentitee.EntityMetadata
	studentId     uuid.UUID
	libraryExamId uuid.UUID

	timeLimit       int64
	timeOfTimeLimit *time.Time
	startedAt       *time.Time
	completedAt     *time.Time

	state ExamState

	questions []*Question
}

func NewExam(studentId, libraryExamId uuid.UUID, timeLimit int64, questions []*Question) *Exam {
	return &Exam{
		EntityMetadata:  baseentitee.NewEntityMetadata(),
		studentId:       studentId,
		libraryExamId:   libraryExamId,
		timeLimit:       timeLimit,
		timeOfTimeLimit: nil, // timeOfTimeLimit is set when the exam is started
		startedAt:       nil, // startedAt is set when the exam is started
		completedAt:     nil, // completedAt is set when the exam is completed
		state:           StateNotStarted,
		questions:       questions,
	}
}

// Getters for Exam fields.
func (e Exam) GetQuestions() []*Question       { return e.questions }
func (e Exam) GetLibraryExamIdUUID() uuid.UUID { return e.libraryExamId }
func (e Exam) GetFirstQuestion() *Question     { return e.GetQuestionByIndex(1) }
func (e Exam) GetTimeLimitSeconds() int64      { return e.timeLimit }
func (e Exam) GetTimeOfTimeLimit() *time.Time  { return e.timeOfTimeLimit }
func (e Exam) GetCompletedAtTime() *time.Time  { return e.completedAt }
func (e Exam) GetStudentIdUUID() uuid.UUID     { return e.studentId }
func (e Exam) GetStartedAtTime() *time.Time    { return e.startedAt }
func (e Exam) GetStudentIdString() string      { return e.studentId.String() }
func (e Exam) IsCompleted() bool               { return e.state == StateCompleted }
func (e Exam) GetState() ExamState             { return e.state }

func (e Exam) GetQuestionByIndex(index int32) *Question {
	if index < 1 || int(index) > len(e.questions) {
		return nil
	}

	return e.questions[index-1]
}

type ExamState int

const (
	StateNotStarted ExamState = iota
	StateInProgress
	StateCompleted
)

var stateName = map[ExamState]string{ //nolint:gochecknoglobals // global is ok here for enum
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

func (q Question) GetExamId() uuid.UUID          { return q.examId }
func (q Question) GetIndex() int32               { return q.index }
func (q Question) IsAnswered() bool              { return q.answered }
func (q Question) GetQuestionText() string       { return q.questionText }
func (q Question) GetQuestionType() QuestionType { return q.questionType }
func (q Question) GetProvidedAnswer() *string    { return q.providedAnswer }
func (q Question) GetResponseOptions() *[]string { return q.responseOptions }

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
