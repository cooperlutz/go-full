CREATE SCHEMA IF NOT EXISTS grading;

CREATE TABLE IF NOT EXISTS grading.exams (
    exam_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    -- 
    student_id UUID NOT NULL,
    library_exam_id UUID NOT NULL,
    examination_exam_id UUID NOT NULL,
    grading_completed BOOLEAN NOT NULL DEFAULT FALSE,
    total_points_received INTEGER,
    total_points_possible INTEGER
);

CREATE TABLE IF NOT EXISTS grading.questions (
    question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    exam_id UUID NOT NULL REFERENCES grading.exams(exam_id) ON DELETE CASCADE,
    -- 
    index INT NOT NULL,
    question_type TEXT NOT NULL,
    graded BOOLEAN NOT NULL DEFAULT FALSE,
    feedback TEXT,
    provided_answer TEXT NOT NULL,
    correct_answer TEXT,
    correctly_answered BOOLEAN,
    points_received INTEGER,
    points_possible INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_questions_exam_id ON grading.questions(exam_id);