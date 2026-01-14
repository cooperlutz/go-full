package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/event"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

// Exam represents an exam entity in the exam library.
// Exam serves as the aggregate root for exam library related operations.
type Exam struct {
	*baseentitee.EntityMetadata
	name       string
	gradeLevel valueobject.GradeLevel
	questions  *[]ExamQuestion
}

// NewExam creates a new Exam entity.
func NewExam(name string, gradeLevel valueobject.GradeLevel, questions *[]ExamQuestion) *Exam {
	e := &Exam{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		name:           name,
		gradeLevel:     gradeLevel,
		questions:      questions,
	}
	e.RaiseDomainEvent(
		event.ExamAddedToLibrary{ExamID: e.GetIdString()},
	)

	return e
}

// AddQuestion adds a new question to the exam.
func (e *Exam) AddQuestion(question ExamQuestion) {
	if e.questions == nil {
		e.questions = &[]ExamQuestion{}
	}

	*e.questions = append(*e.questions, question)
	e.MarkUpdated()
}

// GetQuestions returns all questions associated with the exam.
func (e Exam) GetQuestions() []ExamQuestion {
	if e.questions == nil {
		return []ExamQuestion{}
	}

	return *e.questions
}

// GetQuestionByIndex retrieves a question by its index in the exam.
func (e Exam) GetQuestionByIndex(index int) (ExamQuestion, error) {
	if e.questions == nil || index < 0 || index >= len(*e.questions) {
		return ExamQuestion{}, exception.ErrInvalidIndex{}
	}

	return (*e.questions)[index], nil
}

// GetQuestionById retrieves a question by its unique identifier.
func (e Exam) GetQuestionById(id uuid.UUID) (*ExamQuestion, error) {
	for _, q := range *e.questions {
		if q.GetIdUUID() == id {
			return &q, nil
		}
	}

	return nil, exception.ErrQuestionNotFound{}
}

// GetName returns the name of the exam.
func (e *Exam) GetName() string {
	return e.name
}

// GetGradeLevel returns the grade level of the exam.
func (e *Exam) GetGradeLevel() valueobject.GradeLevel {
	return e.gradeLevel
}

// MapToExamEntity maps raw data to an Exam entity.
// This is typically used when reconstructing an Exam from persistent storage.
func MapToExamEntity(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	name string,
	gradeLevel valueobject.GradeLevel,
	questions *[]ExamQuestion,
) Exam {
	exam := Exam{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		name:       name,
		gradeLevel: gradeLevel,
		questions:  questions,
	}

	return exam
}
