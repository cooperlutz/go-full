# System Modules

As noted in [ADR-00002](../decisions/00002_system_arch.md), the system architecture adopts a modular monolith approach, adhering to principles of Domain-Driven Design (DDD).

## Module Structure

Each system module is structured to encapsulate the various layers of the application, ensuring separation of concerns and maintainability. The typical structure of a module is as follows:

### core / backend

```shell
internal/DOMAIN_MODULE
├── api          # Interface Layer: API controllers and routes
├── app          # Application Layer: Application services and use cases
├── domain       # Domain Layer: Core business logic and domain entities
├── infra        # Infrastructure Layer: Data access and external services
└── DOMAIN_MODULE.go  # Module definition and entry point
```

#### Mock Implementations

```shell
test/mocks/DOMAIN_MODULE/
```

In order to configure the relevant mock implementations for use within tests, each *system module* will require updates to the `.mockery` configuration file to outline relevant *system module* packages.

### frontend

```shell
api/frontend/src/DOMAIN_MODULE
├── __tests__   # Tests specific to the module
├── assets       # Assets specific to the module
├── components   # Vue components specific to the module
├── composables  # Vue composables specific to the module
├── configs      # Configuration files specific to the module
├── layouts      # Vue layouts specific to the module
├── router      # Vue router definitions specific to the module
├── services     # Services (API clients) specific to the module
├── stores        # Vuex store modules specific to the module
├── styles       # Styles specific to the module
├── utils       # Utility functions specific to the module
├── views        # Vue views specific to the module
└── DOMAIN_MODULE.ts  # Module definition and entry point
```

### REST API

```shell
api/rest/DOMAIN_MODULE/{version}
├── client # Generated API client for the individual module version
│   ├── cfg.yaml # Configuration for the individual module version server
│   ├── client.gen.go # Generated client code for the individual module version
│   └── generate.go # Code generation script for the individual module version
├── server # Generated API server for the individual module
│   ├── cfg.yaml # Configuration for the individual module version server
│   └── generate.go # Code generation script for the individual module version
└── api.yaml # OpenAPI specification for the individual module version
```

### Database

Note: Migration files must be defined within a consolidated `db/migrations/` directory to ensure proper execution order during migration runs.

#### Migrations

```shell
db/migrations/ # Database migration files
```

#### Queries and Schemas

```shell
db/DOMAIN_MODULE
├── queries  # Database query files specific to the module
└── schemas  # Database schema files specific to the module
```

Additionally, each *system module* will require updates to the corresponding `.sqlc` file to outline the relevant configuration.
