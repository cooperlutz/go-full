# DB - Database Files

The `db` directory contains components relevant at the database layer.

Core application code MUST NOT import from `db/`

Relevant Project Tooling:

- sqlc
- golang-migrate

## DB Folder Structure

- `{domain}/db/{dbtype}/migrations/` contains sql statements to support database migrations
- `{domain}/db/{dbtype}/queries/` contains sql query statements to be leveraged by sqlc in order to generate the associated go code
- `{domain}/db/{dbtype}/schemas/` contains sql statements defining the database/table schema
