# Development Tools

## Azure Developer CLI (azd)

Documentation: [Azure Developer CLI (azd)](https://learn.microsoft.com/en-us/azure/developer/azure-developer-cli/)

---

## Commitizen

| category | detail |
| --- | --- |
| Documentation | [Commitizen](https://commitizen-tools.github.io/commitizen/) |
| System Context | Project |
| Purpose | Standardizing commit messages according to Conventional Commits |
| Configuration | [.cz.yaml](../../../.cz.yaml) |
| Usage | Run `make commit` to create a commit with a standardized message |

---

## Docker

| category | detail |
| --- | --- |
| Documentation | [docker](https://docs.docker.com/) |
| System Context | Project |
| Purpose | Containerization of the application |
| Configuration | [build/docker/Dockerfile](../../../build/docker/Dockerfile) |
| Usage |  |

---

## Docker Compose

| category | detail |
| --- | --- |
| Documentation | [docker compose](https://docs.docker.com/compose/) |
| System Context | Project |
| Purpose | Development environment orchestration |
| Configuration | [deploy/compose/](../../../deploy/compose/) |
| Usage | `make compose` |

---

## ESLint

| category | detail |
| --- | --- |
| Documentation | [ESLint](https://eslint.org/docs/latest/) |
| System Context | Frontend |
| Purpose | Linting for JavaScript/TypeScript code in the frontend application |
| Configuration | [.eslintrc.json](../../../api/frontend/.eslintrc.json) |
| Usage | `make lint-fe` |

---

## GoFumpt

| category | detail |
| --- | --- |
| Documentation | [GoFumpt](https://github.com/mvdan/gofumpt) |
| System Context | Core / Backend |
| Purpose | Go code formatter, stricter than `gofmt` |
| Configuration | N/A |
| Usage | `make format`, `make format-be` |

---

## Golangci-lint

| category | detail |
| --- | --- |
| Documentation | [Golangci-lint](https://golangci-lint.run/docs/) |
| System Context | Core / Backend |
| Purpose | Go linting |
| Configuration | [.golangci.yml](../../../.golangci.yml) |
| Usage | `make lint`, `make lint-be` |

---

## Golang-Migrate

| category | detail |
| --- | --- |
| Documentation | [golang-migrate](https://github.com/golang-migrate/migrate) |
| System Context | Core / Backend |
| Purpose | Database schema migrations |
| Configuration | [migrations/](../../../db/migrations) |
| Usage |  |

---

## GoReleaser

| category | detail |
| --- | --- |
| Documentation | [GoReleaser](https://goreleaser.com/) |
| System Context | Core / Backend |
| Purpose | Automation and orchestration of project releases |
| Configuration | [.goreleaser.yml](../../../.goreleaser.yml) |
| Usage |  |

---

## Knip

| category | detail |
| --- | --- |
| Documentation | [Knip](https://knip.dev/) |
| System Context | Frontend |
| Purpose | Identifying unused dependencies in the frontend application |
| Configuration | [knip.config.js](../../../api/frontend/knip.json) |
| Usage | `make deps`, `pnpm declutter` |

---

## Make

| category | detail |
| --- | --- |
| Documentation | [Make (Makefile)](https://www.gnu.org/software/make/manual/make.html) |
| System Context | Project |
| Purpose | Automation & Orchestration of common tasks and commands for project development |
| Configuration | [Makefile](../../../Makefile) |
| Usage | `make help` - displays all available make commands |

---

## Mockery

| category | detail |
| --- | --- |
| Documentation | [mockery](https://vektra.github.io/mockery/latest/) |
| System Context | Core / Backend |
| Purpose | Mock generation for Go interfaces |
| Configuration | [mockery.yml](../../../.mockery.yml) |
| Usage | Run `make mock` to generate mocks for the Go interfaces |

---

## MkDocs + Material for MkDocs

| category | detail |
| --- | --- |
| Documentation | [mkdocs](https://www.mkdocs.org/), [Material for Mkdocs](https://squidfunk.github.io/mkdocs-material/) |
| System Context | Documentation |
| Purpose | Documentation site generation and theming |
| Configuration | [mkdocs.yml](../../../mkdocs.yml) |
| Usage | Run `make docs` to build the documentation site |

---

## OapiCodeGen

| category | detail |
| --- | --- |
| Documentation | [OapiCodeGen](https://github.com/oapi-codegen/oapi-codegen/tree/main) |
| System Context | Core / Backend |
| Purpose | Generating Go server and client code from OpenAPI 3.0 specifications |
| Configuration | [cfg.yaml](../../../api/rest/pingpong/v1/server/cfg.yaml) |
| Usage | `make gen-api-be` |

---

## OpenAPI-Generator-Cli

| category | detail |
| --- | --- |
| Documentation | [OpenAPI-Generator-Cli](https://github.com/OpenAPITools/openapi-generator-cli) |
| System Context | Frontend |
| Purpose | Generation of the typescript API client library to enable the consumption of REST APIs from OpenAPI specifications |
| Configuration | [openapi-generator-config.json](../../../api/frontend/src/api/openapi-generator-config.json) |
| Usage | `make gen-api-fe` |

---

## Prettier

| category | detail |
| --- | --- |
| Documentation | [Prettier](https://prettier.io/docs/en/index.html) |
| System Context | Frontend |
| Purpose | Code formatter for consistent code style in the frontend application |
| Configuration | [prettier.config.js](../../../api/frontend/prettierrc.json) |
| Usage | `make format-fe` |

---

## PNPM

| category | detail |
| --- | --- |
| Documentation | [pnpm](https://pnpm.io/motivation) |
| System Context | Frontend |
| Purpose | Package management for the frontend application |
| Configuration | [package.json](../../../api/frontend/package.json) |
| Usage |  |

---

## SQLC

| category | detail |
| --- | --- |
| Documentation | [sqlc](https://docs.sqlc.dev/en/latest/) |
| System Context | Core / Backend |
| Purpose | Generation of go code for database interaction from sql queries |
| Configuration | [.sqlc.yaml](../../../.sqlc.yaml) |
| Usage |   |

---

## UV

| category | detail |
| --- | --- |
| Documentation | [uv](https://docs.astral.sh/uv/) |
| System Context | Documentation |
| Purpose | Python environment and package management for the documentation site |
| Configuration | [pyproject.toml](../../pyproject.toml) |
| Usage |  |

---

## Vite

| category | detail |
| --- | --- |
| Documentation | [vite](https://vite.dev/guide/) |
| System Context | Frontend |
| Purpose | build tool for frontend |
| Configuration | [vite.config.ts](../../../api/frontend/vite.config.ts) |
| Usage |  |

---

## Vitest

| category | detail |
| --- | --- |
| Documentation | [Vitest](https://vitest.dev/) |
| System Context | Frontend |
| Purpose | Testing framework for the frontend application |
| Configuration | [vite.config.ts](../../../api/frontend/vite.config.ts) |
| Usage | `make test-fe` |
