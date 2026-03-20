package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"go.yaml.in/yaml/v2"

	"github.com/cooperlutz/go-full/pkg/ai"
	"github.com/cooperlutz/go-full/tools/modularizer"
	"github.com/cooperlutz/go-full/tools/modularizer/module"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: make generate-poc prompt=\"path/to/business_case.md\"")
	}

	promptFile := os.Args[1]

	content, err := os.ReadFile(promptFile)
	if err != nil {
		log.Fatalf("Error reading prompt file %s: %v\n", promptFile, err)
	}

	fmt.Println("=== AI PoC Factory Generator ===")
	fmt.Printf("[1/6] Reading business case from: %s\n", promptFile)

	aiClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Phase 1: Ask the AI to generate a modularizer YAML config
	fmt.Println("[2/6] Generating module configuration via AI...")

	modConfig, aiIntegration, err := generateModuleConfig(aiClient, string(content))
	if err != nil {
		log.Fatalf("Failed to generate module config: %v\n", err)
	}

	moduleName := modConfig.Modules[0].Name.Flat()
	fmt.Printf("       Module name: %s\n", moduleName)
	fmt.Printf("       AI integration: %v\n", aiIntegration)
	fmt.Printf("       Aggregates: %d\n", len(modConfig.Modules[0].Aggregates))
	fmt.Printf("       Commands: %d\n", len(modConfig.Modules[0].Commands))
	fmt.Printf("       Events: %d\n", len(modConfig.Modules[0].Events))

	// Save the generated YAML for auditing
	if err := saveGeneratedConfig(modConfig, moduleName); err != nil {
		log.Printf("Warning: could not save generated config: %v\n", err)
	}

	// Phase 2: Invoke the modularizer to scaffold the full module
	fmt.Println("[3/6] Scaffolding module via modularizer...")

	if err := runModularizer(modConfig); err != nil {
		log.Fatalf("Modularizer failed: %v\n", err)
	}

	// Phase 3: If AI integration is needed, inject the AI client dependency
	if aiIntegration {
		fmt.Println("[4/6] Injecting AI client dependency into module...")
		if err := injectAIDependency(moduleName); err != nil {
			log.Printf("Warning: AI dependency injection failed: %v\n", err)
			fmt.Println("       You will need to manually add pkg/ai to the module.")
		}
	} else {
		fmt.Println("[4/6] No AI runtime dependency needed, skipping injection.")
	}

	// Phase 4: Auto-wire into app/app.go
	fmt.Println("[5/6] Wiring module into app/app.go...")
	if err := wireIntoApp(moduleName, aiIntegration); err != nil {
		log.Printf("Warning: auto-wiring failed: %v\n", err)
		fmt.Println("       You will need to manually wire the module into app/app.go.")
	}

	// Phase 5: Update .mockery.yml and run codegen
	fmt.Println("[6/6] Updating mockery config and running codegen...")
	if err := updateMockeryConfig(moduleName); err != nil {
		log.Printf("Warning: mockery config update failed: %v\n", err)
	}
	if err := runGoModTidy(); err != nil {
		log.Printf("Warning: go mod tidy failed: %v\n", err)
	}

	fmt.Println("\n=== PoC Generation Complete ===")
	fmt.Printf("Module '%s' has been scaffolded and wired into the application.\n", moduleName)
	fmt.Println("\nNext steps:")
	fmt.Println("  1. Review the generated code under internal/" + moduleName + "/")
	fmt.Println("  2. Uncomment and customize the entity fields in the domain and SQL files")
	fmt.Println("  3. Wire the frontend router (if applicable)")
	fmt.Println("  4. Run: make gen")
	fmt.Println("  5. Run: make test")
}

// generateModuleConfig calls the AI to produce a modularizer-compatible YAML config
// from a business case description. Returns the parsed config and whether AI integration is needed.
func generateModuleConfig(aiClient *ai.Client, businessCase string) (module.ModuleConfig, bool, error) {
	systemPrompt := `You are an expert backend Go module architect. Given a business case description, you must output a YAML configuration that will be used to scaffold a Go module.

Output ONLY a single YAML codeblock (enclosed in triple backticks with 'yaml' language tag). No other text.

The YAML must follow this exact schema:

` + "```yaml" + `
ai_integration: false  # set to true ONLY if the business case requires calling an LLM/AI at runtime
modules:
- name: module_name  # snake_case, short, descriptive
  description: "Brief description of the module"
  defaultQueries: true  # generates find_all and find_one queries per aggregate
  aggregates:
    - name: entity_name  # snake_case
      description: "What this entity represents"
      fields:
        - name: field_name  # snake_case
          type: string  # allowed: string, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, time, date, uuid, object
          optional: false  # true if nullable
          list: false  # true if array
  commands:
    - name: command_name  # snake_case
      description: "What this command does"
      events_emitted:
        - event_name  # references an event below
      params:
        - name: param_name
          type: string
          optional: false
  events:
    - name: event_name  # snake_case
      description: "What triggers this event"
      kind: emitted  # emitted or consumed
      fields:
        - name: field_name
          type: string
          optional: false
` + "```" + `

RULES:
- All names MUST be in snake_case
- Field types must be one of: string, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, time, date, uuid, object
- The module name should be short and descriptive (e.g. "startup_rater", "pet_registry")
- Include at least one aggregate, one command, and one event
- Set ai_integration to true ONLY if the business case explicitly requires calling an AI/LLM service at runtime (e.g. "rate using AI", "generate with LLM", etc.)
- Do NOT include standard CRUD metadata fields (id, created_at, updated_at, deleted, deleted_at) in the aggregate fields - those are added automatically
- DO include all business-specific fields

Here is an example for a veterinary clinic pet registry:

` + "```yaml" + `
ai_integration: false
modules:
- name: pet_registry
  description: Module for managing pet registrations in a veterinary clinic system
  defaultQueries: true
  aggregates:
    - name: pet
      description: Represents a pet registered in the veterinary clinic system
      fields:
        - name: name
          type: string
          optional: false
        - name: species
          type: string
          optional: false
        - name: age
          type: float32
          optional: true
    - name: owner
      description: Represents the owner of a pet
      fields:
        - name: name
          type: string
          optional: false
  commands:
    - name: register_pet
      description: Register a new pet in the system
      events_emitted:
        - pet_registered
      params:
        - name: name
          type: string
          optional: false
        - name: species
          type: string
          optional: false
        - name: age
          type: float32
          optional: true
  events:
    - name: pet_registered
      description: Event triggered when a new pet is registered
      kind: emitted
      fields:
        - name: name
          type: string
          optional: false
        - name: species
          type: string
          optional: false
` + "```"

	fullPrompt := systemPrompt + "\n\nBusiness Case:\n" + businessCase

	resp, err := aiClient.GenerateCompletion(context.Background(), fullPrompt)
	if err != nil {
		return module.ModuleConfig{}, false, fmt.Errorf("AI call failed: %w", err)
	}

	yamlContent := extractCodeBlock(resp, "yaml")
	if yamlContent == "" {
		return module.ModuleConfig{}, false, fmt.Errorf("no YAML codeblock found in AI response.\nRaw response:\n%s", resp)
	}

	// Parse the ai_integration flag separately since it's not part of ModuleConfig
	aiIntegration := false
	var rawConfig map[string]any
	if err := yaml.Unmarshal([]byte(yamlContent), &rawConfig); err == nil {
		if val, ok := rawConfig["ai_integration"]; ok {
			if boolVal, ok := val.(bool); ok {
				aiIntegration = boolVal
			}
		}
	}

	// Parse into ModuleConfig
	var modConfig module.ModuleConfig
	if err := yaml.Unmarshal([]byte(yamlContent), &modConfig); err != nil {
		return module.ModuleConfig{}, false, fmt.Errorf("failed to parse YAML config: %w\nYAML content:\n%s", err, yamlContent)
	}

	if len(modConfig.Modules) == 0 {
		return module.ModuleConfig{}, false, fmt.Errorf("AI generated empty modules list.\nYAML content:\n%s", yamlContent)
	}

	// Validate
	for _, mod := range modConfig.Modules {
		if err := mod.Validate(); err != nil {
			return module.ModuleConfig{}, false, fmt.Errorf("validation failed: %w\nYAML content:\n%s", err, yamlContent)
		}
	}

	return modConfig, aiIntegration, nil
}

// saveGeneratedConfig writes the AI-generated config to a file for auditing/regeneration.
func saveGeneratedConfig(config module.ModuleConfig, moduleName string) error {
	outDir := "tools/pocgen/generated"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	outPath := filepath.Join(outDir, moduleName+".yaml")
	return os.WriteFile(outPath, data, 0644)
}

// runModularizer creates a Modularizer and runs it to scaffold the full module.
func runModularizer(config module.ModuleConfig) error {
	modularizers, err := modularizer.FromConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create modularizer: %w", err)
	}

	for _, m := range modularizers {
		if err := m.CreateModule(); err != nil {
			return fmt.Errorf("failed to create module: %w", err)
		}
	}

	return nil
}

// injectAIDependency modifies the generated module.go and app.go files to include
// the pkg/ai.Client dependency, enabling AI calls at runtime.
func injectAIDependency(moduleName string) error {
	// Patch module.go - add ai.Client parameter
	modulePath := filepath.Join("internal", moduleName, "module.go")
	moduleContent, err := os.ReadFile(modulePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", modulePath, err)
	}

	moduleStr := string(moduleContent)

	// Add ai import
	moduleStr = strings.Replace(moduleStr,
		`"github.com/jackc/pgx/v5/pgxpool"`,
		`"github.com/jackc/pgx/v5/pgxpool"`+"\n\n\t"+`"github.com/cooperlutz/go-full/pkg/ai"`,
		1)

	// Add aiClient parameter to NewModule function signature
	moduleStr = strings.Replace(moduleStr,
		"pubSub *eeventdriven.BasePgsqlPubSubProcessor,\n)",
		"pubSub *eeventdriven.BasePgsqlPubSubProcessor,\n\taiClient *ai.Client,\n)",
		1)

	if err := os.WriteFile(modulePath, []byte(moduleStr), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", modulePath, err)
	}

	return nil
}

// wireIntoApp programmatically edits app/app.go to import, instantiate, and mount the new module.
func wireIntoApp(moduleName string, aiIntegration bool) error {
	appPath := "app/app.go"
	content, err := os.ReadFile(appPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", appPath, err)
	}

	appStr := string(content)
	pascalName := snakeToPascal(moduleName)
	camelName := snakeToCamel(moduleName)

	importPath := fmt.Sprintf(`"github.com/cooperlutz/go-full/internal/%s"`, moduleName)

	// Step 1: Add import (if not already present)
	if !strings.Contains(appStr, importPath) {
		// Insert before the last import in the cooperlutz group
		insertAfter := `"github.com/cooperlutz/go-full/pkg/workerbee"`
		if !strings.Contains(appStr, insertAfter) {
			// Fallback: insert before closing paren of imports
			insertAfter = `"github.com/cooperlutz/go-full/pkg/hteeteepee"`
		}
		appStr = strings.Replace(appStr, insertAfter, insertAfter+"\n\t"+importPath, 1)
	}

	// Also add ai import if needed and not already present
	aiImport := `"github.com/cooperlutz/go-full/pkg/ai"`
	if aiIntegration && !strings.Contains(appStr, aiImport) {
		insertAfter := `"github.com/cooperlutz/go-full/pkg/workerbee"`
		if !strings.Contains(appStr, insertAfter) {
			insertAfter = `"github.com/cooperlutz/go-full/pkg/hteeteepee"`
		}
		appStr = strings.Replace(appStr, insertAfter, insertAfter+"\n\t"+aiImport, 1)
	}

	// Step 2: Add module initialization block
	// Find the last module init block and append after it
	var moduleInitBlock string
	if aiIntegration {
		// Add AI client init if not already present
		if !strings.Contains(appStr, "aiClient := ai.NewClient") {
			moduleInitBlock += "\t// AI Client\n"
			moduleInitBlock += "\taiClient := ai.NewClient(os.Getenv(\"OPENAI_API_KEY\"))\n\n"
		}
		moduleInitBlock += fmt.Sprintf("\t// %s\n", pascalName)
		moduleInitBlock += fmt.Sprintf("\t%sModule, err := %s.NewModule(\n", camelName, moduleName)
		moduleInitBlock += "\t\tconn,\n"
		moduleInitBlock += "\t\tpubSub,\n"
		moduleInitBlock += "\t\taiClient,\n"
		moduleInitBlock += "\t)\n"
		moduleInitBlock += "\tif err != nil {\n"
		moduleInitBlock += "\t\tos.Exit(1)\n"
		moduleInitBlock += "\t}\n"
	} else {
		moduleInitBlock = fmt.Sprintf("\t// %s\n", pascalName)
		moduleInitBlock += fmt.Sprintf("\t%sModule, err := %s.NewModule(\n", camelName, moduleName)
		moduleInitBlock += "\t\tconn,\n"
		moduleInitBlock += "\t\tpubSub,\n"
		moduleInitBlock += "\t)\n"
		moduleInitBlock += "\tif err != nil {\n"
		moduleInitBlock += "\t\tos.Exit(1)\n"
		moduleInitBlock += "\t}\n"
	}

	// Insert before the "Protected REST API Controller" section
	marker := "\t/* -----------------------------------------------------------------------------------\n\tProtected REST API Controller Initialization:"
	if strings.Contains(appStr, marker) && !strings.Contains(appStr, camelName+"Module, err :=") {
		appStr = strings.Replace(appStr, marker, moduleInitBlock+"\n"+marker, 1)
	}

	// Step 3: Add route mount
	routeMount := fmt.Sprintf("\tprotectedRestApiRouter.Mount(\n\t\t\"/%s\",\n\t\t%sModule.RestApi,\n\t)", moduleName, camelName)

	// Insert before the "Mount Public Routes" section
	publicRoutesMarker := "\t/* -----------------------------------------------------------------------------------\n\tMount Public Routes"
	if strings.Contains(appStr, publicRoutesMarker) && !strings.Contains(appStr, `"/`+moduleName+`"`) {
		appStr = strings.Replace(appStr, publicRoutesMarker, routeMount+"\n\n"+publicRoutesMarker, 1)
	}

	return os.WriteFile(appPath, []byte(appStr), 0644)
}

// updateMockeryConfig adds entries for the new module to .mockery.yml
func updateMockeryConfig(moduleName string) error {
	mockeryPath := ".mockery.yml"
	content, err := os.ReadFile(mockeryPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", mockeryPath, err)
	}

	mockeryStr := string(content)
	pascalName := snakeToPascal(moduleName)

	// Check if already present
	if strings.Contains(mockeryStr, moduleName+"_mocks") {
		return nil // already configured
	}

	// The modularizer uses hexagonal architecture (adapters/outbound), so mock that
	mockeryEntry := fmt.Sprintf(`
# %s Module
  github.com/cooperlutz/go-full/internal/%s/adapters/outbound:
    config:
      filename: mocks_%s_adapters_outbound.gen.go
      all: true
      dir: test/mocks/%s
      pkgname: %s_mocks
      include-auto-generated: true`, pascalName, moduleName, moduleName, moduleName, moduleName)

	mockeryStr += mockeryEntry + "\n"

	return os.WriteFile(mockeryPath, []byte(mockeryStr), 0644)
}

// runGoModTidy runs go mod tidy to ensure all dependencies are resolved.
func runGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// extractCodeBlock pulls the actual code content out of the LLM's markdown fences
func extractCodeBlock(text string, lang string) string {
	re := regexp.MustCompile("(?is)```" + lang + "\\s*(.*?)\\s*```")
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

// snakeToPascal converts a snake_case string to PascalCase
func snakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// snakeToCamel converts a snake_case string to camelCase
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i > 0 && len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
