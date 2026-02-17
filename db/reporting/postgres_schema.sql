CREATE SCHEMA IF NOT EXISTS reporting;

CREATE TABLE IF NOT EXISTS reporting.reports (
    name TEXT PRIMARY KEY,
    description TEXT,
    component_definitions JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS reporting.metrics (
    name TEXT PRIMARY KEY,
    value DOUBLE PRECISION NOT NULL
);

