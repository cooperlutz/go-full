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

-- name: FindQuestionsForExam :many
SELECT * FROM examination.questions
WHERE exam_id = $1;

-- name: AddQuestion :exec
INSERT INTO examination.questions (
    question_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    exam_id,
    index,
    answered,
    question_text,
    question_type,
    provided_answer,
    response_options
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12
);