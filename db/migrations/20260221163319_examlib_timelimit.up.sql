-- add time_limit column to exams table
ALTER TABLE exam_library.exams
    ADD COLUMN time_limit BIGINT DEFAULT 3600;
