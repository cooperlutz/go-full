ALTER TABLE examination.exams
ADD COLUMN completed BOOLEAN NOT NULL DEFAULT FALSE;

UPDATE examination.exams
SET completed = CASE
    WHEN state = 'completed' THEN TRUE
    ELSE FALSE
END;

ALTER TABLE examination.exams
DROP COLUMN state;
