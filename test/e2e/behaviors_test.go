package e2e_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	api_client "github.com/cooperlutz/go-full/api/rest/examination/client"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

/*
Scenario: A user navigates to a page that shouldn't exist

Given:
- a user accesses the frontend ui

When:
- the user navigates to the url `/asdkfjfo2o3falsdflkhjaoishjdfkjnl`

Then:
- the page should contain the 404 Page content
*/
func TestUserAccessesInvalidPage(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/asdkfjfo2o3falsdflkhjaoishjdfkjnl")
	assert.NoError(t, err)

	// Assert
	pageContent, err := page.Locator("#app").All()
	assert.NoError(t, err)
	assert.NotEmpty(t, pageContent)

	for _, content := range pageContent {
		text, err := content.TextContent()
		assert.NoError(t, err)
		assert.Contains(t, text, "404")
	}
}

/*
Scenario: A user inputs and sends a valid `ping` via the Ping Pong application UI, which creates a new ping in the backend system

Given:
- a user accesses the ping pong application UI

When:
- the user manually inputs "ping" into the input field
- and the user clicks the send button

Then:
- a new ping is created in the backend system
*/
func TestUserInputsAndSendsAValidPing(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)
	currentPings, err := pingpongApiClient.GetPingsWithResponse(ctx)
	numPingsBeforeAction := len(*currentPings.JSON200.Pingpongs)

	// Act
	_, err = page.Goto(serverAddr + "/ping-pong/app")
	assert.NoError(t, err, "Error navigating to Ping Pong app page: %v", err)
	err = page.Locator("#pingpong-input").Fill("ping")
	buttons, err := page.Locator("#send-button").All()
	assert.NoError(t, err, "Error locating send button: %v", err)
	err = buttons[0].Click()
	assert.NoError(t, err, "Error clicking send button: %v", err)
	assert.NotEmpty(t, buttons)

	// Assert
	time.Sleep(1 * time.Second)
	pingsAfterAction, err := pingpongApiClient.GetPingsWithResponse(ctx)
	assert.NoError(t, err, "Error getting pings after action: %v", err)
	numPingsAfterAction := len(*pingsAfterAction.JSON200.Pingpongs)
	assert.Equal(t, numPingsBeforeAction+1, numPingsAfterAction, "Expected number of pings to increase by 1")
}

/*
Scenario: A user inputs and sends an invalid input via the Ping Pong application UI, this should NOT make any changes in the backend system

Given:
- a user accesses the ping pong application UI

When:
- the user manually inputs "jqlerjhfljkohqelkrjhglkjahsdkjfhlakjhsdljfhlakjdshflkjashsdf" into the input field
- and the user clicks the send button

Then:
- no new ping is created in the backend system
*/
func TestUserInputsAndSendsAnInvalidPing(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)
	currentPings, err := pingpongApiClient.GetPingsWithResponse(ctx)
	numPingsBefore := len(*currentPings.JSON200.Pingpongs)

	// Act
	_, err = page.Goto(serverAddr + "/ping-pong/app")
	assert.NoError(t, err, "Error navigating to Ping Pong app page: %v", err)
	err = page.Locator("#pingpong-input").Fill("jqlerjhfljkohqelkrjhglkjahsdkjfhlakjhsdljfhlakjdshflkjashsdf")
	buttons, err := page.Locator("#send-button").All()
	assert.NoError(t, err, "Error locating send button: %v", err)
	err = buttons[0].Click()
	assert.NoError(t, err, "Error clicking send button: %v", err)
	assert.NotEmpty(t, buttons)

	// Assert
	time.Sleep(1 * time.Second)
	pingsAfterAction, err := pingpongApiClient.GetPingsWithResponse(ctx)
	assert.NoError(t, err, "Error getting pings after action: %v", err)
	numPingsAfterAction := len(*pingsAfterAction.JSON200.Pingpongs)
	assert.Equal(t, numPingsBefore, numPingsAfterAction, "Expected number of pings to remain the same")
}

/*
Scenario: A user accesses the Exam Library application UI and is able to view the list of exams available

Given:
- a user accesses the system UI

When:
- the user navigates to the exam library page

Then:
- a table of exams is displayed to the user
*/
func TestUserViewsExamLibrary(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/exam-library")
	assert.NoError(t, err, "Error navigating to Exam Library page: %v", err)

	// Assert
	tableLocator := page.Locator("#exam-library-table")
	assert.NotNil(t, tableLocator, "Exam Library table locator should not be nil")

	tableExists, err := tableLocator.IsVisible()
	assert.NoError(t, err, "Error checking visibility of Exam Library table: %v", err)
	assert.True(t, tableExists, "Exam Library table should be visible on the page")
}

/*
Scenario: A user accesses an Exam from the Exam Library application UI and is able to start taking the exam

Given:
- a student user accesses the system UI

When:
- the user navigates to an exam detail page
- and the user clicks the "Start Exam" button

Then:
- the exam is started for the student

Implementation Notes:
1. Go to /exam-library/{examId} within the frontend UI
2. Click the "Start Exam" button
3. An exam should be created within the examination system for the student, verified via the API + database check
4. An exam started event should be published to the message bus, verified via querying the 'watermill_examination' table in the database
*/
func TestUserStartsAnExamFromExamLibrary(t *testing.T) {
	randomStudentId := uuid.New().String()
	numberThreeInt32 := int32(3)
	expect := api_client.Exam{
		LibraryExamId:     utilitee.StrPtr("11111111-1111-1111-1111-111111111111"),
		AnsweredQuestions: &numberThreeInt32,
		TotalQuestions:    &numberThreeInt32,
		StudentId:         randomStudentId,
		Completed:         true,
		ExamId:            "", // we don't know the exam ID ahead of time
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
				ProvidedAnswer: utilitee.StrPtr("Berlin"),
			},
			{
				QuestionIndex:   2,
				Answered:        true,
				QuestionText:    "What is Go?",
				QuestionType:    "short-answer",
				ResponseOptions: nil,
				ProvidedAnswer:  utilitee.StrPtr("A programming language developed by Google."),
			},
			{
				QuestionIndex:   3,
				Answered:        true,
				QuestionText:    "Explain the concept of concurrency.",
				QuestionType:    "essay",
				ResponseOptions: nil,
				ProvidedAnswer:  utilitee.StrPtr("It's about the relationship between space and time."),
			},
		},
	}
	// Arrange
	ctx := context.Background()
	queryCountOfExaminationEvents := func() (int64, error) {
		return countOfQuery("public", "watermill_examination.examstarted")
	}
	queryCountOfExaminationQuestions := func() (int64, error) {
		return countOfQuery("examination", "questions")
	}
	countOfExaminationEventsBefore, err := queryCountOfExaminationEvents()
	examsBefore, err := examinationApiClient.GetAvailableExamsWithResponse(ctx)
	countOfExaminationExamsBefore := len(*examsBefore.JSON200)
	countOfExaminationQuestionsBefore, err := queryCountOfExaminationQuestions()
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err = page.Goto(serverAddr + "/exam-library/11111111-1111-1111-1111-111111111111")
	modalButtons, err := page.Locator("#start-exam-modal-button").All()
	err = modalButtons[0].Click()
	err = page.Locator("#student-id-input").Fill(randomStudentId)
	buttons, err := page.Locator("#confirm-start-exam-button").All()
	err = buttons[0].Click()
	time.Sleep(1 * time.Second)
	url := page.URL()
	examId := strings.Split(url, "/exam/")[1]
	expect.ExamId = examId
	firstQuestionButton, err := page.Locator("#go-to-first-question-button").All()
	err = firstQuestionButton[0].Click()
	time.Sleep(1 * time.Second)
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

	// Assert
	countOfExaminationEventsAfter, err := queryCountOfExaminationEvents()
	countOfExaminationQuestionsAfter, err := queryCountOfExaminationQuestions()
	examsAfter, err := examinationApiClient.GetAvailableExamsWithResponse(ctx)
	countOfExaminationExamsAfter := len(*examsAfter.JSON200)

	assert.Equal(t, countOfExaminationExamsBefore+1, countOfExaminationExamsAfter)
	assert.Equal(t, countOfExaminationEventsBefore+1, countOfExaminationEventsAfter)
	assert.Greater(t, countOfExaminationQuestionsAfter, countOfExaminationQuestionsBefore)
	actual, err := examinationApiClient.GetExamWithResponse(ctx, examId)
	assert.Equal(t, *expect.LibraryExamId, *actual.JSON200.LibraryExamId)
	assert.Equal(t, expect.StudentId, actual.JSON200.StudentId)
	assert.Equal(t, *expect.AnsweredQuestions, *actual.JSON200.AnsweredQuestions)
	assert.Equal(t, *expect.TotalQuestions, *actual.JSON200.TotalQuestions)
	assert.Equal(t, expect.Completed, actual.JSON200.Completed)
	assert.Equal(t, len(*expect.Questions), len(*actual.JSON200.Questions))
	for i, expectedQuestion := range *expect.Questions {
		actualQuestion := (*actual.JSON200.Questions)[i]
		assert.Equal(t, expectedQuestion.QuestionIndex, actualQuestion.QuestionIndex)
		assert.Equal(t, expectedQuestion.Answered, actualQuestion.Answered)
		assert.Equal(t, expectedQuestion.QuestionText, actualQuestion.QuestionText)
		assert.Equal(t, expectedQuestion.QuestionType, actualQuestion.QuestionType)
		assert.Equal(t, expectedQuestion.ResponseOptions, actualQuestion.ResponseOptions)
		assert.Equal(t, *expectedQuestion.ProvidedAnswer, *actualQuestion.ProvidedAnswer)
	}
	assert.NoError(t, err)
}
