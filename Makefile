.DEFAULT_GOAL := all

# ENVIRONMENT VARIABLES Loading
ifneq ($(wildcard .env),)
include .env
export
else
$(warning WARNING: .env file not found! Using .env.example)
include .env.example
export
endif

# CONSTANTS
GO_MODULE_NAME := $(shell go list -m | head -n 1)
GO_VERSION := $(shell go list -f {{.GoVersion}} -m)
PNPM := pnpm -C ./api/frontend
APP_VERSION := $(shell cz version -p)
APP_NAME := $(GO_MODULE_NAME)
AZD_CONF := ./deploy/azure

export APP_VERSION
export APP_NAME


############################################################################
#                             ALL                                          #
############################################################################
all: install deps deps-audit gen build-fe lintfmt test cover-filter build-local ### run all setup tasks
.PHONY: all

############################################################################
#                            DEPENDENCIES                                  #
############################################################################

deps: ### go mod tidy + verify + vendor
	go mod tidy && go mod verify && go mod vendor
	$(PNPM) deps
.PHONY: deps

deps-audit: ### check dependencies vulnerabilities
	go tool govulncheck ./...
.PHONY: deps-audit

############################################################################
#                            LINT & FORMAT                                 #
############################################################################

lintfmt: lint format ### run all linters and formatters
.PHONY: lintfmt

format: format-be format-fe ### run all formatters
.PHONY: format

lint: lint-be lint-fe ### run all linters
.PHONY: lint

lint-be: ### check by golangci linter
	go tool golangci-lint run --config ./.golangci.yml --fix
.PHONY: lint-be

lint-fe: ### lint frontend code
	$(PNPM) lint
.PHONY: lint-fe

format-fe: ### format frontend code
	$(PNPM) format
.PHONY: format-fe

format-be: ### Run code formatter
	go tool gofumpt -l -w .
	go tool gci write . --skip-generated --skip-vendor -s standard -s default -s "prefix($(GO_MODULE_NAME))"
.PHONY: format-be

############################################################################
#                            GENERATORS                                   #
############################################################################

gen: mock gen-api queries ### run all code generators
.PHONY: gen

gen-api: gen-api-be gen-api-fe ### generate all api
.PHONY: gen-api

gen-api-be: ### generate api
	go generate ./api/...
.PHONY: gen-api-be

gen-api-fe: ### generate frontend api
	$(PNPM) openapi
.PHONY: gen-api-fe

queries: ### generate queries
	sqlc generate -f ./.sqlc.yaml
.PHONY: queries

# Note: .mockery.yml config file is used for mockery settings
mock: ### generate mock interfaces
	mockery
.PHONY: mock

pre-commit: deps deps-audit gen lint format ### run pre-commit
	$(PNPM) pre-commit
.PHONY: pre-commit

MIGMODULE=pingpong
MIGDBTYPE=postgres
MIGNAME=init_schema
migrate-create:  ### create new migration, run it like this "make migrate-create MIGMODULE=pingpong MIGDBTYPE=postgres MIGNAME=init_schema"
	@echo "Creating $(MIGNAME) migration for $(MIGMODULE)/$(MIGDBTYPE) database!"
	migrate create -ext sql -dir db/migrations $(MIGNAME)
.PHONY: migrate-create
	
# .cz.yaml config file is used for commitizen settings
commit: ### commit changes
	cz commit
.PHONY: commit

############################################################################
#                            DEPLOY                                        #
############################################################################

init-deploy:
	azd pipeline config --auth-type federated -C ${AZD_CONF}
.PHONY: init-deploy

azd-show: ### run Azure Developer CLI
	azd show -C ${AZD_CONF}
.PHONY: azd-show

azd-deploy: ### run azd deployment
	azd deploy -C ${AZD_CONF} --debug

azd-down: ### run azd down
	azd down -C ${AZD_CONF} --debug

azd-provision: ### run azd provision
	azd provision -C ${AZD_CONF} --debug

############################################################################
#                            BUILD                                        #
############################################################################

# Notes: 
# .goreleaser.yaml config file is used for goreleaser settings
# the `build-fe` task is called via the `before.hooks` section of the `.goreleaser.yaml` file
build: ## goreleaser build 
	goreleaser build --clean
.PHONY: build

build-local: ## goreleaser build 
	goreleaser release --snapshot --clean 
.PHONY: build-local

release-local: ## goreleaser build 
	goreleaser release --snapshot --clean --skip=publish
.PHONY: release-local

build-fe: ### build frontend
	$(PNPM) build
.PHONY: build-fe

############################################################################
#                            WORKFLOWS                                    #
############################################################################

pre-wflow: build-fe  ### prehook for ci tasks
	if [ ! -d .coverage ]; then mkdir .coverage; else echo ".coverage directory already exists, skipping creation."; fi
.PHONY: pre-wflow

ci: pre-wflow deps deps-audit lint format test cover-filter compose-e2e ### run all ci tasks
.PHONY: ci

############################################################################
#                            DEVELOPMENT ENV                               #
############################################################################

compose: release-local run ### run docker compose
.PHONY: compose

run: ### Run Local
	docker compose --env-file ".env" -f ./deploy/compose/docker-compose.yml up --build
.PHONY: run

install: install-tools install-brews install-playwright ### install all dependencies
.PHONY: install

install-brews: ### install brew packages
	brew update && brew upgrade
# languages
	brew install go
	brew install node
	brew install pnpm
# dev tools
	brew install commitizen
	brew install golang-migrate
	brew install sqlc
	brew install mockery
	brew install goreleaser
# deployment tools
	brew tap azure/azd && brew install azd
.PHONY: install-brews

install-tools: ### install tools
	go install tool
.PHONY: install-tools

install-playwright: ### install playwright browsers
	go run github.com/playwright-community/playwright-go/cmd/playwright@v0.5200.1 install --with-deps
.PHONY: install-playwright

############################################################################
#                            TESTING                                     #
############################################################################

compose-e2e: release-local ### run containerized e2e tests with docker compose
	docker compose -f ./test/e2e/docker-compose.yml up --build --abort-on-container-exit --exit-code-from e2e
.PHONY: compose-e2e

e2e: ### run e2e tests
	go test -v ./test/e2e/...
.PHONY: e2e

test-fe: ### run frontend test once
	$(PNPM) test:once
.PHONY: test-fe

test-be: ### run test
	go test -v -race -covermode atomic -coverprofile=.coverage/coverage.out ./internal/... ./app/... ./pkg/...
.PHONY: test-be

test: test-be test-fe ### run test
.PHONY: test

cover-filter: ### filter coverage report
	cat .coverage/coverage.out | grep -v ".gen.go" > .coverage/coverage.filtered.out  
.PHONY: cover-filter

cover-be: test-be cover-filter ### backend coverage report
# 	Filter out generated files from coverage report
	cat .coverage/coverage.out | grep -v ".gen.go" > .coverage/coverage.filtered.out  
# 	Display coverage report
	go tool cover -html=.coverage/coverage.filtered.out
# 	Print total coverage percentage
	go tool cover -func=.coverage/coverage.filtered.out | grep total | awk '{print "Total Coverage: " $$3}'
.PHONY: cover-be

cover-fe: ### frontend coverage report
	$(PNPM) coverage
.PHONY: cover-fe

############################################################################
#                            UTILITIES                                     #
############################################################################

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

module: ### print go module name
	@echo $(GO_MODULE_NAME)
.PHONY: module

go-version:  ### print go version
	@echo $(GO_VERSION)
.PHONY: go-version

loc: ### lines of code
	git ls-files | xargs wc -l
.PHONY: loc

version:  ### print app version
	@echo $(APP_VERSION)
.PHONY: version

docs: ### generate documentation and run mkdocs server
	cd docs && uv sync && uv run mkdocs serve
.PHONY: docs

coverage-directory: ### create .coverage directory
	@mkdir -p .coverage
.PHONY: coverage-directory

init-env-file: ### initialize .env file from .env.example if it does not exist
	if [ ! -f ".env" ]; then cp .env.example .env; else echo ".env file already exists, skipping copy."; fi
.PHONY: init-env-file

init: coverage-directory init-env-file all compose ### initialize project
.PHONY: init