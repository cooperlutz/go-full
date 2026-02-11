package grading

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

var FixtureExamQuestions = []*Question{
	NewQuestion(
		QuestionMultipleChoice,
		1,
		"Go",
		utilitee.StrPtr("Go"),
		int32(5),
	),
	NewQuestion(
		QuestionMultipleChoice,
		1,
		"fortran",
		utilitee.StrPtr("Go"),
		int32(5),
	),
	NewQuestion(
		QuestionShortAnswer,
		2,
		"Python",
		nil,
		int32(10),
	),
	NewQuestion(
		QuestionEssay,
		3,
		"Java",
		nil,
		int32(15),
	),
}

func TestExam(t *testing.T) {
	exam := NewExam(
		uuid.MustParse("00000000-0000-0000-0000-000000000123"),
		uuid.MustParse("00000000-0000-0000-0000-000000000123"),
		uuid.MustParse("00000000-0000-0000-0000-000000000123"),
		FixtureExamQuestions,
	)

	// initial state of the exam should be correct
	assert.WithinDuration(t, time.Now(), exam.GetCreatedAtTime(), time.Microsecond*10)
	assert.Nil(t, exam.GetDeletedAtTime())
	assert.False(t, exam.IsDeleted())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.GetStudentId())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.GetExamLibraryExamId())
	assert.Equal(t, uuid.MustParse("00000000-0000-0000-0000-000000000123"), exam.GetExaminationExamId())
	assert.Equal(t, exam.GetStudentIdString(), uuid.MustParse("00000000-0000-0000-0000-000000000123").String())
	assert.Equal(t, FixtureExamQuestions, exam.GetQuestions())
	assert.Equal(t, int32(35), exam.totalPossiblePoints)
	assert.False(t, exam.gradingCompleted)
	assert.Nil(t, exam.totalPointsReceived)
	grade, err := exam.GetGrade()
	assert.Error(t, err)
	assert.Equal(t, "grading not completed", err.Error())

	// Multiple choice questions should be graded automatically
	err = exam.GradeMultipleChoiceQuestions()
	assert.NoError(t, err)

	firstQuestion := exam.GetFirstQuestion()
	assert.NotNil(t, firstQuestion)
	assert.True(t, firstQuestion.graded)
	assert.Equal(t, int32(1), firstQuestion.GetIndex())
	assert.NotNil(t, firstQuestion.GetPointsReceived())
	assert.Equal(t, int32(5), *firstQuestion.GetPointsReceived())

	// from this point on, we want to be able to manually grade the remaining questions
	ungradedQuestions := exam.GetUngradedQuestions()
	assert.Len(t, ungradedQuestions, 2)

	thirdQuestion := ungradedQuestions[0]
	err = thirdQuestion.GradeQuestion(GradeQuestionOption{
		Feedback: "Good job",
		Points:   8,
	})
	assert.NoError(t, err)
	assert.NotNil(t, thirdQuestion.GetPointsReceived())
	assert.Equal(t, int32(8), *thirdQuestion.GetPointsReceived())

	// an error should be returned if points received exceed possible points
	fourthQuestion := ungradedQuestions[1]
	err = fourthQuestion.GradeQuestion(GradeQuestionOption{
		Feedback: "Needs improvement",
		Points:   16,
	})
	assert.Error(t, err)
	assert.Equal(t, "points received cannot exceed possible points", err.Error())

	// grading should not be finalized until all questions are graded
	// given that the third question is still ungraded, grading should not be finalized
	assert.False(t, exam.CheckIfGradingCompletedAndFinalize())

	fourthQuestion = exam.GetQuestionByIndex(4)
	err = fourthQuestion.GradeQuestion(
		GradeQuestionOption{
			Feedback: "this was horrible",
			Points:   1,
		},
	)
	assert.Equal(t, int32(1), *fourthQuestion.GetPointsReceived())
	assert.True(t, exam.CheckIfGradingCompletedAndFinalize())
	assert.True(t, exam.CheckIfGradingCompletedAndFinalize())
	assert.True(t, exam.IsCompleted())
	assert.Equal(t, int32(14), *exam.totalPointsReceived)
	grade, err = exam.GetGrade()
	assert.NoError(t, err)
	assert.Equal(t, float64(40), grade) // 14/35 = 40%
}
