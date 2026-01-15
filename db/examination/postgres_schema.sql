CREATE SCHEMA IF NOT EXISTS examination;

CREATE TABLE IF NOT EXISTS examination.exams (
    name TEXT NOT NULL,
    grade_level INT
);