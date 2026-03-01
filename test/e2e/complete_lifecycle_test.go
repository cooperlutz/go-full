package e2e_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	api_client "github.com/cooperlutz/go-full/api/rest/examination/client"
)

func TestUserStartsAnExamFromExamLibrary(t *testing.T) {
	numberThreeInt32 := int32(3)
	expect := api_client.Exam{
		LibraryExamId:     new("11111111-1111-1111-1111-111111111111"),
		AnsweredQuestions: &numberThreeInt32,
		TotalQuestions:    &numberThreeInt32,
		// StudentId:         ,
		State:  "completed",
		ExamId: "", // we don't know the exam ID ahead of time
		Questions: &[]api_client.Question{
			{
				QuestionIndex: 1,
				Answered:      true,
				QuestionText:  "What is the capital of France?",
				QuestionType:  "multiple-choice",
				ResponseOptions: &[]string{
					"Berlin",
					"Madrid",
					"Paris",
					"Rome",
				},
				ProvidedAnswer: new("Berlin"),
			},
			{
				QuestionIndex:   2,
				Answered:        true,
				QuestionText:    "What is Go?",
				QuestionType:    "short-answer",
				ResponseOptions: nil,
				ProvidedAnswer:  new("A programming language developed by Google."),
			},
			{
				QuestionIndex:   3,
				Answered:        true,
				QuestionText:    "Explain the concept of concurrency.",
				QuestionType:    "essay",
				ResponseOptions: nil,
				ProvidedAnswer:  new("It's about the relationship between space and time."),
			},
		},
	}
	// Arrange
	ctx := context.Background()
	queryCountOfExaminationEvents := func() (int64, error) {
		return countOfQuery("public", "watermill_examination.exam_started")
	}
	queryCountOfExaminationQuestions := func() (int64, error) {
		return countOfQuery("examination", "questions")
	}
	countOfExaminationEventsBefore, err := queryCountOfExaminationEvents()
	examsBefore, err := examinationApiClient.FindAllExamsWithResponse(ctx)
	countOfExaminationExamsBefore := len(*examsBefore.JSON200)
	countOfExaminationQuestionsBefore, err := queryCountOfExaminationQuestions()
	metricNumberOfExamsInProgressBefore, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_in_progress")
	valueMetricNumberOfExamsInProgressBefore := *metricNumberOfExamsInProgressBefore.JSON200.MetricValue
	metricNumberOfExamsCompletedBefore, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_completed")
	valueMetricNumberOfExamsCompletedBefore := *metricNumberOfExamsCompletedBefore.JSON200.MetricValue
	metricNumberOfExamsBeingGradedBefore, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_being_graded")
	valueMetricNumberOfExamsBeingGradedBefore := *metricNumberOfExamsBeingGradedBefore.JSON200.MetricValue
	metricNumberOfExamsGradingCompletedBefore, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_grading_completed")
	valueMetricNumberOfExamsGradingCompletedBefore := *metricNumberOfExamsGradingCompletedBefore.JSON200.MetricValue
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err = page.Goto(serverAddr + "/exam-library/11111111-1111-1111-1111-111111111111")
	modalButtons, err := page.Locator("#start-exam-modal-button").All()
	err = modalButtons[0].Click()
	buttons, err := page.Locator("#confirm-start-exam-button").All()
	err = buttons[0].Click()
	time.Sleep(1 * time.Second)
	url := page.URL()
	examId := strings.Split(url, "/exam/")[1]
	expect.ExamId = examId
	firstQuestionButton, err := page.Locator("#go-to-first-question-button").All()
	err = firstQuestionButton[0].Click()
	time.Sleep(1 * time.Second)
	metricNumberOfExamsInProgressDuring, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_in_progress")
	valueMetricNumberOfExamsInProgressDuring := *metricNumberOfExamsInProgressDuring.JSON200.MetricValue
	mcRadioOpt1, err := page.Locator("#multiple-choice-radio-option-0").All()
	err = mcRadioOpt1[0].Click()
	subQuestionBtn, err := page.Locator("#record-answer-button").All()
	err = subQuestionBtn[0].Click()
	questionNavItem2, err := page.Locator("#question-nav-item-2").All()
	err = questionNavItem2[0].Click()
	time.Sleep(1 * time.Second)
	shortAnswerInput, err := page.Locator("#short-answer-input").All()
	err = shortAnswerInput[0].Fill("A programming language developed by Google.")
	subQuestionBtn2, err := page.Locator("#record-answer-button").All()
	err = subQuestionBtn2[0].Click()
	questionNavItem3, err := page.Locator("#question-nav-item-3").All()
	err = questionNavItem3[0].Click()
	err = page.Locator("#essay-question-input").Fill("It's about the relationship between space and time.")
	subQuestionBtn3, err := page.Locator("#record-answer-button").All()
	err = subQuestionBtn3[0].Click()
	submitExamBtn, err := page.Locator("#exam-submission-button").All()
	err = submitExamBtn[0].Click()
	submissionValidationBtn, err := page.Locator("#confirm-exam-submission-button").All()
	err = submissionValidationBtn[0].Click()

	// Grade the Exam
	_, err = page.Goto(serverAddr + "/grading")
	time.Sleep(2 * time.Second)
	gradingGradeExamBtn, err := page.Locator("#grading-grade-exam-button-0").All()
	err = gradingGradeExamBtn[0].Click()
	time.Sleep(1 * time.Second)
	gradingGradeQuestion2Btn, err := page.Locator("#grading-grade-question-button-1").All()
	err = gradingGradeQuestion2Btn[0].Click()
	err = page.Locator("#grader-comments").Fill("not great")
	err = page.Locator("#points-to-give").Fill("2")
	saveFeedbackBtnQ2, err := page.Locator("#save-feedback-and-points").All()
	err = saveFeedbackBtnQ2[0].Click()

	metricNumberOfExamsBeingGradedAfter, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_being_graded")
	valueMetricNumberOfExamsBeingGradedAfter := *metricNumberOfExamsBeingGradedAfter.JSON200.MetricValue

	time.Sleep(1 * time.Second)
	gradingGradeQuestion3Btn, err := page.Locator("#grading-grade-question-button-2").All()
	err = gradingGradeQuestion3Btn[0].Click()
	err = page.Locator("#grader-comments").Fill("not great")
	err = page.Locator("#points-to-give").Fill("2")
	saveFeedbackBtnQ3, err := page.Locator("#save-feedback-and-points").All()
	err = saveFeedbackBtnQ3[0].Click()

	time.Sleep(2 * time.Second)
	// Assert
	metricNumberOfExamsCompletedAfter, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_completed")
	valueMetricNumberOfExamsCompletedAfter := *metricNumberOfExamsCompletedAfter.JSON200.MetricValue
	metricNumberOfExamsGradingCompletedAfter, err := reportingApiClient.GetMetricWithResponse(ctx, "number_of_exams_grading_completed")
	valueMetricNumberOfExamsGradingCompletedAfter := *metricNumberOfExamsGradingCompletedAfter.JSON200.MetricValue
	countOfExaminationEventsAfter, err := queryCountOfExaminationEvents()
	countOfExaminationQuestionsAfter, err := queryCountOfExaminationQuestions()
	examsAfter, err := examinationApiClient.FindAllExamsWithResponse(ctx)
	countOfExaminationExamsAfter := len(*examsAfter.JSON200)

	assert.Equal(t, countOfExaminationExamsBefore+1, countOfExaminationExamsAfter, "Expected number of exams to increase by 1")
	assert.Equal(t, countOfExaminationEventsBefore+1, countOfExaminationEventsAfter, "Expected number of events to increase by 1")
	assert.Greater(t, countOfExaminationQuestionsAfter, countOfExaminationQuestionsBefore, "Expected number of questions to increase")
	actual, err := examinationApiClient.GetExamWithResponse(ctx, examId)
	assert.Equal(t, *expect.LibraryExamId, *actual.JSON200.LibraryExamId, "Expected LibraryExamId to match")
	assert.Equal(t, *expect.AnsweredQuestions, *actual.JSON200.AnsweredQuestions, "Expected AnsweredQuestions to match")
	assert.Equal(t, *expect.TotalQuestions, *actual.JSON200.TotalQuestions, "Expected TotalQuestions to match")
	assert.Equal(t, expect.State, actual.JSON200.State, "Expected State to match")
	assert.Equal(t, len(*expect.Questions), len(*actual.JSON200.Questions), "Expected number of questions to match")
	for i, expectedQuestion := range *expect.Questions {
		actualQuestion := (*actual.JSON200.Questions)[i]
		assert.Equal(t, expectedQuestion.QuestionIndex, actualQuestion.QuestionIndex, "Question %d: QuestionIndex does not match", expectedQuestion.QuestionIndex)
		assert.Equal(t, expectedQuestion.Answered, actualQuestion.Answered, "Question %d: Answered does not match", expectedQuestion.QuestionIndex)
		assert.Equal(t, expectedQuestion.QuestionText, actualQuestion.QuestionText, "Question %d: QuestionText does not match", expectedQuestion.QuestionIndex)
		assert.Equal(t, expectedQuestion.QuestionType, actualQuestion.QuestionType, "Question %d: QuestionType does not match", expectedQuestion.QuestionIndex)
		assert.Equal(t, expectedQuestion.ResponseOptions, actualQuestion.ResponseOptions, "Question %d: ResponseOptions do not match", expectedQuestion.QuestionIndex)
		assert.Equal(t, *expectedQuestion.ProvidedAnswer, *actualQuestion.ProvidedAnswer, "Question %d: Provided answer does not match expected", expectedQuestion.QuestionIndex)
	}
	assert.Equal(t, valueMetricNumberOfExamsInProgressBefore+float64(1), valueMetricNumberOfExamsInProgressDuring, "Expected number_of_exams_in_progress metric to increase by 1")
	assert.Equal(t, valueMetricNumberOfExamsCompletedBefore+float64(1), valueMetricNumberOfExamsCompletedAfter, "Expected number_of_exams_completed metric to increase by 1 after exam submission")
	assert.Equal(t, valueMetricNumberOfExamsBeingGradedBefore+float64(1), valueMetricNumberOfExamsBeingGradedAfter, "Expected number_of_exams_being_graded metric to increase by 1 after exam submission")
	assert.Equal(t, valueMetricNumberOfExamsGradingCompletedBefore+float64(1), valueMetricNumberOfExamsGradingCompletedAfter, "Expected number_of_exams_grading_completed metric to increase by 1 after exam submission")
	assert.NoError(t, err)
}
