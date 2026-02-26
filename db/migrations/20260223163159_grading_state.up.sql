-- This migration adds a new "state" column to the "exams" table and removes the "grading_completed" boolean column.
-- The "state" column will be used to track the state of the exam (e.g., "in-progress", "completed", etc.) instead of a simple boolean.

ALTER TABLE grading.exams
ADD COLUMN state TEXT NOT NULL DEFAULT 'not-started';

UPDATE grading.exams
SET state = CASE
    WHEN grading_completed THEN 'completed'
    WHEN NOT grading_completed AND updated_at IS NOT NULL THEN 'in-progress'
    ELSE 'not-started'
END;

ALTER TABLE grading.exams
DROP COLUMN grading_completed;