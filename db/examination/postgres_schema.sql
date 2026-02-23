CREATE SCHEMA IF NOT EXISTS examination;

CREATE TABLE IF NOT EXISTS examination.exams (
    exam_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    -- 
    student_id UUID NOT NULL,
    library_exam_id UUID NOT NULL,
    state TEXT NOT NULL,
    completed_at TIMESTAMPTZ,
    started_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS examination.questions (
    question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    exam_id UUID NOT NULL REFERENCES examination.exams(exam_id) ON DELETE CASCADE,
    -- 
    index INT NOT NULL,
    answered BOOLEAN NOT NULL DEFAULT FALSE,
    question_text TEXT NOT NULL,
    question_type TEXT NOT NULL,
    provided_answer TEXT,
    response_options TEXT[]
);