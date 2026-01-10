package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/event"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Exam struct {
	*baseentitee.EntityMetadata
	name       string
	gradeLevel valueobject.GradeLevel
	questions  *[]ExamQuestion
}

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

func (e *Exam) AddQuestion(question ExamQuestion) {
	*e.questions = append(*e.questions, question)
	e.MarkUpdated()
}

func (e Exam) GetQuestions() *[]ExamQuestion {
	return e.questions
}

func (e Exam) GetQuestionByIndex(index int) (ExamQuestion, error) {
	if e.questions == nil || index < 0 || index >= len(*e.questions) {
		return ExamQuestion{}, exception.ErrInvalidIndex{}
	}

	return (*e.questions)[index], nil
}

func (e Exam) GetQuestionById(id uuid.UUID) (*ExamQuestion, error) {
	for _, q := range *e.questions {
		if q.GetIdUUID() == id {
			return &q, nil
		}
	}

	return nil, exception.ErrQuestionNotFound{}
}

func (e *Exam) GetName() string {
	return e.name
}

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
