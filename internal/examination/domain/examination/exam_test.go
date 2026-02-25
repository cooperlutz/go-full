package examination

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var FixtureExamQuestions = []*Question{
	NewQuestion(1, "What is 2 + 2?", QuestionMultipleChoice, &[]string{"3", "4", "5"}),
	NewQuestion(2, "What is the capital of France?", QuestionShortAnswer, nil),
	NewQuestion(3, "Write an essay", QuestionEssay, nil),
}

func TestExam(t *testing.T) {
	exam := NewExam(uuid.MustParse("00000000-0000-0000-0000-000000000123"), uuid.MustParse("00000000-0000-0000-0000-000000000123"), 3600, FixtureExamQuestions)
	assert.WithinDuration(t, time.Now(), exam.GetCreatedAtTime(), time.Microsecond*10)
	assert.Nil(t, exam.GetStartedAtTime())
	assert.Nil(t, exam.GetCompletedAtTime())
	assert.Nil(t, exam.GetTimeOfTimeLimit())
	assert.Nil(t, exam.GetDeletedAtTime())
	assert.False(t, exam.IsDeleted())
	assert.Equal(t, int64(3600), exam.GetTimeLimitSeconds())
	assert.Equal(t, StateNotStarted, exam.GetState())

	err := exam.StartExam()
	assert.Nil(t, err)

	assert.Equal(t, StateInProgress, exam.GetState())
	assert.WithinDuration(t, time.Now(), *exam.GetStartedAtTime(), time.Microsecond*10)
	assert.Nil(t, exam.GetCompletedAtTime())
	assert.Nil(t, exam.GetDeletedAtTime())
	assert.False(t, exam.IsDeleted())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.GetStudentIdUUID())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.GetLibraryExamIdUUID())
	assert.Equal(t, FixtureExamQuestions, exam.GetQuestions())

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
	assert.WithinDuration(t, time.Now(), exam.GetUpdatedAtTime(), time.Microsecond*10)
	assert.WithinDuration(t, firstQuestion.GetUpdatedAtTime(), time.Now(), time.Microsecond*10)
	assert.Nil(t, err)
	answeredQuestion := exam.GetQuestionByIndex(1)
	assert.Equal(t, "4", *answeredQuestion.GetProvidedAnswer())

	err = exam.Submit()
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

	err = exam.Submit()
	assert.Nil(t, err)
	assert.Equal(t, StateCompleted, exam.GetState())
	assert.True(t, exam.IsCompleted())
	assert.NotNil(t, exam.GetCompletedAtTime())
	assert.WithinDuration(t, time.Now(), *exam.GetCompletedAtTime(), time.Microsecond*10)
	assert.WithinDuration(t, time.Now(), exam.GetUpdatedAtTime(), time.Microsecond*10)
}

func TestExamStateFromString(t *testing.T) {
	state, err := ExamStateFromString("not-started")
	assert.Nil(t, err)
	assert.Equal(t, StateNotStarted, state)

	state, err = ExamStateFromString("in-progress")
	assert.Nil(t, err)
	assert.Equal(t, StateInProgress, state)

	state, err = ExamStateFromString("completed")
	assert.Nil(t, err)
	assert.Equal(t, StateCompleted, state)

	_, err = ExamStateFromString("INVALID STATE")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidExamState{})
	assert.Equal(t, "invalid exam state", err.Error())
}

func TestQuestionTypeFromString(t *testing.T) {
	questionType, err := QuestionTypeFromString("multiple-choice")
	assert.Nil(t, err)
	assert.Equal(t, QuestionMultipleChoice, questionType)

	questionType, err = QuestionTypeFromString("short-answer")
	assert.Nil(t, err)
	assert.Equal(t, QuestionShortAnswer, questionType)

	questionType, err = QuestionTypeFromString("essay")
	assert.Nil(t, err)
	assert.Equal(t, QuestionEssay, questionType)

	_, err = QuestionTypeFromString("INVALID TYPE")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidQuestionType{})
	assert.Equal(t, "invalid question type", err.Error())
}

func TestQuestion(t *testing.T) {
	question := NewQuestion(1, "What is 2 + 2?", QuestionMultipleChoice, &[]string{"3", "4", "5"})
	assert.Equal(t, int32(1), question.GetIndex())
	assert.Equal(t, "What is 2 + 2?", question.GetQuestionText())
	assert.Equal(t, QuestionMultipleChoice, question.GetQuestionType())
	assert.Equal(t, []string{"3", "4", "5"}, *question.GetResponseOptions())
}

func TestQuestionType_Int(t *testing.T) {
	assert.Equal(t, 0, QuestionMultipleChoice.Int())
	assert.Equal(t, 1, QuestionEssay.Int())
	assert.Equal(t, 2, QuestionShortAnswer.Int())
}

func TestQuestionType_String(t *testing.T) {
	assert.Equal(t, "multiple-choice", QuestionMultipleChoice.String())
	assert.Equal(t, "essay", QuestionEssay.String())
	assert.Equal(t, "short-answer", QuestionShortAnswer.String())
}
