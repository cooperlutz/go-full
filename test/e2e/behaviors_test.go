package e2e_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
