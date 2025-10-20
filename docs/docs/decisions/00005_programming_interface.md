# 00005: Initial Application Programming Interface

## Status

Accepted

## Context

In order to enable programmatic access to the core / backend, an interface needs to be exposed.

## Decision

The initial programming interface will be a RESTful API adhering to OpenAPI specifications. The OpenAPI specification will be utilized to generate go server code for the backend, go client code which can be used for e2e tests and other go based clients, and a typescript client which can be used by the frontend.

## Alternatives Considered

- GraphQL
- gRPC

## Consequences

The OpenAPI specification will need to be maintained as the system evolves. The OpenAPI specification will drive much of the development of the system as it will be used to generate significant portions of code for both the backend and frontend.
