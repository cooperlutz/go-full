# 00009: Event Driven

## Status

Accepted

## Context

In the context of the system, the system must be able to respond and react to events that occur within the domain.

## Decision

The decision was made to implement a mechanism (a reusable package) to support domain events and event-driven patterns, primarily and initially, pub-sub. The addition of this package and initial implementation of this solution will provide support for emitting, storing, reading, and responding to events that happen across the system.

### Implementation

Given we would like to minimize infrastructure, but maintain the ability to keep the system loosely coupled, while also keeping records of events that occur within the system, the decision has been made to utilize Postgres as the event / message store.

The decision has been made to leverage the [Watermill (Lic MIT)](https://github.com/ThreeDotsLabs/watermill) framework, which provide OOTB Postgres adapter along with available pubsub functionality.

## Alternatives Considered

- Custom in-memory Event Bus. While this is easily accomplished thanks to Go's concurrency functionality and would enable us to send and receive messages across channels implemented via each domain's module, it would require additional effort to support recording and storing events.

## Consequences

- We may find limitations in what postgres is able to handle which may make us look to an alternate infrastructure solution (e.g. RabbitMQ, Kafka, etc.)
- Additional network traffic will occur between the application and postgres database which may introduce performance challenges
- We become dependent on the functionality of Watermill and its development progress
