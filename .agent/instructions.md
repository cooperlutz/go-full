# AI Agent Repository Instructions

Welcome to the Go-Full repository. You are operating inside a configuration-driven, highly optimized "AI PoC Factory". This repository is designed to minimize boilerplate and enforce type safety through specific code-generation patterns.

When developing new backend features or an AI Proof-of-Concept, you **must** adhere strictly to the following workflows:

## 1. Database & Schema changes
Do NOT write raw Go code for database interactions. All data access is managed via `sqlc`.
1. **Schema modifications:** If you need new tables, create a new SQL migration in `db/migrations/` (use `make migrate-create MIGNAME=<name>` if available, or manually create).
2. **Writing Queries:** Define all your SQL queries with `sqlc` annotations in `db/query/`.
3. **Generate DB Code:** Whenever you change migration schemas or query definitions, you MUST run:
   ```bash
   make queries
   ```
   This will auto-generate the type-safe Go data models inside `internal/store/db/`.

## 2. API & Endpoint modifications
Do NOT manually wire Handlers or Models in Go for HTTP endpoints. The HTTP contract is strictly driven by OpenAPI.
1. **Swagger Update:** Modify `api/openapi.yaml` (or the respective Swagger definition) to define your new routes, request bodies, and response schemas.
2. **Generate API Code:** After modifying the OpenAPI spec, you MUST run:
   ```bash
   make gen-api
   ```
   This auto-generates the Chi HTTP router interfaces and data models in `api/`.
3. **Implement the Controller:** After generation, create the actual endpoint logic in `internal/controllers/` or `app/api/` logic layer by implementing the auto-generated ServerInterface. 

## 3. Using the AI Subsystem
If you need to incorporate Large Language Models (LLMs):
1. Import `github.com/cooperlutz/go-full/pkg/ai`.
2. Instantiate the centralized `ai.Client`.
3. The client is already wired with OpenTelemetry. Do not wrap AI calls in your own spans unless absolutely necessary; let `pkg/ai` handle usage tracking.

## General Guidelines
- **Run the Suite:** Always run `make lintfmt` and `make test-be` before considering a backend task complete.
- **Frontend changes:** The Vue.js frontend relies on auto-generated clients from the OpenAPI spec. When `make gen-api` runs, it updates both the backend and frontend configurations. UI logic should go into `api/frontend/src/app/`.
