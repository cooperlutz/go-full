# Modularizer Tool

The Modularizer Tool is designed to instantiate a new Go Full module adhering to the recommended module structure. It provides a streamlined process for creating a new module with as much boilerplate code as possible, allowing developers to focus on implementing the core logic of their module.

## Example Usage

The below command demonstrates how to use the Modularizer Tool to create a new module named "work" with an aggregate root of "work_item". It includes commands for "assign_to" and "start_work", events for "work_item_created" and "work_item_started", and a query for "find_all_work_items".

```bash
go run tools/modularizer/cmd/modularize/main.go modularize --config tools/modularizer/modularizer.yaml
```

## Outputs

Upon running the above command, the Modularizer Tool will generate a new module with the following structure along with generating the relevant files for the module, its adapters, its domain, and its application commands, events, and queries.

```bash
├── api
│   ├── frontend/
│   │   ├── src/work
│   │   └── test/mocks/work
│   └── rest/work
├── db
│   ├── migrations/
│   └── work/
├── internal/work/
│   ├── adapters
│   │   ├── inbound/
│   │   └── outbound/
│   ├── app
│   │   ├── command/
│   │   ├── event/
│   │   └── query/
│   └── domain/
│       └── work/
```

## How it works

The Modularizer Tool makes use of the `text/template` package to generate files based on predefined templates. It takes in command-line arguments to specify the module name, aggregate root, commands, events, and queries, and then creates the necessary directory structure and files for the new module.

It also builds on top of the other code generation tools utilized within the project, SqlC, OapiCodeGen, openapi-generator to further simplify the process.
