-- name: FindAllExams :many
SELECT * FROM exam_library.exams;

-- name: FindAllExamQuestions :many
SELECT * FROM exam_library.exam_questions
WHERE exam_id = $1;
