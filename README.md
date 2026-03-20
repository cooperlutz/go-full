<div align="center">

# Go Full

[![Go Report Card](https://goreportcard.com/badge/github.com/cooperlutz/go-full)](https://goreportcard.com/report/github.com/cooperlutz/go-full) [![codecov](https://codecov.io/gh/cooperlutz/go-full/graph/badge.svg?token=5QZGPZU2SG)](https://codecov.io/gh/cooperlutz/go-full) ![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white) ![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=&logo=vuedotjs&logoColor=%234FC08D) ![Microsoft Azure](https://custom-icon-badges.demolab.com/badge/Microsoft%20Azure-0089D6?logo=msazure&logoColor=white) ![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?logo=github-actions&logoColor=white) ![Dependabot](https://img.shields.io/badge/dependabot-025E8C?style=&logo=dependabot&logoColor=white) ![OpenTelemetry](https://img.shields.io/badge/OpenTelemetry-4f62ad?&style=&logo=opentelemetry&logoColor=f5a800) ![OpenAPI](https://img.shields.io/badge/OpenAPI-6BA539?logo=openapiinitiative&logoColor=white) ![Tailwind CSS](https://img.shields.io/badge/Tailwind%20CSS-%2338B2AC.svg?logo=tailwind-css&logoColor=white) ![DaisyUI](https://img.shields.io/badge/DaisyUI-5A0EF8?logo=daisyui&logoColor=fff) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff) ![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=fff) ![Vitest](https://img.shields.io/badge/Vitest-6E9F18?logo=vitest&logoColor=fff) ![pnpm](https://img.shields.io/badge/pnpm-F69220?logo=pnpm&logoColor=fff) ![MkDocs](https://img.shields.io/badge/MkDocs-526CFE?logo=materialformkdocs&logoColor=fff) [![Playwright](https://custom-icon-badges.demolab.com/badge/Playwright-2EAD33?logo=playwright&logoColor=fff)](https://playwright-community.github.io/playwright-go/) ![sqlc](https://img.shields.io/badge/sqlc-cyan)
</div>

*Go Full* is a full-stack project boilerplate template, learning tool, and reference implementation demonstrating a variety of development patterns and practices.

## Motivation

In the movie, [Ratatouille](https://www.imdb.com/title/tt0382932/), the late Chef Gusteau has a saying that serves as a core theme of the film; "anyone can cook". Recently, AI Coding, Vibe Coding, and the like, have taken off and shifted the pardigm of development toward extremely fast, hands-off, AI developing AI, and various other rabbit holes. Ironically, the concepts and capabilities that AI Code Generation promise aren't necessarily all that new... code generation tools have been around since, well forever, as have low-code or no-code tools. At the end of the day, what matters most is the adoption, value, and functionality of the system being developed, along with the maintainability of the code that makes up the system.

Many higher level tools take away flexibility (and the fun) of development. Oftentimes, complex packages and libraries can abstract away an understanding of what/how/why the system does what it does. This project intends *not* to abstract away core functinality with light and intentional usage of more basic packages (not fully functional highly opinionated frameworks).

This project aims to provide a well-formed, "just enough" featured, minimalist boilerplate examplar under the idea that **"anyone can code"** *(warning: we may have surpassed our minimalist and just enough by a hair)*. And with that, we don't want to just spit out code that works (or looks like it works) as quickly as tokens could possibly be consumed, we instead focus on providing the components and resources to develop applications or services in a manner that prioritizes developers truly understanding the why, what, and how behind the things they are developing and making industry terminolgy and Buzzwords more real. We leverage code generation tools (mockery, oapi code generators, sqlc) to reduce boilerplate code development. Regardless of what is developed by hand or machine utilizing this project, the intent is to prioritize code **quality** over quantity, or more commonly, [Clean Code](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29) along with Product quality.

## Priorities & Decision Criteria

- **Anyone can code**
- Code and Feature Quality
- Go as a first class citizen
- Development focus and emphasis should be on the domain logic and user features
- Development should be fun!
- Semantic & Idiomatic code and tools
- Tools should be configuration-based via yaml
- Languages should be typesafe
- Everything as code
- not overengineered, but smart and scalable

## AI PoC Factory

The **AI PoC Factory** (`generate-poc`) lets you go from a plain-English business case to a fully scaffolded, API-ready module in a single command.

### How it works

1. You write a short markdown file describing a feature or product idea
2. The AI generates a module configuration (aggregates, commands, events, field types)
3. The modularizer scaffolds the full module: domain entities, application layer, REST API, database schema, migrations, frontend components, and generated code (sqlc, oapi-codegen)
4. The new module is auto-wired into the main application (imports, initialization, route mounting)

### What you get

A single `make generate-poc` command produces:

| Layer | What's generated |
|-------|-----------------|
| **REST API** | OpenAPI 3.0 spec, auto-generated Go server + TypeScript client, HTTP adapter with Chi router |
| **Database** | PostgreSQL schema, sqlc-annotated queries, type-safe Go query code, timestamped up/down migrations |
| **Domain** | Entity structs with base metadata (ID, timestamps, soft-delete), repository interfaces |
| **Application** | Command handlers, query handlers, event handlers, application service wiring |
| **Events** | Pub/sub event definitions (emitted/consumed) via Watermill + PostgreSQL |
| **Frontend** | Vue.js components, composables, router config, OpenAPI-generated TypeScript client |
| **Observability** | OpenTelemetry tracing spans baked into every layer |
| **Config** | `.sqlc.yaml`, `.mockery.yml`, `.openapitools.json` all updated automatically |
| **App wiring** | Module imported, instantiated, and route-mounted in `app/app.go` |

If the business case involves AI/LLM at runtime (e.g. "rate pitches using AI"), the tool detects this and injects the `pkg/ai` client dependency into the module.

### Quick start

```shell
# Write a business case
echo "# Expense Tracker
We need an API to track business expenses with description, amount, and category." > docs/my_idea.md

# Generate the full module
make generate-poc prompt="docs/my_idea.md"

# Review, customize entity fields, then run codegen and tests
make gen
make test
```

See [exampleguide.md](exampleguide.md) for a detailed walkthrough.

## Getting Started

> [!WARNING]
> This project has only been tested for development on macOS

### Prerequisites

- [Git](https://git-scm.com/)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Brew](https://brew.sh/)

### Installation

```shell
# Clone the repository
git clone https://github.com/cooperlutz/go-full.git

# Change directory to the project folder
cd go-full

# run make init to initialize and run the project
make init
```

Then open your browser to `http://app.lvh.me` to see the running application

### Additional Helpful Development Commands

```shell
make # end to end development tools

make compose # builds, deploys, and runs development environment

make commit # provides a mechanism to simplify conventional commits

make generate-poc prompt="docs/my_idea.md" # scaffold a new module from a business case

make modularizer # scaffold a new module from a modularizer.yaml config
```

For further details, please consult our [docs](https://cooperlutz.github.io/go-full/development/)
