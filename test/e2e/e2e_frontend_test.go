package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Tests in this file do not adhere to encapsulate full end to end functionality, nor do they necessarily follow Behavior Driven Development,
but instead encapsulate tests associated with checking frontend ui pages, components, or other building blocks
as part of the e2e test suite
*/

// tests basic functionality of the /health view rendering
func TestHealthPage(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/health")
	assert.NoError(t, err)

	// Assert
	pageContent, err := page.Locator("#app").All()
	assert.NoError(t, err)
	assert.NotEmpty(t, pageContent)

	for _, content := range pageContent {
		text, err := content.TextContent()
		assert.NoError(t, err)
		assert.Contains(t, text, "if you can see this, things are probably fine")
	}
}

// Tests basic functionality of the /dashboard view rendering
func TestDashboardView(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/dashboard")
	assert.NoError(t, err)

	// Assert
	appGridContent, err := page.Locator("#apps-grid").All()
	assert.NoError(t, err)
	assert.NotEmpty(t, appGridContent)

	for _, content := range appGridContent {
		text, err := content.TextContent()
		assert.NoError(t, err)
		assert.Contains(t, text, "Ping Pong")
	}
}

// tests basic functionality of the sidebar rendering
func TestSidebar(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/")
	assert.NoError(t, err)

	// Assert
	sidebarContent, err := page.Locator("#app").All()
	assert.NoError(t, err)
	assert.NotEmpty(t, sidebarContent)

	for _, content := range sidebarContent {
		text, err := content.TextContent()
		assert.NoError(t, err)
		assert.Contains(t, text, "Dashboard")
		assert.Contains(t, text, "Ping Pong")
	}
}

// tests basic functionality of the footer rendering
func TestFooter(t *testing.T) {
	// Arrange
	_, page := newBrowserContextAndPage(t, defaultBrowserContextOptions)

	// Act
	_, err := page.Goto(serverAddr + "/")
	assert.NoError(t, err)

	// Assert
	footerContent, err := page.Locator(".footer").All()
	assert.NoError(t, err)
	assert.NotEmpty(t, footerContent)

	for i, content := range footerContent {
		text, err := content.TextContent()
		assert.NoError(t, err)
		t.Logf("Footer content %d: %s\n", i+1, text)
	}
}
