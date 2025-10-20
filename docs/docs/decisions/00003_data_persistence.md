# 00003: Data Persistence

## Status

Accepted

## Context

In the context of the overall system, persistent data storage will be required to store data. This decision influences and directly impacts development complexity, tooling, and system performance.

## Decision

The chosen approach for data persistence is to use a **relational database management system (RDBMS)** for structured data storage, leveraging its ACID properties to ensure data integrity and consistency. This decision was made to align with the project's requirements for complex querying and reporting.

### Implementation

- Database Schema Design: The database schema will be designed to reflect the domain model, ensuring that tables and relationships align with the application's requirements. Each module will have its own schema to encapsulate its data.
- ORM Usage: An Object-Relational Mapping (ORM) tool will **NOT** be used. Instead, direct SQL queries will be employed to interact with the database, providing greater control over query optimization and performance tuning. These queries will be defined and generated via [sqlc](https://sqlc.dev/)
- **PostgreSQL** has been selected as the specific RDBMS.

## Alternatives Considered

### Storage Options

- NoSQL Databases: While NoSQL databases offer flexibility in handling unstructured data and can scale horizontally, they may lack the robust transactional support required for this project. Given the current data requirements, the complexity introduced by NoSQL systems is not justified.
- In-Memory Databases: Although in-memory databases provide high-speed data access, they are not suitable for long-term data persistence due to volatility.
- Flat File Storage: While flat file storage is simple to implement, it lacks the advanced querying capabilities and data integrity features provided by RDBMS. This option was rejected due to scalability and performance concerns.

### Database Management Systems

- SQLite: While SQLite is lightweight and easy to set up, it is not suitable for applications requiring concurrent access by multiple users or high transaction volumes. Further, it would limit future scalability options.

## Consequences

As a consequence of this decision:

- Increased complexity to maintain the postgres database server
- Need to manage database migrations and versioning
- Improved data integrity and consistency
