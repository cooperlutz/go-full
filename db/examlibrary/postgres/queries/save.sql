-- name: SaveExam :exec
INSERT INTO exam_library.exams (
    exam_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    name,
    grade_level,
    time_limit
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);

-- name: SaveExamQuestion :exec
INSERT INTO exam_library.exam_questions (
    exam_question_id,
    created_at,
    updated_at,
    deleted_at,
    deleted,
    exam_id,
    index,
    question_text,
    answer_text,
    question_type,
    possible_points,
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