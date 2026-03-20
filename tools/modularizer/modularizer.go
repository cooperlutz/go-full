package modularizer

import (
	"encoding/json"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"go.yaml.in/yaml/v2"

	"github.com/cooperlutz/go-full/tools/modularizer/module"
	"github.com/cooperlutz/go-full/tools/modularizer/templates"
	"github.com/cooperlutz/go-full/tools/modularizer/utils"
)

type Modularizer struct {
	templateData module.Module
}

func NewModularizer(mod module.Module) (*Modularizer, error) {
	if err := mod.Validate(); err != nil {
		return nil, err
	}
	return &Modularizer{
		templateData: mod,
	}, nil
}

func FromConfig(cfg module.ModuleConfig) ([]*Modularizer, error) {
	var modularizers []*Modularizer
	for _, mod := range cfg.Modules {
		modularizer, err := NewModularizer(mod)
		if err != nil {
			return nil, err
		}
		modularizers = append(modularizers, modularizer)
	}
	return modularizers, nil
}

// CreateModule creates a new module based on the modularizer's template data, including the directory structure,
// module file, entity files, repository files, application file, command files, query files,
// event files, Postgres adapter files, database migration files, REST API files, OpenAPI schema file,
// and updates to the SQLC configuration and OpenAPI Generator configuration,
// and returns an error if there was an issue creating any of the files or directories
func (m *Modularizer) CreateModule() error {
	if err := m.createDirectoryStructure(); err != nil {
		return err
	}
	if err := m.createModuleFile(); err != nil {
		return err
	}
	if err := m.createEntityFiles(); err != nil {
		return err
	}
	if err := m.createRepositoryFiles(); err != nil {
		return err
	}
	if err := m.createAppFile(); err != nil {
		return err
	}
	if err := m.createCommandFiles(); err != nil {
		return err
	}
	if err := m.createCommandTypesFile(); err != nil {
		return err
	}
	if m.templateData.DefaultQueries {
		if err := m.createDefaultQueryFiles(); err != nil {
			return err
		}
	}
	if err := m.createQueryTypesFile(); err != nil {
		return err
	}
	if err := m.createEventFiles(); err != nil {
		return err
	}
	if err := m.createPostgresQueriesFile(); err != nil {
		return err
	}
	if err := m.createPostgresSchemaFile(); err != nil {
		return err
	}
	if err := m.createRestServerConfFile(); err != nil {
		return err
	}
	if err := m.createRestServerGenerateFile(); err != nil {
		return err
	}
	if err := m.createRestClientConfFile(); err != nil {
		return err
	}
	if err := m.createRestClientGenerateFile(); err != nil {
		return err
	}
	if err := m.createHttpAdapterFile(); err != nil {
		return err
	}
	if err := m.updateFrontendOpenAPIConfig(); err != nil {
		return err
	}
	if err := m.createBackendOpenApiSchemaFile(); err != nil {
		return err
	}
	if err := m.createSqlSubscriberAdapterFile(); err != nil {
		return err
	}
	if err := m.updateSqlCConfig(); err != nil {
		return err
	}
	if err := m.createPostgresAdapterFiles(); err != nil {
		return err
	}
	err := m.checkIfMigrationInitFilesExist() // returns an error if there are any existing module initialization migration files in the db/migrations directory
	if err == nil {
		if err := m.createDbMigrationsFiles(); err != nil {
			return err
		}
	}
	if err := m.createFrontendFiles(); err != nil {
		return err
	}
	if err := m.generateOapiCodegen(); err != nil {
		return err
	}
	if err := m.runFormattingAndLinting(); err != nil {
		return err
	}

	return nil
}

// createDirectoryStructure creates the directory structure for the new module
func (m *Modularizer) createDirectoryStructure() error {
	slog.Info("creating directory structure for the new module")

	workDir, err := os.Getwd()
	if err != nil {
		return err
	}
	basePath := workDir + "/internal/" + m.templateData.Name.Flat()
	subDirs := []string{
		"adapters/inbound",
		"adapters/outbound",
		"app",
		"app/command",
		"app/query",
		"app/event",
		"domain/" + m.templateData.Name.Flat(),
	}

	for _, subDir := range subDirs {
		dirPath := basePath + "/" + subDir
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	dbPath := workDir + "/db/" + m.templateData.Name.Flat()
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return err
	}

	apiRestPaths := []string{
		workDir + "/api/rest/" + m.templateData.Name.Flat() + "/server",
		workDir + "/api/rest/" + m.templateData.Name.Flat() + "/client",
	}
	for _, apiRestPath := range apiRestPaths {
		err := os.MkdirAll(apiRestPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	apiFrontendPaths := []string{
		workDir + "/api/frontend/test/mocks/" + m.templateData.Name.Flat(),
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat(),
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/components",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/components/__tests__",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/assets",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/composables",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/config",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/router",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/stores",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/views",
		workDir + "/api/frontend/src/" + m.templateData.Name.Flat() + "/views/__tests__",
	}
	for _, apiFrontendPath := range apiFrontendPaths {
		err := os.MkdirAll(apiFrontendPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// createModuleFile creates the /{module}/module.go file
func (m *Modularizer) createModuleFile() error {
	slog.Info("Backend: creating module file for the new module")

	err := m.getTemplateAndWriteToFile(
		"/templates/internal/module.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/module.go",
	)
	if err != nil {
		return err
	}
	return nil
}

// createRepositoryFiles creates the /{module}/domain/{module}/repository.go file
func (m *Modularizer) createRepositoryFiles() error {
	slog.Info("Backend: creating repository file for the new module")
	err := m.getTemplateAndWriteToFile(
		"/templates/internal/domain/repository.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/domain/"+m.templateData.Name.Flat()+"/repository.go",
	)
	if err != nil {
		return err
	}
	return nil
}

// createAppFile creates the /{module}/app/app.go file
func (m *Modularizer) createAppFile() error {
	slog.Info("Backend: creating application file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/app/app.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/app/app.go",
	); err != nil {
		return err
	}
	return nil
}

// createPostgresQueriesFile creates the /db/postgres_queries.sql file
func (m *Modularizer) createPostgresQueriesFile() error {
	slog.Info("Backend: creating Postgres queries file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/db/postgres_queries.sql.templ",
		"db/"+m.templateData.Name.Flat()+"/postgres_queries.sql",
	); err != nil {
		return err
	}
	return nil
}

// createPostgresSchemaFile creates the /db/postgres_schema.sql file
func (m *Modularizer) createPostgresSchemaFile() error {
	slog.Info("Backend: creating Postgres schema file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/db/postgres_schema.sql.templ",
		"db/"+m.templateData.Name.Flat()+"/postgres_schema.sql",
	); err != nil {
		return err
	}
	return nil
}

// createHttpAdapterFile creates the /{module}/adapters/inbound/http.go file
func (m *Modularizer) createHttpAdapterFile() error {
	slog.Info("Backend: creating HTTP adapter file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/inbound/http.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/inbound/http.go",
	); err != nil {
		return err
	}
	return nil
}

// createRestServerConfFile creates the /{module}/api/rest/server/cfg.yaml file
func (m *Modularizer) createRestServerConfFile() error {
	slog.Info("Backend: creating REST server configuration file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/api/rest/server/cfg.yaml.templ",
		"api/rest/"+m.templateData.Name.Flat()+"/server/cfg.yaml",
	); err != nil {
		return err
	}
	return nil
}

// createRestServerGenerateFile creates the /{module}/api/rest/server/generate.go file
func (m *Modularizer) createRestServerGenerateFile() error {
	slog.Info("Backend: creating REST server generator file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/api/rest/server/generate.go.templ",
		"api/rest/"+m.templateData.Name.Flat()+"/server/generate.go",
	); err != nil {
		return err
	}
	return nil
}

// createRestClientConfFile creates the /{module}/api/rest/client/cfg.yaml file
func (m *Modularizer) createRestClientConfFile() error {
	slog.Info("Backend: creating REST client configuration file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/api/rest/client/cfg.yaml.templ",
		"api/rest/"+m.templateData.Name.Flat()+"/client/cfg.yaml",
	); err != nil {
		return err
	}
	return nil
}

// createRestClientGenerateFile creates the /{module}/api/rest/client/generate.go file
func (m *Modularizer) createRestClientGenerateFile() error {
	slog.Info("Backend: creating REST client generator file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/api/rest/client/generate.go.templ",
		"api/rest/"+m.templateData.Name.Flat()+"/client/generate.go",
	); err != nil {
		return err
	}
	return nil
}

// createBackendOpenApiSchemaFile creates the OpenAPI schema file for the new module in the api/rest/{module}/api.yaml file
func (m *Modularizer) createBackendOpenApiSchemaFile() error {
	slog.Info("Backend: creating OpenAPI file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/api/rest/api.yaml.templ",
		"api/rest/"+m.templateData.Name.Flat()+"/api.yaml",
	); err != nil {
		return err
	}
	return nil
}

// generateOapiCodegen runs the OpenAPI Generator command to generate the REST API server and client code for the new module based on the OpenAPI schema file created in the previous step, and returns an error if there was an issue running the command
func (m *Modularizer) generateOapiCodegen() error {
	slog.Info("Backend: generating OpenAPI codegen for the new module")
	cmd := exec.Command("make", "gen-api-be")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type ErrModuleInitMigrationFileExists struct{}

func (e ErrModuleInitMigrationFileExists) Error() string {
	return "a module initialization migration file already exists in the db/migrations directory. Please remove or rename any existing migration files that start with 'init_' before creating a new module."
}

func (m *Modularizer) checkIfMigrationInitFilesExist() error {
	slog.Info("Backend: checking for existing module initialization migration files in the db/migrations directory")
	migrationsDirectory := "db/migrations/"
	// walk the directory
	err := filepath.WalkDir(migrationsDirectory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".sql") && strings.Contains(d.Name(), "init_") && strings.Contains(d.Name(), m.templateData.Name.Snake()+"_module") {
			return ErrModuleInitMigrationFileExists{}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// createDbMigrationsFiles runs the command to create the database migration files for the new module using the name of the module as part of the migration name, and returns an error if there was an issue running the command
// func (m *Modularizer) createDbMigrationsFiles() error {
// 	slog.Info("Backend: creating database migrations files for the new module")
// 	cmd := exec.Command("make", "migrate-create", "MIGNAME=init_"+m.templateData.Name.Flat())
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }

func (m *Modularizer) createDbMigrationsFiles() error {
	timeStamp := time.Now().UTC().Format("20060102150405")
	if err := m.getTemplateAndWriteToFile(
		"/templates/db/migrations/migration_init_up.sql.templ",
		"db/migrations/"+timeStamp+"_init_"+m.templateData.Name.Snake()+"_module.up.sql",
	); err != nil {
		return err
	}

	if err := m.getTemplateAndWriteToFile(
		"/templates/db/migrations/migration_init_down.sql.templ",
		"db/migrations/"+timeStamp+"_init_"+m.templateData.Name.Snake()+"_module.down.sql",
	); err != nil {
		return err
	}

	return nil
}

// runFormattingAndLinting runs the command to format and lint the code for the new module, and logs a warning if it fails
func (m *Modularizer) runFormattingAndLinting() error {
	slog.Info("Backend: running formatting and linting for the new module")
	cmd := exec.Command("make", "lintfmt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		slog.Warn("formatting/linting failed (non-fatal)", "error", err)
	}
	return nil
}

// createCommandFiles creates the command files for the new module in the /{module}/app/command directory,
// and also creates the corresponding entity method files in the /{module}/domain/{module} directory,
// and returns an error if there was an issue creating any of the files
func (m *Modularizer) createCommandFiles() error {
	slog.Info("Backend: creating command files for the new module")

	for _, command := range m.templateData.Commands {
		if err := m.getTemplateAndWriteToFileWithTemplateData(
			"/templates/internal/app/command/command.go.templ",
			"internal/"+m.templateData.Name.Flat()+"/app/command/"+command.Name.Snake()+".go",
			struct {
				Mod module.Module
				Cmd module.Command
			}{
				Mod: m.templateData,
				Cmd: command,
			},
		); err != nil {
			return err
		}
	}

	// for _, command := range m.templateData.Commands {
	// 	if err := m.getTemplateAndWriteToFileWithTemplateData(
	// 		"/templates/internal/domain/entity_method.go.templ",
	// 		"internal/"+m.templateData.Name.Flat()+"/domain/"+m.templateData.Name.Flat()+"/"+command.Name.Snake()+".go",
	// 		struct {
	// 			Mod module.Module
	// 			Cmd module.Command
	// 		}{
	// 			Mod: m.templateData,
	// 			Cmd: command,
	// 		},
	// 	); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

// createEventFiles creates the event files for the new module in the /{module}/app/event directory,
// and returns an error if there was an issue creating any of the files
func (m *Modularizer) createEventFiles() error {
	slog.Info("Backend: creating event files for the new module")

	for _, event := range m.templateData.Events {
		if err := m.getTemplateAndWriteToFileWithTemplateData(
			"/templates/internal/app/event/event.go.templ",
			"internal/"+m.templateData.Name.Flat()+"/app/event/"+event.Name.Snake()+".go",
			struct {
				Mod module.Module
				Evt module.Event
			}{
				Mod: m.templateData,
				Evt: event,
			},
		); err != nil {
			return err
		}
	}
	return nil
}

// createDefaultQueryFiles creates the default query files for the new module in the /{module}/app/query directory
// based on the aggregate root entity of the module, and returns an error if there was an issue creating any of the files
func (m *Modularizer) createDefaultQueryFiles() error {
	slog.Info("Backend: creating default query files for the new module")

	for _, aggregate := range m.templateData.Aggregates {
		if err := m.getTemplateAndWriteToFileWithTemplateData(
			"/templates/internal/app/query/find_all_query.go.templ",
			"internal/"+m.templateData.Name.Flat()+"/app/query/find_all_"+aggregate.Name.PluralSnake()+".go",
			struct {
				Mod module.Module
				Agg module.Entity
			}{
				Mod: m.templateData,
				Agg: aggregate,
			},
		); err != nil {
			return err
		}

		if err := m.getTemplateAndWriteToFileWithTemplateData(
			"/templates/internal/app/query/find_one_query.go.templ",
			"internal/"+m.templateData.Name.Flat()+"/app/query/find_one_"+aggregate.Name.Snake()+".go",
			struct {
				Mod module.Module
				Agg module.Entity
			}{
				Mod: m.templateData,
				Agg: aggregate,
			},
		); err != nil {
			return err
		}

	}

	return nil
}

// createQueryTypesFile creates the /{module}/app/query/types.go file
func (m *Modularizer) createQueryTypesFile() error {
	slog.Info("Backend: creating query types file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/app/query/types.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/app/query/types.go",
	); err != nil {
		return err
	}
	return nil
}

// createSqlSubscriberAdapterFile creates the /{module}/adapters/inbound/sql_subscriber.go file
func (m *Modularizer) createSqlSubscriberAdapterFile() error {
	slog.Info("Backend: creating SQL subscriber adapter file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/inbound/sql_subscriber.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/inbound/sql_subscriber.go",
	); err != nil {
		return err
	}
	return nil
}

// createFrontendFiles creates the frontend files for the new module, including components, views, composables, router, and config
func (m *Modularizer) createFrontendFiles() error {
	slog.Info("Frontend: creating files for the new module")

	// components
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/components/DashboardCard.vue.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/components/DashboardCard.vue",
	); err != nil {
		return err
	}
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/components/__tests__/DashboardCard.test.ts.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/components/__tests__/DashboardCard.test.ts",
	); err != nil {
		return err
	}

	// config
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/config/index.ts.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/config/index.ts",
	); err != nil {
		return err
	}

	// router
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/router/index.ts.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/router/index.ts",
	); err != nil {
		return err
	}

	// composables
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/composables/useModule.ts.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/composables/use"+m.templateData.Name.Pascal()+".ts",
	); err != nil {
		return err
	}

	// views
	if err := m.getTemplateAndWriteToFile(
		"/templates/api/frontend/src/views/ModuleView.vue.templ",
		"api/frontend/src/"+m.templateData.Name.Flat()+"/views/"+m.templateData.Name.Pascal()+"View.vue",
	); err != nil {
		return err
	}
	return nil
}

// createEntityFiles creates the entity file for the aggregate root entity of the new module
func (m *Modularizer) createEntityFiles() error {
	slog.Info("Backend: creating entity file for the new module")

	for _, aggregate := range m.templateData.Aggregates {
		if err := m.getTemplateAndWriteToFileWithTemplateData(
			"/templates/internal/domain/entity.go.templ",
			"internal/"+m.templateData.Name.Flat()+"/domain/"+m.templateData.Name.Flat()+"/"+aggregate.Name.Snake()+".go",
			struct {
				Mod module.Module
				Agg module.Entity
			}{
				Mod: m.templateData,
				Agg: aggregate,
			},
		); err != nil {
			return err
		}
	}
	return nil
}

// createCommandTypesFile creates the /{module}/app/command/types.go file
func (m *Modularizer) createCommandTypesFile() error {
	slog.Info("Backend: creating command types file for the new module")

	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/app/command/types.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/app/command/types.go",
	); err != nil {
		return err
	}
	return nil
}

// getTemplateAndWriteToFile is a helper function that takes in a template path and an output path, parses the template, executes it with the modularizer's template data, and writes the output to the specified file path
func (m *Modularizer) getTemplateAndWriteToFile(templateRelativePath string, outputPath string) error {
	tmpl := template.Must(template.ParseFiles(utils.GetDirectoryOfCurrentFile() + templateRelativePath))
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := tmpl.Execute(f, m.templateData); err != nil {
		return err
	}
	return nil
}

// getTemplateAndWriteToFileWithTemplateData is a helper function that takes in a template path, an output path, and custom template data, parses the template, executes it with the provided template data, and writes the output to the specified file path
func (m *Modularizer) getTemplateAndWriteToFileWithTemplateData(templateRelativePath string, outputPath string, templateData any) error {
	tmpl := template.Must(template.ParseFiles(utils.GetDirectoryOfCurrentFile() + templateRelativePath))
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := tmpl.Execute(f, templateData); err != nil {
		return err
	}
	return nil
}

// createPostgresAdapterFiles creates the Postgres adapter files for the new module, including the query interface wrapper, mapping, and main adapter file
func (m *Modularizer) createPostgresAdapterFiles() error {
	slog.Info("Backend: creating postgres adapter files for the new module")

	// create the query interface wrapper file
	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/outbound/postgres_query_interface_wrapper.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/outbound/postgres_query_interface_wrapper.go",
	); err != nil {
		return err
	}

	// create the query interface wrapper test file
	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/outbound/postgres_query_interface_wrapper_test.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/outbound/postgres_query_interface_wrapper_test.go",
	); err != nil {
		return err
	}

	// generate the sqlc files
	if err := m.generateSQLC(); err != nil {
		return err
	}

	// create the main adapter file
	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/outbound/postgres.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/outbound/postgres.go",
	); err != nil {
		return err
	}

	// create the mapping file
	if err := m.getTemplateAndWriteToFile(
		"/templates/internal/adapters/outbound/postgres_mapping.go.templ",
		"internal/"+m.templateData.Name.Flat()+"/adapters/outbound/postgres_mapping.go",
	); err != nil {
		return err
	}
	return nil
}

// generateSQLC runs the command to generate the SQLC files for the new module using the sqlc configuration specified in the .sqlc.yaml file
func (m *Modularizer) generateSQLC() error {
	slog.Info("Backend: generating SqlC queries for the new module")
	cmd := exec.Command("make", "queries")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type SqlCSqlConfig struct {
	Name    string `json:"name" yaml:"name"`
	Schema  string `json:"schema" yaml:"schema"`
	Queries string `json:"queries" yaml:"queries"`
	Engine  string `json:"engine" yaml:"engine"`
}

type SQLCConfig struct {
	Sql []SqlCSqlConfig `json:"sql" yaml:"sql"`
}

// checkIfSQLCConfigExists checks if the SQLC configuration for the new module already exists in the .sqlc.yaml file, and returns true if it does, false if it doesn't, and an error if there was an issue reading or parsing the file
func (m *Modularizer) updateSqlCConfig() error {
	slog.Info("Backend: updating SqlC configuration for the new module")

	exists, err := m.checkIfSQLCConfigExists()
	if err != nil {
		return err
	}
	if !exists {
		if err := m.appendToSQLCConfig(); err != nil {
			return err
		}
	}
	return nil
}

// checkIfSQLCConfigExists checks if the SQLC configuration for the new module already exists in the .sqlc.yaml file, and returns true if it does, false if it doesn't, and an error if there was an issue reading or parsing the file
func (m *Modularizer) checkIfSQLCConfigExists() (bool, error) {
	slog.Info("checking if SqlC configuration for this module already exists")
	//   parse the sqlC string to check if it's valid yaml
	f, err := os.ReadFile(".sqlc.yaml")
	if err != nil {
		return false, err
	}
	sqlC := string(f)

	var config SQLCConfig
	err = yaml.Unmarshal([]byte(sqlC), &config)
	if err != nil {
		return false, err
	}

	// check if the config has the module's sqlC config already, if yes, return an error
	for _, sqlConfig := range config.Sql {
		if sqlConfig.Name == m.templateData.Name.Flat()+"_postgres" {
			return true, nil
		}
	}

	return false, nil
}

// appendToSQLCConfig appends the SQLC configuration for the new module to the .sqlc.yaml file using the sqlc configuration template and the modularizer's template data, and returns an error if there was an issue writing to the file
func (m *Modularizer) appendToSQLCConfig() error {
	slog.Info("Backend: appending SqlC configuration for the new module to .sqlc.yaml")

	f, err := os.OpenFile(".sqlc.yaml", os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	sqlcConfTmpl := template.Must(template.New("sqlcConfig").Parse(templates.SqlcConfig))
	var sqlcConfBuilder strings.Builder
	err = sqlcConfTmpl.Execute(&sqlcConfBuilder, m.templateData)
	if err != nil {
		return err
	}
	_, err = f.WriteString("\n" + sqlcConfBuilder.String())
	if err != nil {
		return err
	}

	return nil
}

// addNewGeneratorConfig adds the OpenAPI Generator configuration for the new module to the .openapitools.json file, and returns the updated configuration as a map[string]any, or an error if there was an issue updating the configuration
func (m *Modularizer) addNewGeneratorConfig(config map[string]any) map[string]any {
	slog.Info("Frontend: appending OpenAPI Generator configuration for the new module to .openapitools.json")

	generators := config["generator-cli"].(map[string]any)["generators"].(map[string]any)
	moduleName := m.templateData.Name.Flat()
	generators[moduleName] = map[string]any{
		"generatorName": "typescript-fetch",
		"output":        "#{cwd}/src/" + moduleName + "/services",
		"inputSpec":     "../rest/" + moduleName + "/api.yaml",
		"additionalProperties": map[string]any{
			"withInterfaces": true,
		},
	}
	return config
}

// updateFrontendOpenAPIConfig reads the .openapitools.json file, updates it with the new module's OpenAPI Generator configuration using the addNewGeneratorConfig helper function, and writes the updated configuration back to the .openapitools.json file, returning an error if there was an issue reading, updating, or writing the file
func (m *Modularizer) updateFrontendOpenAPIConfig() error {
	slog.Info("Frontend: updating OpenAPI Generator configuration for the new module")

	f, err := os.ReadFile("api/frontend/.openapitools.json")
	if err != nil {
		return err
	}

	var config map[string]any
	if err = json.Unmarshal(f, &config); err != nil {
		return err
	}

	updatedConfig := m.addNewGeneratorConfig(config)

	marshalledConfig, err := json.MarshalIndent(updatedConfig, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("api/frontend/.openapitools.json", marshalledConfig, 0o644)
	if err != nil {
		return err
	}
	return nil
}
