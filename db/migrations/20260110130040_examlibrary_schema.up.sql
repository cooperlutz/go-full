-- Initialize exam_library schema
CREATE SCHEMA IF NOT EXISTS exam_library;

-- create exam_questions table
CREATE TABLE IF NOT EXISTS exam_library.exam_questions (
    exam_question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    exam_id UUID NOT NULL,
    -- 
    index INT NOT NULL,
    question_text TEXT NOT NULL,
    answer_text TEXT NULL,
    question_type TEXT NOT NULL,
    possible_points INT NOT NULL,
    response_options TEXT[]
);

-- create exams table
CREATE TABLE IF NOT EXISTS exam_library.exams (
    exam_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    -- 
    name TEXT NOT NULL,
    grade_level INT
);

-- create relationship between exam_questions and exams
-- add foreign key constraint to exam_questions table
ALTER TABLE exam_library.exam_questions
ADD CONSTRAINT fk_exam
FOREIGN KEY (exam_id)
REFERENCES exam_library.exams (exam_id)
ON DELETE CASCADE;

-- create index to optimize queries searching by exam_id
CREATE INDEX IF NOT EXISTS idx_exam_questions_exam_id
ON exam_library.exam_questions (exam_id);

-- create index to optimize queries searching by grade_level
CREATE INDEX IF NOT EXISTS idx_exams_grade_level
ON exam_library.exams (grade_level);