-- name: FindExamByID :one
SELECT * FROM exam_library.exams WHERE exam_id = $1;

-- name: FindExamQuestionByID :one
SELECT * FROM exam_library.exam_questions WHERE exam_question_id = $1;