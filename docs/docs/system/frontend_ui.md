# Frontend User Interface (UI)

## Frontend / User Interface

ADR Ref: [ADR-00005](../decisions/00005_frontend.md)

The Frontend (User Interface) is implemented as a Single-Page Application (SPA) developed in Vue. The frontend application code can be found in `/api/frontend` intentionally because following our goal to align the overall project to be semantic and idiomatic... the frontend web ui is really just *an interface* in the context of the overall system.

The Vue application is compiled, then the compiled frontend is embedded within the Go application. this can be seen within `api/frontend/frontend.go`.

Note: when Go compiles it looks for `.go` files and ignores anything else, therefore, the all of the Vue/Typescript/node_modules/etc. have no impact to Go compilation

Vite is simply used for building and development.

## Design System

ADR Ref: [ADR-00006](../decisions/00006_base_design_system.md)
