-- name: FindAllIncompleteExams :many
SELECT * FROM grading.exams WHERE grading_completed = FALSE;

-- name: GetExam :one
SELECT * FROM grading.exams
WHERE exam_id = $1;

-- name: GetQuestionsForExam :many
SELECT * FROM grading.questions
WHERE exam_id = $1
ORDER BY index ASC;

-- name: FindQuestionByExamIdAndQuestionIndex :one
SELECT * FROM grading.questions
WHERE exam_id = $1 AND index = $2;

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
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5,
    student_id = $6,
    library_exam_id = $7,
    examination_exam_id = $8,
    grading_completed = $9,
    total_points_received = $10,
    total_points_possible = $11
WHERE exam_id = $1;

-- name: UpdateQuestion :exec
UPDATE grading.questions SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5,
    exam_id = $6,
    index = $7,
    question_type = $8,
    graded = $9,
    feedback = $10,
    provided_answer = $11,
    correct_answer = $12,
    correctly_answered = $13,
    points_received = $14,
    points_possible = $15
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