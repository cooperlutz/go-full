CREATE TABLE IF NOT EXISTS exam_library.exams (
    exam_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    -- 
    name TEXT NOT NULL,
    grade_level INT,
    time_limit BIGINT
);
