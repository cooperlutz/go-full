package e2e_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

func newBrowserContextAndPage(t *testing.T, options playwright.BrowserNewContextOptions) (playwright.BrowserContext, playwright.Page) {
	t.Helper()
	context, err := browser.NewContext(options)
	if err != nil {
		t.Fatalf("could not create new context: %v", err)
	}
	t.Cleanup(func() {
		if err := context.Close(); err != nil {
			t.Errorf("could not close context: %v", err)
		}
	})
	p, err := context.NewPage()
	if err != nil {
		t.Fatalf("could not create new page: %v", err)
	}
	return context, p
}
