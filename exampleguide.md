# AI PoC Factory - Example Guide

This guide walks through the full AI PoC Factory pipeline: from a written business case to a running, API-exposed application module.

## What does this produce?

The `generate-poc` command takes a markdown file describing a product idea and produces a **fully scaffolded, deployable module** that includes:

- **REST API endpoints** accessible by any HTTP client or web frontend (OpenAPI 3.0 spec + auto-generated Go server code)
- **PostgreSQL database** with schema, type-safe queries, and migrations
- **Domain logic** with entity structs, repository interfaces, command/query/event handlers
- **Vue.js frontend** scaffolding with components, router, composables, and a TypeScript API client
- **Event-driven architecture** with pub/sub event definitions
- **Observability** with OpenTelemetry tracing spans on every layer
- **Auto-wiring** into the main application (imports, initialization, route mounting, config updates)

The result is a module that serves HTTP endpoints a web frontend (or any client) can call, backed by a real database, with full event sourcing and observability. The Go binary embeds the Vue.js SPA and serves it at the root path, so you get a single deployable artifact that includes both the API and the frontend.

## Prerequisites

- Go 1.24+
- Make
- Docker & Docker Compose (for running the full stack)
- A valid `.env` file with Azure OpenAI or OpenAI credentials

### Required `.env` variables for PoC generation

```
OPENAI_API_KEY=your-api-key
AZURE_OPENAI_ENDPOINT=https://your-resource.openai.azure.com
AZURE_OPENAI_DEPLOYMENT_NAME=gpt-4o
AZURE_OPENAI_API_VERSION=2024-10-21
```

## Example 1: Standard module (no AI at runtime)

### Step 1: Write a business case

Create a markdown file describing your idea:

```markdown
# Expense Tracker
We need an API to track business expenses. Users can submit an expense
with a description, amount, and category (e.g., "travel", "food",
"office supplies"). They should be able to retrieve all submitted
expenses. We need a database table to store these expense records.
```

Save this as `docs/expense_tracker_case.md`.

### Step 2: Generate the module

```bash
make generate-poc prompt="docs/expense_tracker_case.md"
```

This single command will:
1. Send the business case to the AI, which generates a module config (aggregates, commands, events, field types)
2. Scaffold the full module using the modularizer (30+ files across domain, app, API, infra, frontend layers)
3. Generate database queries with sqlc and REST server code with oapi-codegen
4. Wire the module into `app/app.go` (import, init, route mount)
5. Update `.sqlc.yaml`, `.mockery.yml`, and `.openapitools.json`

### Step 3: Review the generated code

```bash
# See the AI-generated module config
cat tools/pocgen/generated/expensetracker.yaml

# Browse the generated module
ls internal/expensetracker/

# Check the API spec
cat api/rest/expensetracker/api.yaml

# Check the database schema
cat db/expensetracker/postgres_schema.sql
```

### Step 4: Customize entity fields

The generated domain entity and SQL schema have fields commented out as scaffolding stubs. Uncomment and adjust them to match your needs, then regenerate:

```bash
make gen    # runs sqlc + oapi-codegen + mockery
make test   # verify everything compiles and passes
```

### Step 5: Run the application

```bash
make run
```

This starts the full Docker Compose stack (app, PostgreSQL, Nginx, Prometheus, Jaeger). The application will be available at `http://app.lvh.me`.

### Step 6: Call the API

The generated module exposes endpoints under `/api/expensetracker/`:

```bash
# Submit an expense
curl -X POST http://localhost:8080/api/expensetracker/v1/expenses \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{"description": "Team lunch", "amount": 45.50, "category": "food"}'

# Get all expenses
curl http://localhost:8080/api/expensetracker/v1/expenses \
  -H "Authorization: Bearer <your-jwt-token>"
```

> **Note:** The `/api` routes are behind auth middleware. Authenticate first via `POST /auth/login` to get a JWT token.

## Example 2: AI-integrated module (LLM at runtime)

If your business case involves calling an AI/LLM at runtime, the tool detects this automatically.

### Write a business case that needs AI

```markdown
# AI Startup Idea Rater
We need a single API endpoint that accepts a "Startup Pitch".
The API will pass the pitch to an LLM, returning a JSON response
with a "score" (1-10) and a "critique". We also need a database
table to log these scores over time.
```

### Generate it

```bash
make generate-poc prompt="docs/startup_rater_case.md"
```

The AI will set `ai_integration: true` in the generated config. The tool will:
- Scaffold the standard module (same as above)
- Additionally inject `pkg/ai.Client` as a dependency into the module
- Pass the AI client through to the application layer

You will then need to add the AI call logic in your command handler (e.g., call `aiClient.GenerateCompletion()` in the submit command).

## What gets generated (full file list)

For a module named `expense_tracker`, the tool creates:

```
internal/expensetracker/
  module.go                              # Module entry point & dependency wiring
  adapters/inbound/
    http.go                              # HTTP adapter (REST controller)
    http_openapi.gen.go                  # Generated OpenAPI server code
    sql_subscriber.go                    # Event subscriber adapter
  adapters/outbound/
    postgres.go                          # PostgreSQL repository adapter
    postgres_mapping.go                  # DB <-> domain mapper
    postgres_query_interface_wrapper.go  # sqlc interface wrapper
    postgres_*.gen.go                    # Generated sqlc code
  app/
    app.go                               # Application service (commands, queries, events)
    command/submit_expense.go            # Command handler
    command/types.go                     # Command type definitions
    query/find_all_expenses.go           # Query: list all
    query/find_one_expense.go            # Query: get by ID
    query/types.go                       # Query type definitions
    event/expense_submitted.go           # Event handler
  domain/expensetracker/
    expense.go                           # Domain entity
    repository.go                        # Repository interface

db/expensetracker/
  postgres_schema.sql                    # Database schema
  postgres_queries.sql                   # sqlc-annotated queries

db/migrations/
  {timestamp}_init_expense_tracker_module.up.sql
  {timestamp}_init_expense_tracker_module.down.sql

api/rest/expensetracker/
  api.yaml                               # OpenAPI 3.0 spec
  server/cfg.yaml, generate.go           # oapi-codegen server config
  client/cfg.yaml, generate.go           # oapi-codegen client config

api/frontend/src/expensetracker/
  components/DashboardCard.vue           # Vue component
  composables/useExpensetracker.ts       # Vue composable
  config/index.ts                        # Module config
  router/index.ts                        # Frontend routes
  views/ExpensetrackerView.vue           # Page view

tools/pocgen/generated/
  expensetracker.yaml                    # AI-generated config (for auditing)
```

Config files updated:
- `.sqlc.yaml` -- new sqlc entry
- `.mockery.yml` -- new mock generation entries
- `.openapitools.json` -- new frontend API client entry
- `app/app.go` -- module import, init, route mount

## Running tests

```bash
# Test a specific module
go test -v ./internal/expensetracker/...

# Test everything
go test ./internal/... ./pkg/...

# Full CI pipeline
make test
```

## Architecture overview

Each generated module follows a hexagonal (ports & adapters) architecture:

```
                    HTTP Request
                        |
                   [Inbound Adapter]
                   (http.go / OpenAPI)
                        |
                  [Application Layer]
                  (commands / queries / events)
                        |
                   [Domain Layer]
                   (entities / repository interface)
                        |
                  [Outbound Adapter]
                  (postgres.go / sqlc queries)
                        |
                   PostgreSQL
```

All API routes are mounted under `/api/{module_name}/` behind JWT auth middleware. The Vue.js SPA is embedded in the Go binary and served at the root path, so the final deployable is a single binary that includes both the frontend and backend.

## Troubleshooting

**"Resource not found" (404) from Azure OpenAI**
- Verify `AZURE_OPENAI_ENDPOINT` is the base URL only (no path or query params)
- Verify `AZURE_OPENAI_DEPLOYMENT_NAME` matches an actual deployment in your Azure portal
- Try a different `AZURE_OPENAI_API_VERSION` (e.g., `2024-10-21`)

**"sqlc: command not found"**
- The Makefile uses `go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest` as a fallback. Ensure you have Go installed and network access.

**Frontend generation fails**
- This is expected if `openapi-generator-cli` is not installed. The backend code is fully generated regardless. Install frontend dependencies with `pnpm install` in `api/frontend/` if you need the TypeScript client.

**Lint errors about `dist/*`**
- This is a pre-existing issue: the Vue.js frontend hasn't been built yet. Run `make build-fe` to build it, or ignore for backend-only development.
