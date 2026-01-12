CREATE SCHEMA IF NOT EXISTS exam_library;

CREATE TABLE IF NOT EXISTS exam_library.exam_questions (
    exam_question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    exam_id UUID NOT NULL REFERENCES exam_library.exams(exam_id) ON DELETE CASCADE,
    -- 
    index INT NOT NULL,
    question_text TEXT NOT NULL,
    answer_text TEXT,
    question_type TEXT NOT NULL,
    possible_points INT NOT NULL,
    response_options TEXT[]
);