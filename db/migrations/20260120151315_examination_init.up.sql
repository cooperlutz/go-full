CREATE SCHEMA IF NOT EXISTS examination;

CREATE TABLE IF NOT EXISTS examination.exams (
    exam_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
);