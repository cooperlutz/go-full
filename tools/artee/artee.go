package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

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

var businessModelPrompt = `
You are an expert at refactoring codebases into Go code.
Adhering to the concepts of Domain-Driven Design, you are responsible for reviewing the code base and identifying the aggregate root entities.

When defining the aggregate root entities, commands, events, and queries for your domain, ensure that you are adhering to the principles of Domain-Driven Design and that you are modeling the core concepts and behaviors of your domain in a way that is clear and maintainable.

Business model
{{.story}}

Step 1: Review the business model and Define the relevant business Domains and bounded contexts which will be broken out into individual microservices. Each Domain should represent a specific area of the business logic and should be designed to be cohesive and loosely coupled with other Domains.

The output must be provided in yaml and conform to the following example structure:
'''yaml
# values must be in snake_case!
modules:
- name: work
  description: Module for managing work items
  bounded_context: The core domain of the system, responsible for managing work items and their associated tasks. This module encapsulates the business logic and rules related to work items, including their creation, assignment, and completion. It serves as the central point of interaction for all operations related to work items within the system.
  responsibility: This module is responsible for managing work items, including their creation, assignment, and completion. It defines the core business logic and rules associated with work items, and serves as the central point of interaction for all operations related to work items within the system.
  defaultQueries: true # if true, a default set of queries (get by id, list all) will be generated for each aggregate
  aggregates:
    - name: work_item
      description: Represents a work item in the system
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
		- name: story_points
		  type: int32
		  optional: true
		- name: tasks
		  optional: true
    - name: task
      description: Represents a task associated with a work item
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
  commands:
    - name: assign_to_user
      description: Assign a work item to a user
      events_emitted:
        - work_availeble_for_assignment
      params:
        - name: user_id
          type: string
          optional: false
        - name: work_item_id
          type: string
          optional: false
    - name: complete_work
      description: Complete a work item
      events_emitted:
        - work_item_completed
      params:
        - name: user_id
          type: string
          optional: false
        - name: work_item_id
          type: string
          optional: false
  events:
    - name: work_item_created
      description: Event triggered when a work item is created
      kind: emitted
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
    - name: work_availeble_for_assignment
      description: Event triggered when a work item is available for assignment
      kind: consumed
      fields:
        - name: work_item_id
          type: string
          optional: false
        - name: name
          type: string
          optional: false
        - name: description
          type: string
          optional: true
  queries:
    - name: find_work_assigned_to_user
      description: Query to find work items assigned to a specific user
      params:
        - name: user_id
          type: string
          optional: false
'''
`

const aiModel = "claude-opus-4-6"

func NewBusinessModeller(dir, inputFileName string) (*ModularizerAIGenerator, error) {
	llm, err := anthropic.New(anthropic.WithModel(aiModel))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dmTemplate := prompts.NewPromptTemplate(businessModelPrompt, []string{"story"})

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
		))
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
	outputContent += "// InputTokens:" + fmt.Sprint(c.generationInfo["InputTokens"].(int)) + "\n"
	outputContent += "// OutputTokens:" + fmt.Sprint(c.generationInfo["OutputTokens"].(int)) + "\n"
	outputContent += "// Model:" + aiModel + "\n"
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
