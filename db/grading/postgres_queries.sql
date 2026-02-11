-- name: FindAllIncompleteExams :many
SELECT * FROM grading.exams WHERE grading_completed = FALSE;

-- name: GetExam :one
SELECT * FROM grading.exams
WHERE exam_id = $1;

-- name: GetQuestion :one
SELECT * FROM grading.questions
WHERE question_id = $1;

-- name: AddExam :exec
INSERT INTO grading.exams (
    exam_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    student_id,
    library_exam_id,
    examination_exam_id,
    grading_completed,
    total_points_received,
    total_points_possible
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
    $11
);

-- name: UpdateExam :exec
UPDATE grading.exams SET
    updated_at = $2,
    deleted_at = $3,
    deleted = $4,
    student_id = $5,
    library_exam_id = $6,
    examination_exam_id = $7,
    grading_completed = $8,
    total_points_received = $9,
    total_points_possible = $10
WHERE exam_id = $1;

-- name: UpdateQuestion :exec
UPDATE grading.questions SET
    updated_at = $2,
    deleted_at = $3,
    deleted = $4,
    exam_id = $5,
    index = $6,
    question_type = $7,
    graded = $8,
    feedback = $9,
    provided_answer = $10,
    correct_answer = $11,
    correctly_answered = $12,
    points_received = $13,
    points_possible = $14
WHERE question_id = $1;

-- name: AddQuestion :exec
INSERT INTO grading.questions (
    question_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    exam_id,
    index,
    question_type,
    graded,
    feedback,
    provided_answer,
    correct_answer,
    correctly_answered,
    points_received,
    points_possible
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
    $12,
    $13,
    $14,
    $15
);