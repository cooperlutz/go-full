# Development Tools

## Make

#### Documentation

[Make (Makefile)](https://www.gnu.org/software/make/manual/make.html)

#### System Context

All

#### Purpose

Automation & Orchestration of common tasks and commands for project development

#### Configuration: [Makefile](../../../Makefile)

#### Usage

---

### Commitizen

#### Documentation

[commitizen](https://commitizen-tools.github.io/commitizen/)

#### System Context

All

#### Purpose

Standardizing commit messages according to Conventional Commits

#### Configuration

[.cz.yaml](../../../.cz.yaml)

#### Usage: 

Run `make commit` to create a commit with a standardized message

---

## GoReleaser

#### Documentation: [GoReleaser](https://goreleaser.com/)

#### System Context

Core / Backend

#### Purpose

Automating the release process for Go projects

#### Configuration

[.goreleaser.yml](../../../.goreleaser.yml)

#### Usage

Run `make release` and follow the prompts to create a new release

---

## OapiCodeGen

#### Documentation

[OapiCodeGen](https://github.com/oapi-codegen/oapi-codegen/tree/main)

#### System Context

Core / Backend

#### Purpose

Generating Go server and client code from OpenAPI 3.0 specifications

#### Configuration

[cfg.yaml](../../../api/rest/pingpong/v1/server/cfg.yaml)

#### Usage

Run `make gen` to generate the server and client code

---

## OpenAPI-Generator-Cli

#### Documentation

[OpenAPI-Generator-Cli](https://github.com/OpenAPITools/openapi-generator-cli)

#### System Context

Frontend

#### Purpose

Generation of the typescript API client library to enable the consumption of REST APIs from OpenAPI specifications

#### Usage

---

## Golangci-lint

#### Documentation

[Golangci-lint](https://golangci-lint.run/docs/)

#### System Context

Core / Backend

#### Purpose

#### Usage

---

## GoFumpt

#### Documentation

[GoFumpt](https://github.com/mvdan/gofumpt)

#### System Context

Core / Backend

#### Purpose

Go code formatter, stricter than `gofmt`

#### Configuration

N/A

#### Usage

---

## Knip

#### Documentation

[Knip](https://knip.dev/)

#### System Context

Frontend

#### Purpose

Identifying unused dependencies in the frontend application

#### Configuration

[knip.config.js](../../../api/frontend/knip.json)

#### Usage

---

## Mockery

#### Documentation

[mockery](https://vektra.github.io/mockery/latest/)

#### System Context

Core / Backend

#### Purpose

Mock generation for Go interfaces

#### Configuration

[mockery.yml](../../../.mockery.yml)

#### Usage

---

## MkDocs + Material for MkDocs

#### Documentation

[mkdocs](https://www.mkdocs.org/), [Material for Mkdocs](https://squidfunk.github.io/mkdocs-material/)

#### System Context

Documentation

#### Purpose

#### Configuration

[mkdocs.yml](../../../mkdocs.yml)

#### Usage

---

## PNPM

#### Documentation

[pnpm](https://pnpm.io/motivation)

#### System Context

Frontend

#### Purpose

Package management for the frontend application

#### Configuration

[package.json](../../../api/frontend/package.json)

#### Usage

---

## UV

#### Documentation

[uv](https://docs.astral.sh/uv/)

#### System Context

Documentation

#### Purpose

Python environment and package management for the documentation site

#### Configuration

[pyproject.toml](../../pyproject.toml)

#### Usage

---

## Docker

Documentation: [docker](https://docs.docker.com/)

## Docker Compose

Documentation: [docker compose](https://docs.docker.com/compose/)

## Azure Developer CLI (azd)

Documentation: [Azure Developer CLI (azd)](https://learn.microsoft.com/en-us/azure/developer/azure-developer-cli/)

## Golang-Migrate

#### Documentation

[golang-migrate](https://github.com/golang-migrate/migrate)

#### System Context

Core / Backend

#### Purpose

Database migrations for Go applications

#### Configuration

[migrations/](../../../db/migrations)

#### Usage

---

## SQLC

#### Documentation

[sqlc](https://docs.sqlc.dev/en/latest/)

#### System Context

Core / Backend

#### Purpose

Generate type-safe Go code from SQL queries

#### Configuration

[.sqlc.yaml](../../../.sqlc.yaml)

#### Usage

---

## Vite

#### Documentation

[vite](https://vite.dev/guide/)

#### System Context

Frontend

#### Purpose

A development server and build tool for web projects

#### Configuration

[vite.config.ts](../../../api/frontend/vite.config.ts)

#### Usage

---

## Vue

#### Documentation

[vue](https://vuejs.org/guide/introduction)

#### System Context

Frontend

#### Purpose

#### Configuration

#### Usage
