ALTER TABLE grading.exams
ADD COLUMN grading_completed BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE grading.exams
SET grading_completed = CASE
    WHEN state = 'completed' THEN TRUE
    ELSE FALSE
END;

ALTER TABLE grading.exams
DROP COLUMN state;
