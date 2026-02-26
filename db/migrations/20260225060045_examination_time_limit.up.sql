ALTER TABLE examination.exams
    ADD COLUMN time_limit BIGINT NOT NULL DEFAULT 3600, -- time limit in seconds, default to 1 hour
    ADD COLUMN time_of_time_limit TIMESTAMPTZ;

-- Update existing exams to set time_of_time_limit based on started_at and time_limit
UPDATE examination.exams
SET time_of_time_limit = started_at + (time_limit * INTERVAL '1 second')
WHERE started_at IS NOT NULL;