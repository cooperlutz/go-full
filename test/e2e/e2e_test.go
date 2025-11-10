package e2e_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/playwright-community/playwright-go"

	pingpong_api_client_v1 "github.com/cooperlutz/go-full/api/rest/pingpong/v1/client"
)

var (
	serverAddr                   = "http://app.lvh.me:8080"
	pw                           *playwright.Playwright
	browser                      playwright.Browser
	pingpongApiClient            *pingpong_api_client_v1.ClientWithResponses
	defaultBrowserContextOptions = playwright.BrowserNewContextOptions{}
)

func TestMain(m *testing.M) {
	err := setupEndToEndTests()
	if err != nil {
		slog.Error("error setting up e2e tests", slog.String("error", err.Error()))
	}
	// Create a new client instance
	pingpongApiClient, err = pingpong_api_client_v1.NewClientWithResponses(serverAddr + "/pingpong/api/v1")
	if err != nil {
		slog.Error("Error creating pingpong api client:", slog.String("error", err.Error()))
	}

	exitCode := m.Run()

	err = teardownEndToEndTests()
	if err != nil {
		slog.Error("error tearing down e2e tests", slog.String("error", err.Error()))
	}

	os.Exit(exitCode)
}

// setupEndToEndTests - this function runs BEFORE test functions, initializing global variables and encapsulate any functions
// that should run prior to the test functions running
func setupEndToEndTests() error {
	slog.Info("setting up e2e tests...")

	var err error
	pw, err = playwright.Run()
	if err != nil {
		return err
	}

	browser, err = pw.Chromium.Launch()
	if err != nil {
		return err
	}

	return nil
}

// teardownEndToEndTests - this function runs AFTER test functions
func teardownEndToEndTests() error {
	slog.Info("tearing down e2e tests...")

	err := browser.Close()
	if err != nil {
		return err
	}

	err = pw.Stop()
	if err != nil {
		return err
	}

	return nil
}
