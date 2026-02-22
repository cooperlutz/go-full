package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

var (
	meta      = baseentitee.NewEntityMetadata()
	meta2     = baseentitee.NewEntityMetadata()
	meta3     = baseentitee.NewEntityMetadata()
	dog       = "dog"
	examTime  = 30 * time.Minute
	questions = []ExamQuestion{
		{
			EntityMetadata:  meta,
			index:           1,
			questionText:    "What animal is known to bark?",
			questionType:    valueobject.QuestionMultipleChoice,
			possiblePoints:  5,
			correctAnswer:   &dog,
			responseOptions: &[]string{"cat", "dog", "fish", "bird"},
		},
		{
			EntityMetadata:  meta2,
			index:           2,
			questionText:    "Describe the habitat of a dolphin.",
			questionType:    valueobject.QuestionShortAnswer,
			possiblePoints:  5,
			correctAnswer:   nil,
			responseOptions: nil,
		},
		{
			EntityMetadata:  meta3,
			index:           3,
			questionText:    "Explain the theory of relativity.",
			questionType:    valueobject.QuestionEssay,
			possiblePoints:  10,
			correctAnswer:   nil,
			responseOptions: nil,
		},
	}
)

func TestNewExam(t *testing.T) {
	// Act
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)
	// Assertions
	assert.Equal(t, "Animal Exam", exam.GetName())
	questionByIndex, err := exam.GetQuestionByIndex(0)
	assert.NoError(t, err)
	assert.Equal(t,
		ExamQuestion{
			EntityMetadata:  meta,
			index:           1,
			questionText:    "What animal is known to bark?",
			questionType:    valueobject.QuestionMultipleChoice,
			possiblePoints:  5,
			correctAnswer:   &dog,
			responseOptions: &[]string{"cat", "dog", "fish", "bird"},
		},
		questionByIndex,
	)
	result, err := exam.GetQuestionById(meta.GetIdUUID())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t,
		"dog",
		*result.correctAnswer,
	)
}

func TestExam_GetQuestionByIndex(t *testing.T) {
	// Act
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)
	// Assertions
	result, err := exam.GetQuestionByIndex(1)
	assert.NoError(t, err)
	assert.Equal(t,
		ExamQuestion{
			EntityMetadata:  meta2,
			index:           2,
			questionText:    "Describe the habitat of a dolphin.",
			questionType:    valueobject.QuestionShortAnswer,
			possiblePoints:  5,
			correctAnswer:   nil,
			responseOptions: nil,
		},
		result,
	)

	_, err = exam.GetQuestionByIndex(-1)
	assert.Error(t, err)

	_, err = exam.GetQuestionByIndex(10)
	assert.Error(t, err)
}

func TestExam_GetQuestionById(t *testing.T) {
	// Act
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)

	result, err := exam.GetQuestionById(meta.GetIdUUID())
	assert.NoError(t, err)
	// Assertions
	assert.Equal(t,
		ExamQuestion{
			EntityMetadata:  meta,
			index:           1,
			questionText:    "What animal is known to bark?",
			questionType:    valueobject.QuestionMultipleChoice,
			possiblePoints:  5,
			correctAnswer:   &dog,
			responseOptions: &[]string{"cat", "dog", "fish", "bird"},
		},
		*result,
	)
	assert.NotNil(t, result)

	_, err = exam.GetQuestionById(uuid.New())
	assert.Error(t, err)
}

func TestExam_GetQuestions(t *testing.T) {
	// Act
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)
	// Assertions
	assert.Equal(t, questions, exam.GetQuestions())
}

func TestExam_AddQuestion(t *testing.T) {
	// Arrange
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		nil,
	)
	newQuestion := ExamQuestion{
		EntityMetadata:  meta,
		index:           1,
		questionText:    "What animal is known to bark?",
		questionType:    valueobject.QuestionMultipleChoice,
		possiblePoints:  5,
		correctAnswer:   &dog,
		responseOptions: &[]string{"cat", "dog", "fish", "bird"},
	}
	// Act
	exam.AddQuestion(newQuestion)
	// Assertions
	assert.Equal(t, 1, len(exam.GetQuestions()))
	assert.Equal(t, newQuestion, (exam.GetQuestions())[0])
}

func TestExam_GetGradeLevel(t *testing.T) {
	// Act
	exam := NewExam(
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)
	// Assertions
	assert.Equal(t, valueobject.GradeLevelThird, exam.GetGradeLevel())
}

func TestMapToExamEntity(t *testing.T) {
	// Act
	exam := MapToExamEntity(
		meta.GetIdUUID(),
		meta.GetCreatedAtTime(),
		meta.GetUpdatedAtTime(),
		meta.IsDeleted(),
		meta.GetDeletedAtTime(),
		"Animal Exam",
		valueobject.GradeLevelThird,
		examTime,
		&questions,
	)
	// Assertions
	assert.Equal(t, meta.GetIdUUID(), exam.GetIdUUID())
	assert.Equal(t, "Animal Exam", exam.GetName())
	result, err := exam.GetQuestionByIndex(0)
	assert.NoError(t, err)
	assert.Equal(t,
		ExamQuestion{
			EntityMetadata:  meta,
			index:           1,
			questionText:    "What animal is known to bark?",
			questionType:    valueobject.QuestionMultipleChoice,
			possiblePoints:  5,
			correctAnswer:   &dog,
			responseOptions: &[]string{"cat", "dog", "fish", "bird"},
		},
		result,
	)
}
