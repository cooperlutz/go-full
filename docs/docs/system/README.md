# System Architecture

## Full Stack

This project exemplifies a full stack system architecture inclusive of a frontend User Interface (UI), a backend Application Programming Interface (API) encompassing the core logic, and a data persistence layer.

![full stack](../_img/full_stack.drawio.png)

## Domain-Driven Design

This project adopts and adheres to the [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html) approach. Our bounded contexts are implemented as modules within the system. Each module adheres to the principles and patterns outlined in DDD, including domain entities, infrastructure repositories, and application services. Module implementations enable loose coupling of system components and an ability to easily extract modules into their own services if necessary. Further, each module is designed to enable the system to implement inversion of control and dependency injection principles.

### Further DDD Reading

- [DDD Reference](https://www.domainlanguage.com/wp-content/uploads/2016/05/DDD_Reference_2015-03.pdf)
- DDD Burger: A [good article](https://medium.com/@remast/the-ddd-hamburger-for-go-61dba99c4aaf), particularly referencing DDD in Go, that I think does a nice job reflecting the various layers as if we were building a hamburger.

## Modular Monolith

*ARE YOU KIDDING ME? WHERE ARE THE MICROSERVICES????* Yeah yeah yeah... well first and foremost, you could think of this project itself as a microservice... but more imporantly, that's not the *intent* of this project. This project intends to focus on developing a well formed full stack system in an easy to understand / learn / develop manner, and introducing microservices would exponentially complicate things. If you would prefer a series of microservices, you can easily separate the SPA and create a series of services leveraging this project as a starting point.

![modular mono](../_img/modular_monolith.drawio.png)

## Event-Driven

The system provides capabilities to support event-driven architecture patterns through the use of an event bus in the form of PostgreSQL tables. This is facilitated via implementation of the [Watermill](https://watermill.io/) framework. Read more about the decision: [ADR-00005](./decisions/00009_event_driven.md)

In order to reduce Direct dependency on Watermill, the [eeventdriven package](https://pkg.go.dev/github.com/cooperlutz/go-full@v0.1.22/pkg/eeventdriven) is provided to abstract away the Watermill implementation details from the rest of the system. This allows for easier swapping of the underlying api in the future, if desired.

The below [Component Diagram](https://c4model.com/diagrams/component) illustrates how the event-driven components are integrated into the overall system architecture.

![event driven](../_img/representative_component_diagram_eventdriven.drawio.png)

## Mono Repo Project Layout / Structure

This project is intentionally structured as a monorepo to display how all of the various pieces of the system are stitched together. We can follow the logical flow of the system from frontend view all the way through the backend database queries and every layer in between.

Given the core of the system is defined in Go, we adhere to commonly accepted Go project layout best practices with slight adaptations to suit our needs.

- Directory Structure Reference Point: [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

Root Directory Layout:

```shell
├── api # api
├── app # core application implementation
├── build # build files (dockerfiles)
├── cmd # primary commands for running the application
├── configs # configuration files
├── db # database code (migrations, queries, etc.)
├── deploy # deploy targets
├── docs # documentation site and content
├── examples # example code
├── internal # core logic
├── pkg # go packages, cross-cutting
├── test # e2e tests and test utilities
├── tools # tools for the project
└── vendor # vendor packages  https://go.dev/ref/mod#vendoring
```

### Frontend Vue Layout

The frontend application is defined as an embedded SPA, written in Vue.js, the project layout for the frontend vue application follows the pattern defined here [Vue Reference](https://vue-faq.org/en/development/project-structure.html#suitable-architecture-for-vue-3-web-application)
