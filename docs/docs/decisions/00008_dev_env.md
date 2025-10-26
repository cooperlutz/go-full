
# 00008: Development Environment

## Status

Accepted

## Context

In the context of Development of the system, a consistent environment is required in order to ensure the system is reproducible.

## Decision

The decision was made to utilize Docker & Docker Compose on developer workstations in order to provide consistency

### Implementation

Hardware: Local Workstation
Platform: MacOS Assumed
Tooling: installed to local workstation via brew
Deployment Configuration & Orchestration: Docker Compose

## Alternatives Considered

- Local K8s implementation (kind, minikube): these solutions would more effectively scale under the assumption that the 'production' deployment is running within a k8s cluster, however it also adds additional complexity and requires further development. Ultimately, docker compose is lighter weight and simpler to implement and suits our current needs, and enables us to shift toward a local k8s if and when the need presents itself.

## Consequences

- There are limitations with utilizing Postgres in containers, which could result in issues when running the application in production database servers.
- The current solution does not provide support for Windows Desktop OS, but given a majority of functionality is containerized, support can be added at a later point in time.
