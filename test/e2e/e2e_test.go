package e2e_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/playwright-community/playwright-go"

	examination_api_client "github.com/cooperlutz/go-full/api/rest/examination/client"
	examLibrary_api_client_v1 "github.com/cooperlutz/go-full/api/rest/examlibrary/v1/client"
	iam_api_client "github.com/cooperlutz/go-full/api/rest/iam/client"
	pingpong_api_client_v1 "github.com/cooperlutz/go-full/api/rest/pingpong/v1/client"
)

var (
	serverAddr                   = "http://app.lvh.me:8080"
	bearerToken                  string
	pw                           *playwright.Playwright
	browser                      playwright.Browser
	pingpongApiClient            *pingpong_api_client_v1.ClientWithResponses
	examLibraryApiClient         *examLibrary_api_client_v1.ClientWithResponses
	examinationApiClient         *examination_api_client.ClientWithResponses
	iamApiClient                 *iam_api_client.ClientWithResponses
	defaultBrowserContextOptions = playwright.BrowserNewContextOptions{}
)

func TestMain(m *testing.M) {
	err := setupEndToEndTests()
	if err != nil {
		slog.Error("error setting up e2e tests", slog.String("error", err.Error()))
	}

	// Create a new client instance
	pingpongApiClient, err = pingpong_api_client_v1.NewClientWithResponses(serverAddr+"/api/pingpong/v1", pingpong_api_client_v1.WithRequestEditorFn(
		ReqWithBearerToken(bearerToken)),
	)
	if err != nil {
		slog.Error("Error creating pingpong api client:", slog.String("error", err.Error()))
	}

	examLibraryApiClient, err = examLibrary_api_client_v1.NewClientWithResponses(serverAddr+"/api/examlibrary/v1",
		examLibrary_api_client_v1.WithRequestEditorFn(
			ReqWithBearerToken(bearerToken)),
	)
	if err != nil {
		slog.Error("Error creating examLibrary api client:", slog.String("error", err.Error()))
	}

	examinationApiClient, err = examination_api_client.NewClientWithResponses(serverAddr+"/api/examination",
		examination_api_client.WithRequestEditorFn(
			ReqWithBearerToken(bearerToken)),
	)
	if err != nil {
		slog.Error("Error creating examination api client:", slog.String("error", err.Error()))
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
	authentication()

	err := seedTestData()
	if err != nil {
		return err
	}

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

func authentication() {
	client, err := iam_api_client.NewClientWithResponses(serverAddr)
	if err != nil {
		slog.Error("Error creating iam api client:", slog.String("error", err.Error()))
	}
	iamApiClient = client

	_, err = iamApiClient.RegisterUserWithResponse(context.Background(), iam_api_client.RegisterRequest{
		Email:    openapi_types.Email("user@example.com"),
		Password: "SecureP@ssw0rd!",
	})

	resp, err := iamApiClient.LoginUserWithResponse(context.Background(), iam_api_client.LoginRequest{
		Email:    "user@example.com",
		Password: "SecureP@ssw0rd!",
	})
	if err != nil {
		slog.Error("Error logging in test user:", slog.String("error", err.Error()))
	}
	bearerToken = resp.JSON200.AccessToken
}
