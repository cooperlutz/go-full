package endtoend

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"

	api_client "github.com/cooperlutz/go-full/api/rest/pingpong/v1/client"
)

var serverAddr = "http://app.lvh.me:8080"

/*
Open Frontend UI

Navigate to PingPong page

Send a ping

- hit the rest api, there should be a ping
- hit the database, there should be a ping
*/

func TestIntegrations(t *testing.T) {
	ctx := context.Background()

	// Create a new client instance
	clientInstance, err := api_client.NewClientWithResponses(serverAddr)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	val := "ping"
	req := api_client.PingPongJSONRequestBody{Message: &val}
	// Call the Ping method
	response, err := clientInstance.PingPongWithResponse(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Ping: %v", err)
	}

	assert.Equal(t, 200, response.StatusCode(), "Expected response message to be 'pong'")
	// Print the response
	// log.Printf("Response from server: %s", response)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(serverAddr + "/dashboard"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// get a basic element and print its text content
	mainContent, err := page.Locator(".link").All()
	if err != nil {
		log.Fatalf("could not get main content: %v", err)
	}
	for i, content := range mainContent {
		text, err := content.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("Main content %d: %s\n", i+1, text)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
