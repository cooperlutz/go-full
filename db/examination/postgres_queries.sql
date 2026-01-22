-- name: FindAllExams :many
SELECT * FROM examination.exams;

-- name: AddExam :exec
INSERT INTO examination.exams (
    exam_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    student_id,
    completed,
    completed_at,
    started_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
);