package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/cooperlutz/go-full/tools/modularizer/utils"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

type ModularizerAIGenerator struct {
	llm                        llms.Model
	busDomainModellingTemplate *prompts.PromptTemplate

	directory                  string
	inputFileName              string
	outputFileName             string
	generatedModularizerConfig string
	story                      string
	generationInfo             map[string]any
}

const aiModel = "claude-opus-4-6"

func NewBusinessModeller(dir, inputFileName string) (*ModularizerAIGenerator, error) {

	llm, err := anthropic.New(anthropic.WithModel(aiModel))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	businessModelTemplate, err := os.ReadFile(utils.GetDirectoryOfCurrentFile() + "/templates/business_model_prompt.md")
	if err != nil {
		return nil, fmt.Errorf("reading prompt template: %w", err)
	}

	dmTemplate := prompts.NewPromptTemplate(string(businessModelTemplate), []string{"story"})

	return &ModularizerAIGenerator{
		llm:                        llm,
		busDomainModellingTemplate: &dmTemplate,
		directory:                  dir,
		inputFileName:              inputFileName,
		outputFileName:             "domain_model_" + time.Now().UTC().Format("20060102150405") + ".yaml",
	}, nil
}

func (c *ModularizerAIGenerator) Run(ctx context.Context) error {
	contents, err := os.ReadFile(fmt.Sprintf("%s/%s", c.directory, c.inputFileName))
	if err != nil {
		return fmt.Errorf("reading file contents: %w", err)
	}
	c.story = string(contents)

	prompt, err := c.busDomainModellingTemplate.Format(map[string]any{
		"story": c.story,
	})
	if err != nil {
		return fmt.Errorf("formatting prompt: %w", err)
	}

	response, err := c.llm.GenerateContent(
		ctx,
		[]llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, prompt),
		},
		llms.WithOptions(
			llms.CallOptions{
				// Temperature: 0.7,
				MaxTokens: 32768,
			},
		),
	)
	if err != nil {
		return fmt.Errorf("generating review: %w", err)
	}
	c.generatedModularizerConfig = response.Choices[0].Content
	c.generationInfo = response.Choices[0].GenerationInfo

	// save final response to file
	err = c.saveModularizerConfig()
	if err != nil {
		return fmt.Errorf("saving modularizer config: %w", err)
	}
	slog.Info("llm generation info: ", "input_tokens", c.generationInfo["InputTokens"], "output_tokens", c.generationInfo["OutputTokens"])
	return nil
}

func (c *ModularizerAIGenerator) saveModularizerConfig() error {
	err := c.trimExtraneousContent()
	if err != nil {
		return fmt.Errorf("trimming extraneous content: %w", err)
	}
	c.addContentToOutput()
	err = os.WriteFile(fmt.Sprintf("%s/%s", c.directory, c.outputFileName), []byte(c.generatedModularizerConfig), 0o644)
	if err != nil {
		return fmt.Errorf("writing final review to file: %w", err)
	}
	return nil
}

func (c *ModularizerAIGenerator) addContentToOutput() {
	var outputContent string
	outputContent += "# InputTokens:" + fmt.Sprint(c.generationInfo["InputTokens"].(int)) + "\n"
	outputContent += "# OutputTokens:" + fmt.Sprint(c.generationInfo["OutputTokens"].(int)) + "\n"
	outputContent += "# Model:" + aiModel + "\n"
	outputContent += c.generatedModularizerConfig

	c.generatedModularizerConfig = outputContent
}

func (c *ModularizerAIGenerator) trimExtraneousContent() error {
	// Find the index of the first occurrence of "```yaml"
	startIndex := strings.Index(c.generatedModularizerConfig, "```yaml")
	if startIndex == -1 {
		return fmt.Errorf("start marker not found, unable to trim content")
	}

	// Find the index of the first occurrence of "```" after the start marker
	endIndex := strings.Index(c.generatedModularizerConfig[startIndex+len("```yaml"):], "```")
	if endIndex == -1 {
		return fmt.Errorf("end marker not found, unable to trim content")
	}

	// Extract the content between the markers
	trimmedContent := c.generatedModularizerConfig[startIndex+len("```yaml") : startIndex+len("```yaml")+endIndex]
	c.generatedModularizerConfig = trimmedContent
	return nil
}

func main() {
	// ANTHROPIC_API_KEY=your_api_key needs to be added to .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	modeller, err := NewBusinessModeller("tools/artee/pet_clinic", "business_model.md")
	if err != nil {
		log.Fatalf("Error initializing business modeller: %v", err)
	}

	err = modeller.Run(context.Background())
	if err != nil {
		log.Fatalf("Error running modeller: %v", err)
	}

	log.Println("Domain model generated successfully!")
}
