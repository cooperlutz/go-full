package grading

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Exam struct {
	*baseentitee.EntityMetadata
	studentId         uuid.UUID
	libraryExamId     uuid.UUID // corresponds to the examId in the exam library domain
	examinationExamId uuid.UUID // corresponds to the examId in the examination domain
	questions         []*Question

	gradingCompleted    bool
	totalPointsReceived *int32
	totalPossiblePoints int32
}

func NewExam(studentId, libraryExamId, examinationExamId uuid.UUID, questions []*Question) *Exam {
	e := &Exam{
		EntityMetadata:    baseentitee.NewEntityMetadata(),
		studentId:         studentId,
		libraryExamId:     libraryExamId,
		examinationExamId: examinationExamId,
		questions:         questions,
	}
	e.CalculateTotalPossiblePoints()

	return e
}

type ErrGradingNotCompleted struct{}

func (e ErrGradingNotCompleted) Error() string {
	return "grading not completed"
}

func (e Exam) GetGrade() (float64, error) {
	if !e.gradingCompleted {
		return 0, ErrGradingNotCompleted{}
	}

	// If there are no possible points, we can consider the grade to be 0% to avoid division by zero.
	if e.totalPossiblePoints == 0 {
		return 0, nil
	}

	grade := float64(*e.totalPointsReceived) / float64(e.totalPossiblePoints) * 100 //nolint:mnd // 100 is a constant for percentage calculation

	return grade, nil
}

func (e Exam) GetQuestions() []*Question {
	return e.questions
}

func (e Exam) GetExaminationExamId() uuid.UUID {
	return e.examinationExamId
}

func (e *Exam) CalculateTotalPossiblePoints() {
	var total int32
	for _, q := range e.GetQuestions() {
		total += q.pointsPossible
	}

	e.totalPossiblePoints = total
	e.MarkUpdated()
}

func (e *Exam) CalculateTotalPointsReceived() {
	var total int32

	for _, q := range e.GetQuestions() {
		if q.pointsReceived != nil {
			total += *q.pointsReceived
		}
	}

	e.totalPointsReceived = &total
	e.MarkUpdated()
}

func (e Exam) GetTotalPointsReceived() *int32 {
	return e.totalPointsReceived
}

func (e Exam) GetTotalPointsPossible() int32 {
	return e.totalPossiblePoints
}

func (e Exam) GetMultiplChoiceQuestions() []*Question {
	var questions []*Question

	for _, q := range e.GetQuestions() {
		if q.GetQuestionType() == QuestionMultipleChoice {
			questions = append(questions, q)
		}
	}

	return questions
}

func (e Exam) GetUngradedQuestions() []*Question {
	var questions []*Question

	for _, q := range e.GetQuestions() {
		if !q.graded {
			questions = append(questions, q)
		}
	}

	return questions
}

type ErrMultipleChoiceGradingFailed struct{}

func (e ErrMultipleChoiceGradingFailed) Error() string {
	return "grading multiple-choice question failed"
}

func (e *Exam) GradeMultipleChoiceQuestions() error {
	for _, q := range e.GetMultiplChoiceQuestions() {
		err := q.gradeQuestion(GradeQuestionOption{})
		if err != nil {
			return ErrMultipleChoiceGradingFailed{}
		}
	}

	e.MarkUpdated()

	return nil
}

func (e Exam) GetQuestionByIndex(index int32) *Question {
	if index < 1 || int(index) > len(e.questions) {
		return nil
	}

	return e.questions[index-1]
}

func (e Exam) GetExamLibraryExamId() uuid.UUID {
	return e.libraryExamId
}

func (e Exam) GetFirstQuestion() *Question {
	return e.GetQuestionByIndex(1)
}

func (e Exam) IsCompleted() bool {
	return e.gradingCompleted
}

func (e Exam) GetStudentId() uuid.UUID {
	return e.studentId
}

func (e Exam) GetStudentIdString() string {
	return e.studentId.String()
}

type Question struct {
	*baseentitee.EntityMetadata
	examId       uuid.UUID
	questionType QuestionType
	index        int32

	graded   bool
	feedback *string

	providedAnswer    string
	correctAnswer     *string // nil for non-multiple-choice questions
	correctlyAnswered *bool

	pointsPossible int32
	pointsReceived *int32
}

// GetIndex returns the index of the question in the exam.
func (q Question) GetIndex() int32 {
	return q.index
}

// GetQuestionType returns the type of the question.
func (q Question) GetQuestionType() QuestionType {
	return q.questionType
}

func (q Question) GetPointsReceived() *int32 {
	return q.pointsReceived
}

func (q Question) GetPointsPossible() int32 {
	return q.pointsPossible
}

func (q Question) IsGraded() bool {
	return q.graded
}

func (q Question) GetFeedback() *string {
	return q.feedback
}

func (q Question) GetProvidedAnswer() string {
	return q.providedAnswer
}

func (q Question) GetCorrectAnswer() *string {
	return q.correctAnswer
}

func (q Question) IsCorrectlyAnswered() *bool {
	return q.correctlyAnswered
}

func (q *Question) markAsGraded() {
	q.graded = true
	q.MarkUpdated()
}

// NewQuestion creates a new Question entity.
func NewQuestion(
	// examId uuid.UUID,
	questionType QuestionType,
	index int32,
	providedAnswer string,
	correctAnswer *string,
	pointsPossible int32,
) *Question {
	return &Question{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		index:          index,
		// examId:         examId,
		questionType:   questionType,
		graded:         false,
		feedback:       nil,
		providedAnswer: providedAnswer,
		correctAnswer:  correctAnswer,
		pointsPossible: pointsPossible,
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
