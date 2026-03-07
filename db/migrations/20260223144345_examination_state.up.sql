-- This migration adds a new "state" column to the "exams" table and removes the "completed" boolean column.
-- The "state" column will be used to track the state of the exam (e.g., "in-progress", "completed", etc.) instead of a simple boolean.

ALTER TABLE examination.exams
ADD COLUMN state TEXT NOT NULL DEFAULT 'in-progress';

UPDATE examination.exams
SET state = CASE
    WHEN completed THEN 'completed'
    WHEN NOT completed AND started_at IS NOT NULL THEN 'in-progress'
    ELSE 'not-started'
END;

ALTER TABLE examination.exams
DROP COLUMN completed;