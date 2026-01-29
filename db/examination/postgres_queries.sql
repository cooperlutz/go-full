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

-- name: GetQuestion :one
SELECT * FROM examination.questions
WHERE question_id = $1;

-- name: GetQuestionsByExam :many
SELECT * FROM examination.questions
WHERE exam_id = $1
ORDER BY index ASC;

-- name: GetQuestionByExamAndIndex :one
SELECT * FROM examination.questions
WHERE exam_id = $1 AND index = $2;

-- name: GetExam :one
SELECT * FROM examination.exams
WHERE exam_id = $1;

-- name: SaveExam :exec
UPDATE examination.exams
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5,
    student_id = $6,
    completed = $7,
    completed_at = $8,
    started_at = $9
WHERE exam_id = $1;

-- name: SaveQuestion :exec
UPDATE examination.questions
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5,
    exam_id = $6,
    index = $7,
    answered = $8,
    question_text = $9,
    question_type = $10,
    provided_answer = $11,
    response_options = $12
WHERE question_id = $1;