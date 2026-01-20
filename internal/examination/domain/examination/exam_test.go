package examination

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var FixtureExamQuestions = []*Question{
	NewQuestion(1, "What is 2 + 2?", QuestionMultipleChoice, "", &[]string{"3", "4", "5"}),
	NewQuestion(2, "What is the capital of France?", QuestionShortAnswer, "", nil),
	NewQuestion(3, "Write an essay", QuestionEssay, "", nil),
}

func TestExam(t *testing.T) {
	exam := NewExam(uuid.MustParse("00000000-0000-0000-0000-000000000123"), FixtureExamQuestions)
	assert.WithinDuration(t, time.Now(), exam.GetCreatedAtTime(), time.Microsecond*5)
	err := exam.StartExam()
	assert.Nil(t, err)

	assert.WithinDuration(t, time.Now(), *exam.GetStartedAtTime(), time.Microsecond*5)
	assert.Nil(t, exam.GetCompletedAtTime())
	assert.Nil(t, exam.GetDeletedAtTime())
	assert.False(t, exam.IsDeleted())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.studentId)
	assert.Equal(t, FixtureExamQuestions, exam.questions)

	err = exam.StartExam()
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrExamAlreadyStarted{})
	assert.Equal(t, "exam has already been started", err.Error())

	firstQuestion := exam.GetFirstQuestion()
	assert.NotNil(t, firstQuestion)
	assert.Equal(t, int32(1), firstQuestion.GetIndex())
	assert.Equal(t, "What is 2 + 2?", firstQuestion.GetQuestionText())
	err = exam.AnswerQuestion(1, "WRONG ANSWER")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidAnswer{})
	assert.Equal(t, "invalid answer provided", err.Error())

	err = exam.AnswerQuestion(1, "4")
	assert.WithinDuration(t, time.Now(), exam.GetUpdatedAtTime(), time.Microsecond*5)
	assert.WithinDuration(t, firstQuestion.GetUpdatedAtTime(), time.Now(), time.Microsecond*5)
	assert.Nil(t, err)
	answeredQuestion := exam.GetQuestionByIndex(1)
	assert.Equal(t, "4", answeredQuestion.GetProvidedAnswer())

	err = exam.FinishExam()
	assert.Error(t, err)
	assert.False(t, exam.IsCompleted())
	assert.Nil(t, exam.GetCompletedAtTime())

	question2, err := exam.NextQuestion(answeredQuestion.GetIndex())
	assert.Nil(t, err)
	assert.NotNil(t, question2)
	assert.Equal(t, int32(2), question2.GetIndex())
	err = exam.AnswerQuestion(2, "Paris")
	assert.Nil(t, err)

	question3, err := exam.NextQuestion(2)
	err = exam.AnswerQuestion(3, "this is my essay response")
	assert.Nil(t, err)

	question3, err = exam.NextQuestion(3)
	assert.Nil(t, question3)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoMoreQuestions{})
	assert.Equal(t, "no more questions available", err.Error())

	err = exam.FinishExam()
	assert.Nil(t, err)
	assert.True(t, exam.IsCompleted())
	assert.NotNil(t, exam.GetCompletedAtTime())
	assert.WithinDuration(t, time.Now(), *exam.GetCompletedAtTime(), time.Microsecond*5)
	assert.WithinDuration(t, time.Now(), exam.GetUpdatedAtTime(), time.Microsecond*5)
}
