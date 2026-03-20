package ai

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("pkg/ai")

// Client wraps the OpenAI client and provides automatic OpenTelemetry instrumentation.
type Client struct {
	openAIClient *openai.Client
	defaultModel string
}

// NewClient creates a new configured AI client. It automatically detects if Azure OpenAI
// env vars are present and routes traffic accordingly, otherwise defaults to public OpenAI.
func NewClient(apiKey string) *Client {
	var config openai.ClientConfig

	azureEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	if azureEndpoint != "" {
		// Using Azure OpenAI
		config = openai.DefaultAzureConfig(apiKey, azureEndpoint)
		if version := os.Getenv("AZURE_OPENAI_API_VERSION"); version != "" {
			config.APIVersion = version
		}
	} else {
		// Public OpenAI Standard
		config = openai.DefaultConfig(apiKey)
	}

	// Figure out the default model/deployment to use across the app
	model := os.Getenv("AZURE_OPENAI_DEPLOYMENT_NAME")
	if model == "" {
		model = os.Getenv("OPENAI_MODEL")
	}

	if model == "" {
		model = "gpt-4o" // Safe fallback
	}

	return &Client{
		openAIClient: openai.NewClientWithConfig(config),
		defaultModel: model,
	}
}

// GenerateCompletion is a basic wrapper around ChatCompletion that automatically creates tracing spans.
func (c *Client) GenerateCompletion(ctx context.Context, prompt string) (string, error) {
	ctx, span := tracer.Start(ctx, "GenerateCompletion", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	req := openai.ChatCompletionRequest{
		Model: c.defaultModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	resp, err := c.openAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		span.RecordError(err)

		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", nil
}
